package monitor

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/logger"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/rpc"
)

type ChainState struct {
	Chain            string
	Monitors         []MonitorEntry
	Commands         []Command
	WatchlistModTime time.Time
	CommandsModTime  time.Time
	WatchlistPath    string
	CommandsPath     string
}

type MonitorService struct {
	logger            *slog.Logger
	chains            []string
	config            MonitorConfig
	chainStates       map[string]*ChainState
	paused            bool
	ctx               context.Context
	cancel            context.CancelFunc
	processedCount    atomic.Int64      // Atomic counter for current round
	monitorAddressMap map[string]string // address -> monitor mapping for logging
	processing        atomic.Bool       // True if a round is currently in progress
	currentAddress    atomic.Value      // string: Address currently being processed
}

func NewMonitorService(logger *slog.Logger, chains []string, config MonitorConfig) (*MonitorService, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &MonitorService{
		logger:            logger,
		chains:            chains,
		config:            config,
		chainStates:       make(map[string]*ChainState),
		monitorAddressMap: make(map[string]string),
		paused:            false,
		ctx:               ctx,
		cancel:            cancel,
	}, nil
}

func (s *MonitorService) Name() string {
	return "monitor"
}

func (s *MonitorService) Initialize() error {
	for _, chain := range s.chains {
		var monitors []MonitorEntry
		var err error
		var watchlistPath string

		if s.config.WatchlistDir == "all" {
			monitors, err = DiscoverMonitors(chain)
			if err != nil {
				return fmt.Errorf("failed to discover monitors for chain %s: %v", chain, err)
			}
		} else {
			watchlistPath = filepath.Join(s.config.WatchlistDir, fmt.Sprintf("watchlist-%s.txt", chain))
			monitors, err = ParseWatchlist(watchlistPath)
			if err != nil {
				return fmt.Errorf("failed to parse watchlist for chain %s: %v", chain, err)
			}
		}

		commandsPath := filepath.Join(s.config.CommandsDir, fmt.Sprintf("commands-%s.yaml", chain))
		commands, err := ParseCommands(commandsPath)
		if err != nil {
			return fmt.Errorf("failed to parse commands for chain %s: %v", chain, err)
		}

		// Get initial modification times
		watchlistModTime := getFileModTime(watchlistPath)
		commandsModTime := getFileModTime(commandsPath)

		s.chainStates[chain] = &ChainState{
			Chain:            chain,
			Monitors:         monitors,
			Commands:         commands,
			WatchlistPath:    watchlistPath,
			CommandsPath:     commandsPath,
			WatchlistModTime: watchlistModTime,
			CommandsModTime:  commandsModTime,
		}

		s.logger.Info("Initialized chain", "chain", chain, "monitors", len(monitors), "commands", len(commands))
	}

	return nil
}

func (s *MonitorService) Process(ready chan bool) error {
	ready <- true

	// Event-driven: processing happens via ProcessBlockRange() called by coordinator
	<-s.ctx.Done()
	return nil
}

func (s *MonitorService) getChainHead(chain string) (uint64, error) {
	conn := rpc.TempConnection(chain)
	return uint64(conn.GetLatestBlockNumber()), nil
}

func (s *MonitorService) shouldFailEarly(metrics IterationMetrics) bool {
	if metrics.MonitorsTotal == 0 {
		return false
	}

	failureRate := float64(metrics.MonitorsFailed) / float64(metrics.MonitorsTotal)
	return failureRate > 0.5
}

func (s *MonitorService) logMetrics(metrics IterationMetrics) {
	s.logger.Info("Iteration complete",
		"duration", metrics.Duration,
		"batches_total", metrics.BatchesTotal,
		"batches_failed", metrics.BatchesFailed,
		"monitors_total", metrics.MonitorsTotal,
		"monitors_success", metrics.MonitorsSuccess,
		"monitors_failed", metrics.MonitorsFailed,
	)
}

func (s *MonitorService) Cleanup() {
	s.logger.Info("Cleaning up monitor service")
	s.cancel()
}

func (s *MonitorService) Pause() bool {
	s.paused = true
	s.logger.Info("Monitor service paused")
	return true
}

func (s *MonitorService) Unpause() bool {
	s.paused = false
	s.logger.Info("Monitor service unpaused")
	return true
}

func (s *MonitorService) IsPaused() bool {
	return s.paused
}

func (s *MonitorService) Logger() *slog.Logger {
	return s.logger
}

type IterationMetrics struct {
	StartTime       time.Time
	Duration        time.Duration
	BatchesTotal    int
	BatchesFailed   int
	MonitorsTotal   int
	MonitorsSuccess int
	MonitorsFailed  int
}

