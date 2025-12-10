package monitor

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type ChainState struct {
	Chain    string
	Monitors []MonitorEntry
	Commands []Command
}

type MonitorService struct {
	logger      *slog.Logger
	chains      []string
	config      MonitorConfig
	chainStates map[string]*ChainState
	paused      bool
	ctx         context.Context
	cancel      context.CancelFunc
}

func NewMonitorService(logger *slog.Logger, chains []string, config MonitorConfig) (*MonitorService, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &MonitorService{
		logger:      logger,
		chains:      chains,
		config:      config,
		chainStates: make(map[string]*ChainState),
		paused:      false,
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (s *MonitorService) Name() string {
	return "monitor"
}

func (s *MonitorService) Initialize() error {
	for _, chain := range s.chains {
		var monitors []MonitorEntry
		var err error

		if s.config.WatchlistDir == "all" {
			monitors, err = DiscoverMonitors(chain)
			if err != nil {
				return fmt.Errorf("failed to discover monitors for chain %s: %v", chain, err)
			}
		} else {
			watchlistPath := filepath.Join(s.config.WatchlistDir, fmt.Sprintf("watchlist-%s.txt", chain))
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

		s.chainStates[chain] = &ChainState{
			Chain:    chain,
			Monitors: monitors,
			Commands: commands,
		}

		s.logger.Info("Initialized chain", "chain", chain, "monitors", len(monitors), "commands", len(commands))
	}

	return nil
}

func (s *MonitorService) Process(ready chan bool) error {
	ready <- true

	for {
		if s.paused {
			time.Sleep(time.Second)
			continue
		}

		select {
		case <-s.ctx.Done():
			return nil
		default:
		}

		anyProgress := false
		metrics := IterationMetrics{
			StartTime: time.Now(),
		}

		for _, chain := range s.chains {
			state := s.chainStates[chain]
			if state == nil {
				continue
			}

			chainProgress := s.processChain(state, &metrics)
			if chainProgress {
				anyProgress = true
			}
		}

		metrics.Duration = time.Since(metrics.StartTime)
		s.logMetrics(metrics)

		if s.shouldFailEarly(metrics) {
			return fmt.Errorf("fail-early triggered: high failure rate (%d/%d failed)", metrics.MonitorsFailed, metrics.MonitorsTotal)
		}

		if !anyProgress {
			time.Sleep(time.Duration(s.config.Sleep) * time.Second)
		}
	}
}

func (s *MonitorService) processChain(state *ChainState, metrics *IterationMetrics) bool {
	if len(state.Monitors) == 0 {
		return false
	}

	chainHead, err := s.getChainHead(state.Chain)
	if err != nil {
		s.logger.Warn("Failed to get chain head", "chain", state.Chain, "error", err)
		return false
	}

	batchCount := (len(state.Monitors) + s.config.BatchSize - 1) / s.config.BatchSize
	metrics.BatchesTotal += batchCount

	for i := 0; i < len(state.Monitors); i += s.config.BatchSize {
		end := i + s.config.BatchSize
		if end > len(state.Monitors) {
			end = len(state.Monitors)
		}
		batch := state.Monitors[i:end]

		if err := s.batchFreshen(state.Chain, batch); err != nil {
			s.logger.Warn("Batch freshen failed", "chain", state.Chain, "batch", i/s.config.BatchSize+1, "error", err)
			metrics.BatchesFailed++
		}
	}

	pool := NewWorkerPool(s.config.Concurrency, NewShellExecutor())
	pool.Start()

	for _, entry := range state.Monitors {
		firstBlock := entry.StartingBlock
		lastBlock := chainHead

		if s.config.MaxBlocksPerRun > 0 && lastBlock-firstBlock > uint64(s.config.MaxBlocksPerRun) {
			lastBlock = firstBlock + uint64(s.config.MaxBlocksPerRun)
		}

		if lastBlock <= firstBlock {
			continue
		}

		vars := TemplateVars{
			Address:    entry.Address,
			Chain:      state.Chain,
			FirstBlock: firstBlock,
			LastBlock:  lastBlock,
			BlockCount: lastBlock - firstBlock + 1,
		}

		job := MonitorJob{
			Entry:    entry,
			Commands: state.Commands,
			Vars:     vars,
		}

		pool.Submit(job)
		metrics.MonitorsTotal++
	}

	results := pool.Wait()

	for _, result := range results {
		if result.Success {
			metrics.MonitorsSuccess++
		} else {
			metrics.MonitorsFailed++
			s.logger.Warn("Monitor processing failed", "chain", state.Chain, "address", result.Address, "error", result.Error)
		}
	}

	return metrics.MonitorsTotal > 0
}

func (s *MonitorService) batchFreshen(chain string, batch []MonitorEntry) error {
	if len(batch) == 0 {
		return nil
	}

	addresses := make([]string, len(batch))
	for i, entry := range batch {
		addresses[i] = entry.Address
	}

	args := []string{"export", "--freshen"}
	args = append(args, addresses...)
	args = append(args, "--chain", chain)

	cmd := exec.CommandContext(s.ctx, "chifra", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("chifra export failed: %v: %s", err, string(output))
	}

	return nil
}

func (s *MonitorService) getChainHead(chain string) (uint64, error) {
	cmd := exec.CommandContext(s.ctx, "chifra", "blocks", "--count", "--chain", chain)
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to get chain head: %v", err)
	}

	var head uint64
	_, err = fmt.Sscanf(strings.TrimSpace(string(output)), "%d", &head)
	if err != nil {
		return 0, fmt.Errorf("failed to parse chain head: %v", err)
	}

	return head, nil
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
