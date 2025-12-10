package monitor

import (
	"testing"
)

func TestNewMonitorConfig(t *testing.T) {
	cfg := NewMonitorConfig()

	if cfg.BatchSize != 8 {
		t.Errorf("Expected BatchSize 8, got %d", cfg.BatchSize)
	}
	if cfg.Concurrency < 1 {
		t.Errorf("Expected Concurrency >= 1, got %d", cfg.Concurrency)
	}
	if cfg.Sleep != 12 {
		t.Errorf("Expected Sleep 12, got %d", cfg.Sleep)
	}
	if cfg.MaxBlocksPerRun != 0 {
		t.Errorf("Expected MaxBlocksPerRun 0, got %d", cfg.MaxBlocksPerRun)
	}
}

func TestMonitorConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  MonitorConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: MonitorConfig{
				WatchlistDir:    "/path/to/watchlists",
				CommandsDir:     "/path/to/commands",
				BatchSize:       8,
				Concurrency:     5,
				Sleep:           12,
				MaxBlocksPerRun: 0,
			},
			wantErr: false,
		},
		{
			name: "empty watchlist dir",
			config: MonitorConfig{
				CommandsDir:     "/path/to/commands",
				BatchSize:       8,
				Concurrency:     5,
				Sleep:           12,
				MaxBlocksPerRun: 0,
			},
			wantErr: true,
		},
		{
			name: "empty commands dir",
			config: MonitorConfig{
				WatchlistDir:    "/path/to/watchlists",
				BatchSize:       8,
				Concurrency:     5,
				Sleep:           12,
				MaxBlocksPerRun: 0,
			},
			wantErr: true,
		},
		{
			name: "invalid batch size",
			config: MonitorConfig{
				WatchlistDir:    "/path/to/watchlists",
				CommandsDir:     "/path/to/commands",
				BatchSize:       0,
				Concurrency:     5,
				Sleep:           12,
				MaxBlocksPerRun: 0,
			},
			wantErr: true,
		},
		{
			name: "invalid concurrency",
			config: MonitorConfig{
				WatchlistDir:    "/path/to/watchlists",
				CommandsDir:     "/path/to/commands",
				BatchSize:       8,
				Concurrency:     0,
				Sleep:           12,
				MaxBlocksPerRun: 0,
			},
			wantErr: true,
		},
		{
			name: "negative sleep",
			config: MonitorConfig{
				WatchlistDir:    "/path/to/watchlists",
				CommandsDir:     "/path/to/commands",
				BatchSize:       8,
				Concurrency:     5,
				Sleep:           -1,
				MaxBlocksPerRun: 0,
			},
			wantErr: true,
		},
		{
			name: "negative max blocks",
			config: MonitorConfig{
				WatchlistDir:    "/path/to/watchlists",
				CommandsDir:     "/path/to/commands",
				BatchSize:       8,
				Concurrency:     5,
				Sleep:           12,
				MaxBlocksPerRun: -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr && err == nil {
				t.Error("Expected error, got nil")
			} else if !tt.wantErr && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}