// ProcessMonitors triggers the monitor to process newly scraped blocks for a chain
// This is called by the coordinator when the scraper completes a batch
// Monitors track their own state and determine which blocks to process
func (s *MonitorService) ProcessMonitors(chain string) error {
	defer func() {
		logger.SetLoggerWriter(io.Discard)
	}()
	logger.SetLoggerWriter(os.Stderr)

	fmt.Println("[MONITOR] ProcessMonitors called for chain:", chain)

	// Skip if already processing a round
	if s.processing.Load() {
		processed := s.processedCount.Load()
		state := s.chainStates[chain]
		total := 0
		if state != nil {
			total = len(state.Monitors)
		}
		currentAddr := ""
		if addr := s.currentAddress.Load(); addr != nil {
			currentAddr = addr.(string)
		}
		fmt.Printf("[MONITOR] Skipping - previous round %d/%d (%s) still in progress for chain: %s\n", processed, total, currentAddr[:10]+"...", chain)
		return nil
	}
	s.processing.Store(true)
	defer s.processing.Store(false)

	// Check if config files have been modified and reload if necessary
	if err := s.checkAndReloadConfig(chain); err != nil {
		fmt.Printf("[MONITOR] Warning: Failed to reload config for %s: %v\n", chain, err)
	}

	if s.IsPaused() {
		// fmt.Println("[MONITOR] ProcessMonitors skipped - service is paused for chain:", chain)
		return nil
	}

	state := s.chainStates[chain]
	if state == nil {
		// fmt.Println("[MONITOR] ERROR: ProcessMonitors failed - no state for chain:", chain)
		return fmt.Errorf("no state for chain %s", chain)
	}

	if len(state.Monitors) == 0 {
		// fmt.Println("[MONITOR] ProcessMonitors skipped - no monitors configured for chain:", chain)
		return nil
	}

	metrics := IterationMetrics{
		StartTime: time.Now(),
	}

	// Get current head of chain to determine processing range
	latestBlock, err := s.getChainHead(chain)
	if err != nil {
		// fmt.Println("[MONITOR] ERROR: ProcessMonitors failed - could not get chain head for", chain, "error:", err)
		return fmt.Errorf("failed to get chain head: %w", err)
	}

	fmt.Printf("[MONITOR] Starting monitor processing for chain: %s, latestBlock: %d, monitorCount: %d\n", chain, latestBlock, len(state.Monitors))

	// Reset progress counter for this round
	s.processedCount.Store(0)

	// Build address map for progress logging
	s.monitorAddressMap = make(map[string]string)
	for _, m := range state.Monitors {
		s.monitorAddressMap[m.Address] = m.Address
	}

	// Refresh monitor states to get current LastScanned values
	// fmt.Println("[MONITOR] Refreshing monitor states from .mon.bin files...")
	for i := range state.Monitors {
		lastScanned, err := readMonitorState(chain, state.Monitors[i].Address)
		if err != nil {
			fmt.Printf("[MONITOR] Warning: Could not read state for %s: %v, keeping previous value %d\n",
				state.Monitors[i].Address, err, state.Monitors[i].StartingBlock)
		} else {
			if lastScanned != state.Monitors[i].StartingBlock {
				// fmt.Printf("[MONITOR] Updated %s: StartingBlock %d -> %d\n",
				// 	state.Monitors[i].Address, state.Monitors[i].StartingBlock, lastScanned)
				state.Monitors[i].StartingBlock = lastScanned
			}
		}
	}

	// TEMPORARY: Sequential processing for debugging (bypasses worker pool)
	useSequential := true // Set to false to re-enable worker pool

	if useSequential {
		// Sequential processing with batching
		executor := NewShellExecutor()
		batchSize := s.config.BatchSize

		// Process monitors in batches
		for i := 0; i < len(state.Monitors); i += batchSize {
			end := i + batchSize
			if end > len(state.Monitors) {
				end = len(state.Monitors)
			}
			batch := state.Monitors[i:end]

			// Collect addresses and check if any need processing
			var addresses []string
			var firstAddr string
			skippedAll := true

			for _, entry := range batch {
				startBlock := entry.StartingBlock
				endBlock := latestBlock

				if s.config.MaxBlocksPerRun > 0 && endBlock-startBlock > uint64(s.config.MaxBlocksPerRun) {
					endBlock = startBlock + uint64(s.config.MaxBlocksPerRun)
				}

				if endBlock <= startBlock {
					fmt.Printf("[MONITOR] Monitor skipped - no new blocks: chain=%s address=%s startBlock=%d endBlock=%d\n", chain, entry.Address, startBlock, endBlock)
					continue
				}

				addresses = append(addresses, entry.Address)
				if firstAddr == "" {
					firstAddr = entry.Address
				}
				skippedAll = false
				metrics.MonitorsTotal++
			}

			if skippedAll || len(addresses) == 0 {
				continue
			}

			// Update progress for this batch
			s.currentAddress.Store(firstAddr)
			processed := s.processedCount.Add(int64(len(addresses)))
			fmt.Printf("[MONITOR] Processing batch %d/%d (%d addresses, starting with %s)\n",
				processed, len(state.Monitors), len(addresses), firstAddr[:10]+"...")

			// Build template vars with multiple addresses
			vars := TemplateVars{
				Addresses:  addresses,
				Address:    firstAddr, // Keep for backwards compatibility
				Chain:      state.Chain,
				FirstBlock: batch[0].StartingBlock,
				LastBlock:  latestBlock,
				BlockCount: latestBlock - batch[0].StartingBlock + 1,
			}

			// Execute commands for this batch
			var executeErr error
			for _, cmd := range state.Commands {
				if err := executor.Execute(context.Background(), cmd, vars); err != nil {
					executeErr = err
					break
				}
			}

			if executeErr != nil {
				metrics.MonitorsFailed += len(addresses)
				s.logger.Warn("Batch processing failed", "chain", state.Chain, "addresses", len(addresses), "error", executeErr)
			} else {
				metrics.MonitorsSuccess += len(addresses)
			}
		}
		fmt.Println() // Newline after processing
	} else {
		// Process monitors with worker pool (original code)
		pool := NewWorkerPool(s.config.Concurrency, NewShellExecutor())

		// Set up real-time progress callback
		pool.onProgressUpdate = func(result JobResult) {
			s.currentAddress.Store(result.Address)
			processed := s.processedCount.Add(1)
			fmt.Printf("\r[MONITOR] Processing monitors %d/%d (%s)%s",
				processed, metrics.MonitorsTotal, result.Address[:10]+"...", strings.Repeat(" ", 30))
		}

		pool.Start()

		for _, entry := range state.Monitors {
			startBlock := entry.StartingBlock
			endBlock := latestBlock

			if s.config.MaxBlocksPerRun > 0 && endBlock-startBlock > uint64(s.config.MaxBlocksPerRun) {
				endBlock = startBlock + uint64(s.config.MaxBlocksPerRun)
			}

			if endBlock <= startBlock {
				fmt.Printf("[MONITOR] Monitor skipped - no new blocks: chain=%s address=%s startBlock=%d endBlock=%d\n", chain, entry.Address, startBlock, endBlock)
				continue
			}

			vars := TemplateVars{
				Address:    entry.Address,
				Chain:      state.Chain,
				FirstBlock: startBlock,
				LastBlock:  endBlock,
				BlockCount: endBlock - startBlock + 1,
			}

			job := MonitorJob{
				Entry:    entry,
				Commands: state.Commands,
				Vars:     vars,
			}

			pool.Submit(job)
			metrics.MonitorsTotal++
		}

		fmt.Printf("[MONITOR] %d monitors submitted to worker pool for chain: %s\n", metrics.MonitorsTotal, chain)

		results := pool.Wait()
		fmt.Println() // Newline after progress updates

		for _, result := range results {
			if result.Success {
				metrics.MonitorsSuccess++
			} else {
				metrics.MonitorsFailed++
				s.logger.Warn("Monitor processing failed", "chain", state.Chain, "address", result.Address, "error", result.Error)
			}
		}
	}

	metrics.Duration = time.Since(metrics.StartTime)
	s.logMetrics(metrics)

	if s.shouldFailEarly(metrics) {
		return fmt.Errorf("fail-early triggered: high failure rate (%d/%d failed)", metrics.MonitorsFailed, metrics.MonitorsTotal)
	}

	return nil
}

