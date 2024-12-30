package services

import (
	"io"
	"log/slog"
	"time"
)

type MonitorService struct {
	logger *slog.Logger
}

// NewMonitorService creates a new instance of MonitorService.
func NewMonitorService(logger *slog.Logger) *MonitorService {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	return &MonitorService{
		logger: logger,
	}
}

// Name returns the name of the service.
func (s *MonitorService) Name() string {
	return "monitor"
}

// Initialize performs any setup required for the MonitorService.
func (s *MonitorService) Initialize() error {
	s.logger.Info("Monitor service initialized.")
	return nil
}

// Process starts the monitor logic in a goroutine and signals readiness.
func (s *MonitorService) Process(ready chan bool) error {
	ready <- true // Signal that the service is ready.
	s.logger.Info("Monitor service started.")
	go func() {
		for {
			s.logger.Info("Monitor is running...")
			time.Sleep(3 * time.Second)
		}
	}()
	return nil
}

// Cleanup stops the MonitorService.
func (s *MonitorService) Cleanup() {
	s.logger.Info("Monitor service cleanup complete.")
}

// Logger returns the logger for this service.
func (s *MonitorService) Logger() *slog.Logger {
	return s.logger
}
