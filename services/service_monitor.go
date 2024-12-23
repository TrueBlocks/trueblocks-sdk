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
func (m *MonitorService) Name() string {
	return "Monitor Service"
}

// Initialize performs any setup required for the MonitorService.
func (m *MonitorService) Initialize() error {
	m.logger.Info("Monitor service initialized.")
	return nil
}

// Process starts the monitor logic in a goroutine and signals readiness.
func (m *MonitorService) Process(ready chan bool) error {
	ready <- true // Signal that the service is ready.
	m.logger.Info("Monitor service started.")
	go func() {
		for {
			m.logger.Info("Monitor is running...")
			time.Sleep(3 * time.Second)
		}
	}()
	return nil
}

// Cleanup stops the MonitorService.
func (m *MonitorService) Cleanup() {
	m.logger.Info("Monitor service cleanup complete.")
}

// Logger returns the logger for this service.
func (m *MonitorService) Logger() *slog.Logger {
	return m.logger
}
