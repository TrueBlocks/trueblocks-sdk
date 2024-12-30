package services

import (
	"bytes"
	"log/slog"

	sdk "github.com/TrueBlocks/trueblocks-sdk/v4"
)

type ApiService struct {
	logger *slog.Logger
	apiUrl string
}

func NewApiService(logger *slog.Logger) *ApiService {
	return &ApiService{
		logger: initializeLogger(logger),
	}
}

func (s *ApiService) Name() string {
	return "api"
}

func (s *ApiService) Initialize() error {
	s.apiUrl = getApiUrl()
	return nil
}

func (s *ApiService) Process(ready chan bool) error {
	s.logger.Info("starting API process")
	opts := sdk.DaemonOptions{
		Silent: true,
		Url:    s.apiUrl,
	}
	in := opts.ToInternal()
	buffer := bytes.Buffer{}
	if err := in.DaemonBytes(&buffer); err != nil {
		s.logger.Error("error starting daemon", "error", err)
		return err
	}

	ready <- true
	return nil
}

func (s *ApiService) Cleanup() {
	s.logger.Info("cleaning up API Server")
}

func (s *ApiService) Logger() *slog.Logger {
	return s.logger
}

func (s *ApiService) ApiUrl() string {
	return s.apiUrl
}
