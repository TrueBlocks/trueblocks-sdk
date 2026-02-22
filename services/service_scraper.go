package services

import (
	"context"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/logger"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/output"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v6"
)

// ScrapeCompletedEvent is sent when scraper completes a batch
// This is a "wake up" signal - monitors track their own block state
type ScrapeCompletedEvent struct {
	Chain string
	Meta  interface{} // *types.MetaData from core
}

// ScrapeService implements Servicer, Pauser, and Restarter interfaces
type ScrapeService struct {
	paused           bool
	logger           *slog.Logger
	initMode         string
	configTargets    []string
	sleep            int
	blockCnt         int
	ctx              context.Context
	cancel           context.CancelFunc
	onScrapeComplete func(ScrapeCompletedEvent)
}

func NewScrapeService(logger *slog.Logger, initMode string, configTargets []string, sleep int, blockCnt int) *ScrapeService {
	ctx, cancel := context.WithCancel(context.Background())
	return &ScrapeService{
		logger:        initializeLogger(logger),
		initMode:      initMode,
		configTargets: configTargets,
		sleep:         sleep,
		blockCnt:      blockCnt,
		ctx:           ctx,
		cancel:        cancel,
	}
}

func (s *ScrapeService) Name() string {
	return "scraper"
}

func (s *ScrapeService) Initialize() error {
	s.logger.Info("Scraper initialization started")
	s.logger.Info("Initializing unchained index")

	if s.initMode != "none" {
		reports := make([]*scraperReport, 0, len(s.configTargets))
		for _, chain := range s.configTargets {
			if rep, err := s.initOneChain(chain); err != nil {
				if !strings.HasPrefix(err.Error(), "no record found in the Unchained Index") {
					s.logger.Warn("Warning", "msg", err)
				} else {
					s.logger.Warn("No record found in the Unchained Index for chain", "chain", chain)
				}
			} else {
				reports = append(reports, rep)
			}
		}

		for _, report := range reports {
			reportScrape(s.logger, report)
		}
	}
	s.logger.Info("Scraper initialization complete")
	return nil
}

func (s *ScrapeService) Process(ready chan bool) error {
	s.logger.Info("Starting scraper process", "sleep", s.sleep, "targets", s.configTargets)

	ready <- true
	runCount := 0

	for {
		select {
		case <-s.ctx.Done():
			s.logger.Info("Scrape service process stopping due to context cancellation")
			return nil
		default:
			if !s.IsPaused() {
				caughtUp := true
				for _, chain := range s.configTargets {
					if report, err := s.scrapeOneChain(chain); err != nil {
						s.logger.Warn("Error scraping chain", "chain", chain, "error", err)
						time.Sleep(1 * time.Second)
						continue
					} else if report == nil { // we may be paused
						continue
					} else {
						if report.Staged > (28 + 4) {
							caughtUp = false
						}
					}
				}
				if caughtUp {
					if runCount%5 == 0 || s.sleep > 10 {
						s.logger.Info("All chains caught up")
					}
					runCount++
					time.Sleep(time.Duration(s.sleep) * time.Second)
				} else {
					time.Sleep(1 * time.Second)
				}
			} else {
				time.Sleep(2 * time.Second)
			}
		}
	}
}

func (s *ScrapeService) IsPausable() bool {
	return true
}

func (s *ScrapeService) IsPaused() bool {
	return s.paused
}

func (s *ScrapeService) Pause() bool {
	s.paused = true
	return s.paused
}

func (s *ScrapeService) Unpause() bool {
	s.paused = false
	return s.paused
}

func (s *ScrapeService) Cleanup() {
	s.cancel()
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.logger.Info("Scraper service cleanup completed")
}

func (s *ScrapeService) Logger() *slog.Logger {
	return s.logger
}

// SetScrapeCompleteCallback configures the scraper to call a function when scraping completes
func (s *ScrapeService) SetScrapeCompleteCallback(fn func(ScrapeCompletedEvent)) {
	s.onScrapeComplete = fn
}

func (s *ScrapeService) initOneChain(chain string) (*scraperReport, error) {
	defer func() {
		logger.SetLoggerWriter(io.Discard)
		_ = os.Setenv("TB_SCRAPE_HEADLESS", "")
	}()
	logger.SetLoggerWriter(os.Stderr)
	_ = os.Setenv("TB_SCRAPE_HEADLESS", "true")

	opts := sdk.InitOptions{
		Globals: sdk.Globals{
			Chain: chain,
		},
	}

	var err error
	var meta *types.MetaData
	switch s.initMode {
	case "all":
		_, meta, err = opts.InitAll()
	case "blooms":
		_, meta, err = opts.Init()
	}

	if err != nil {
		logger.Warn("Error during initialization", "error", err)
		return nil, err
	}

	// Generate a scraper report using the metadata
	report := reportScrapeRun(meta, chain, s.blockCnt)
	logger.Info("Initialization completed", "report", report)

	return report, nil
}

func (s *ScrapeService) scrapeOneChain(chain string) (*scraperReport, error) {
	defer func() {
		logger.SetLoggerWriter(io.Discard)
		_ = os.Setenv("TB_SCRAPE_HEADLESS", "")
	}()
	logger.SetLoggerWriter(os.Stderr)
	_ = os.Setenv("TB_SCRAPE_HEADLESS", "true")

	if s.IsPaused() {
		s.logger.Debug("Scraper is paused, skipping scraping step")
		return nil, nil
	}

	// Create streaming context for event handling
	rCtx := output.NewStreamingContext()
	done := make(chan struct{})

	// Start listener goroutine for streaming events
	go func() {
		defer close(done)
		for {
			select {
			case event, ok := <-rCtx.ModelChan:
				if !ok {
					return
				}
				s.handleStreamEvent(event)
			case err, ok := <-rCtx.ErrorChan:
				if !ok {
					return
				}
				s.logger.Error("Scrape error", "error", err)
			case <-rCtx.Ctx.Done():
				return
			}
		}
	}()

	opts := sdk.ScrapeOptions{
		BlockCnt: uint64(s.blockCnt),
		Globals: sdk.Globals{
			Chain: chain,
		},
		RenderCtx: rCtx,
	}

	msg, meta, err := opts.ScrapeRunOnce()

	// Cleanup: close channels and wait for goroutine
	close(rCtx.ModelChan)
	close(rCtx.ErrorChan)
	<-done

	if err != nil {
		return nil, err
	}

	if len(msg) > 0 {
		s.logger.Info(msg[0].String())
	}
	return reportScrapeRun(meta, chain, s.blockCnt), nil
}

// handleStreamEvent processes events from the streaming context
func (s *ScrapeService) handleStreamEvent(event types.Modeler) {
	switch e := event.(type) {
	case *types.Message:
		s.logger.Info(e.Msg)
	default:
		if s.onScrapeComplete != nil {
			m := event.Model("", "", false, nil)
			if chainVal, ok := m.Data["chain"]; ok {
				if chainStr, ok := chainVal.(string); ok {
					go func(chain string) {
						defer func() {
							if r := recover(); r != nil {
								s.logger.Error("Callback panic", "error", r)
							}
						}()
						s.onScrapeComplete(ScrapeCompletedEvent{Chain: chain})
					}(chainStr)
				}
			}
		}
	}
}

// Compile-time interface checks
var (
	_ Servicer  = (*ScrapeService)(nil)
	_ Restarter = (*ScrapeService)(nil)
	_ Pauser    = (*ScrapeService)(nil)
)
