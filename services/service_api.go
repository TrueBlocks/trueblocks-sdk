package services

import (
	"bytes"
	"context"
	"log/slog"

	sdk "github.com/TrueBlocks/trueblocks-sdk/v6"
)

// ApiService implements Servicer and Restarter interfaces
type ApiService struct {
	logger *slog.Logger
	apiUrl string
	paused bool
	ctx    context.Context
	cancel context.CancelFunc
}

func NewApiService(logger *slog.Logger) *ApiService {
	ctx, cancel := context.WithCancel(context.Background())
	return &ApiService{
		logger: initializeLogger(logger),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *ApiService) Name() string {
	return "api"
}

func (s *ApiService) Initialize() error {
	s.apiUrl = getApiUrl()
	s.logger.Info("API service initialized.", "url", s.apiUrl)
	return nil
}

func (s *ApiService) Process(ready chan bool) error {
	s.logger.Info("API service Process() invoked.")

	go func() {
		ready <- true
		opts := sdk.DaemonOptions{
			Silent: true,
			Url:    s.apiUrl,
		}
		in := opts.ToInternal()

		daemonDone := make(chan error, 1)

		go func() {
			buffer := bytes.Buffer{}
			err := in.DaemonBytes(&buffer)
			if err != nil {
				s.logger.Error("Error running DaemonBytes", "error", err)
			}
			daemonDone <- err
		}()

		s.logger.Info("API service process running.")

		select {
		case <-s.ctx.Done():
			s.logger.Info("API service process stopping due to context cancellation.")
		case err := <-daemonDone:
			s.logger.Error("DaemonBytes exited unexpectedly", "error", err)
		}
	}()

	return nil
}

func (s *ApiService) Cleanup() {
	s.logger.Info("API service cleanup started.")
	s.cancel() // Cancel the context to signal shutdown
	s.ctx, s.cancel = context.WithCancel(context.Background())
	// for i := 0; i < 5; i++ {
	// 	s.logger.Info("Api service cleanup in progress.", "i", i)
	// 	time.Sleep(1 * time.Second)
	// }
	s.logger.Info("API service cleanup complete.")
}

func (s *ApiService) Logger() *slog.Logger {
	return s.logger
}

func (s *ApiService) ApiUrl() string {
	return s.apiUrl
}

// Pauser interface implementation (UI consistency - tracks pause state but doesn't stop service)
func (s *ApiService) IsPausable() bool {
	return true // UI shows this service as pausable
}

func (s *ApiService) IsPaused() bool {
	return s.paused
}

func (s *ApiService) Pause() bool {
	s.paused = true
	s.logger.Info("API service paused")
	return s.paused
}

func (s *ApiService) Unpause() bool {
	s.paused = false
	s.logger.Info("API service unpaused")
	return !s.paused
}

// Compile-time interface checks
var _ Servicer = (*ApiService)(nil)
var _ Restarter = (*ApiService)(nil)
var _ Pauser = (*ApiService)(nil)
