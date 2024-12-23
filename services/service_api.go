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

func (a *ApiService) Name() string {
	return "API Server"
}

func (a *ApiService) Initialize() error {
	a.apiUrl = getApiUrl()
	return nil
}

func (a *ApiService) Process(ready chan bool) error {
	a.logger.Info("starting API process")
	opts := sdk.DaemonOptions{
		Silent: true,
		Url:    a.apiUrl,
	}
	in := opts.ToInternal()
	buffer := bytes.Buffer{}
	if err := in.DaemonBytes(&buffer); err != nil {
		a.logger.Error("error starting daemon", "error", err)
		return err
	}

	ready <- true
	return nil
}

func (a *ApiService) Cleanup() {
	a.logger.Info("cleaning up API Server")
}

func (a *ApiService) Logger() *slog.Logger {
	return a.logger
}

func (a *ApiService) ApiUrl() string {
	return a.apiUrl
}