// getFileModTime returns the modification time of a file, or zero time if file doesn't exist
func getFileModTime(path string) time.Time {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}

// checkAndReloadConfig checks if watchlist or commands files have been modified and reloads them
func (s *MonitorService) checkAndReloadConfig(chain string) error {
	state := s.chainStates[chain]
	if state == nil {
		return nil
	}

	reloadNeeded := false

	// Check watchlist file
	if state.WatchlistPath != "" {
		currentModTime := getFileModTime(state.WatchlistPath)
		if !currentModTime.IsZero() && currentModTime.After(state.WatchlistModTime) {
			fmt.Printf("[MONITOR] Watchlist file changed for %s, reloading...\n", chain)
			monitors, err := ParseWatchlist(state.WatchlistPath)
			if err != nil {
				return fmt.Errorf("failed to reload watchlist: %v", err)
			}
			state.Monitors = monitors
			state.WatchlistModTime = currentModTime
			reloadNeeded = true
			fmt.Printf("[MONITOR] Reloaded watchlist for %s: %d monitors\n", chain, len(monitors))
		}
	}

	// Check commands file
	if state.CommandsPath != "" {
		currentModTime := getFileModTime(state.CommandsPath)
		if !currentModTime.IsZero() && currentModTime.After(state.CommandsModTime) {
			fmt.Printf("[MONITOR] Commands file changed for %s, reloading...\n", chain)
			commands, err := ParseCommands(state.CommandsPath)
			if err != nil {
				return fmt.Errorf("failed to reload commands: %v", err)
			}
			state.Commands = commands
			state.CommandsModTime = currentModTime
			reloadNeeded = true
			fmt.Printf("[MONITOR] Reloaded commands for %s: %d commands\n", chain, len(commands))
		}
	}

	if reloadNeeded {
		s.logger.Info("Configuration reloaded", "chain", chain, "monitors", len(state.Monitors), "commands", len(state.Commands))
	}

	return nil
}
