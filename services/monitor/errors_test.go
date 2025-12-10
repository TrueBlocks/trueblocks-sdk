package monitor

import (
	"errors"
	"testing"
)

func TestConfigError(t *testing.T) {
	err := NewConfigError("WatchlistDir", "cannot be empty")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	var configErr *ConfigError
	if !errors.As(err, &configErr) {
		t.Error("Expected error to be ConfigError type")
	}
	if configErr.Field != "WatchlistDir" {
		t.Errorf("Expected field 'WatchlistDir', got '%s'", configErr.Field)
	}
}

func TestWatchlistError(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		line    int
		message string
	}{
		{
			name:    "with line number",
			path:    "/path/to/watchlist.txt",
			line:    5,
			message: "invalid address",
		},
		{
			name:    "without line number",
			path:    "/path/to/watchlist.txt",
			line:    0,
			message: "file not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewWatchlistError(tt.path, tt.line, tt.message)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			var watchlistErr *WatchlistError
			if !errors.As(err, &watchlistErr) {
				t.Error("Expected error to be WatchlistError type")
			}
			if watchlistErr.Path != tt.path {
				t.Errorf("Expected path '%s', got '%s'", tt.path, watchlistErr.Path)
			}
			if watchlistErr.Line != tt.line {
				t.Errorf("Expected line %d, got %d", tt.line, watchlistErr.Line)
			}
		})
	}
}

func TestCommandError(t *testing.T) {
	tests := []struct {
		name      string
		commandID string
		address   string
		message   string
		cause     error
	}{
		{
			name:      "without cause",
			commandID: "export-data",
			address:   "0x1234",
			message:   "execution failed",
			cause:     nil,
		},
		{
			name:      "with cause",
			commandID: "export-logs",
			address:   "0x5678",
			message:   "chifra failed",
			cause:     errors.New("RPC timeout"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewCommandError(tt.commandID, tt.address, tt.message, tt.cause)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			var cmdErr *CommandError
			if !errors.As(err, &cmdErr) {
				t.Error("Expected error to be CommandError type")
			}
			if cmdErr.CommandID != tt.commandID {
				t.Errorf("Expected commandID '%s', got '%s'", tt.commandID, cmdErr.CommandID)
			}
			if cmdErr.Address != tt.address {
				t.Errorf("Expected address '%s', got '%s'", tt.address, cmdErr.Address)
			}

			if tt.cause != nil {
				if !errors.Is(err, tt.cause) {
					t.Error("Expected error chain to contain cause")
				}
			}
		})
	}
}

func TestCommandError_Unwrap(t *testing.T) {
	cause := errors.New("underlying error")
	err := NewCommandError("cmd1", "0xabc", "failed", cause)

	var cmdErr *CommandError
	if !errors.As(err, &cmdErr) {
		t.Fatal("Expected CommandError type")
	}

	unwrapped := cmdErr.Unwrap()
	if unwrapped != cause {
		t.Errorf("Expected unwrapped error to be cause, got %v", unwrapped)
	}
}
