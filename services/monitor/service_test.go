package monitor

import (
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestNewMonitorService(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	chains := []string{"mainnet", "sepolia"}
	config := MonitorConfig{
		WatchlistDir:    "/tmp/watchlists",
		CommandsDir:     "/tmp/commands",
		BatchSize:       8,
		Concurrency:     5,
		Sleep:           12,
		MaxBlocksPerRun: 0,
	}

	service, err := NewMonitorService(logger, chains, config)
	if err != nil {
		t.Fatalf("NewMonitorService failed: %v", err)
	}

	if service == nil {
		t.Fatal("NewMonitorService returned nil")
	}

	if service.Name() != "monitor" {
		t.Errorf("Expected name 'monitor', got '%s'", service.Name())
	}

	if len(service.chains) != 2 {
		t.Errorf("Expected 2 chains, got %d", len(service.chains))
	}

	if service.paused {
		t.Error("Expected service to start unpaused")
	}
}

func TestNewMonitorService_InvalidConfig(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	chains := []string{"mainnet"}
	config := MonitorConfig{
		WatchlistDir: "",
		CommandsDir:  "/tmp/commands",
		BatchSize:    8,
		Concurrency:  5,
		Sleep:        12,
	}

	_, err := NewMonitorService(logger, chains, config)
	if err == nil {
		t.Error("Expected error for invalid config, got nil")
	}
}

func TestMonitorService_PauseUnpause(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	chains := []string{"mainnet"}
	config := MonitorConfig{
		WatchlistDir:    "/tmp/watchlists",
		CommandsDir:     "/tmp/commands",
		BatchSize:       8,
		Concurrency:     5,
		Sleep:           12,
		MaxBlocksPerRun: 0,
	}

	service, err := NewMonitorService(logger, chains, config)
	if err != nil {
		t.Fatalf("NewMonitorService failed: %v", err)
	}

	if service.IsPaused() {
		t.Error("Expected service to start unpaused")
	}

	if !service.Pause() {
		t.Error("Pause() returned false")
	}

	if !service.IsPaused() {
		t.Error("Expected service to be paused")
	}

	if !service.Unpause() {
		t.Error("Unpause() returned false")
	}

	if service.IsPaused() {
		t.Error("Expected service to be unpaused")
	}
}

func TestMonitorService_Cleanup(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	chains := []string{"mainnet"}
	config := MonitorConfig{
		WatchlistDir:    "/tmp/watchlists",
		CommandsDir:     "/tmp/commands",
		BatchSize:       8,
		Concurrency:     5,
		Sleep:           12,
		MaxBlocksPerRun: 0,
	}

	service, err := NewMonitorService(logger, chains, config)
	if err != nil {
		t.Fatalf("NewMonitorService failed: %v", err)
	}

	service.Cleanup()

	select {
	case <-service.ctx.Done():
	case <-time.After(100 * time.Millisecond):
		t.Error("Context was not cancelled after Cleanup()")
	}
}

func TestMonitorService_ShouldFailEarly(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	chains := []string{"mainnet"}
	config := MonitorConfig{
		WatchlistDir:    "/tmp/watchlists",
		CommandsDir:     "/tmp/commands",
		BatchSize:       8,
		Concurrency:     5,
		Sleep:           12,
		MaxBlocksPerRun: 0,
	}

	service, err := NewMonitorService(logger, chains, config)
	if err != nil {
		t.Fatalf("NewMonitorService failed: %v", err)
	}

	tests := []struct {
		name     string
		metrics  IterationMetrics
		expected bool
	}{
		{
			name: "no failures",
			metrics: IterationMetrics{
				MonitorsTotal:   10,
				MonitorsSuccess: 10,
				MonitorsFailed:  0,
			},
			expected: false,
		},
		{
			name: "low failure rate",
			metrics: IterationMetrics{
				MonitorsTotal:   10,
				MonitorsSuccess: 8,
				MonitorsFailed:  2,
			},
			expected: false,
		},
		{
			name: "exactly 50% failure",
			metrics: IterationMetrics{
				MonitorsTotal:   10,
				MonitorsSuccess: 5,
				MonitorsFailed:  5,
			},
			expected: false,
		},
		{
			name: "high failure rate",
			metrics: IterationMetrics{
				MonitorsTotal:   10,
				MonitorsSuccess: 4,
				MonitorsFailed:  6,
			},
			expected: true,
		},
		{
			name: "all failures",
			metrics: IterationMetrics{
				MonitorsTotal:   10,
				MonitorsSuccess: 0,
				MonitorsFailed:  10,
			},
			expected: true,
		},
		{
			name: "no monitors",
			metrics: IterationMetrics{
				MonitorsTotal:   0,
				MonitorsSuccess: 0,
				MonitorsFailed:  0,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.shouldFailEarly(tt.metrics)
			if result != tt.expected {
				t.Errorf("Expected shouldFailEarly=%v, got %v (failed=%d, total=%d)",
					tt.expected, result, tt.metrics.MonitorsFailed, tt.metrics.MonitorsTotal)
			}
		})
	}
}

func TestMonitorService_Logger(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	chains := []string{"mainnet"}
	config := MonitorConfig{
		WatchlistDir:    "/tmp/watchlists",
		CommandsDir:     "/tmp/commands",
		BatchSize:       8,
		Concurrency:     5,
		Sleep:           12,
		MaxBlocksPerRun: 0,
	}

	service, err := NewMonitorService(logger, chains, config)
	if err != nil {
		t.Fatalf("NewMonitorService failed: %v", err)
	}

	if service.Logger() == nil {
		t.Error("Logger() returned nil")
	}

	if service.Logger() != logger {
		t.Error("Logger() did not return the same logger")
	}
}
