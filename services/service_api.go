package services

import (
	"bytes"
	"context"
	"log/slog"

	sdk "github.com/TrueBlocks/trueblocks-sdk/v5"
)

type ApiService struct {
	logger *slog.Logger
	apiUrl string
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
