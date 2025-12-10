package monitor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseWatchlist(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		content     string
		wantCount   int
		wantErr     bool
		errContains string
	}{
		{
			name: "valid addresses only",
			content: `0x1234567890123456789012345678901234567890
0xabcdefabcdefabcdefabcdefabcdefabcdefabcd`,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "addresses with starting blocks",
			content: `0x1234567890123456789012345678901234567890, 15000000
0xabcdefabcdefabcdefabcdefabcdefabcdefabcd, 16000000`,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "with comments",
			content: `# This is a comment
0x1234567890123456789012345678901234567890
# Another comment
0xabcdefabcdefabcdefabcdefabcdefabcdefabcd  # Inline comment`,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "with empty lines",
			content: `0x1234567890123456789012345678901234567890

0xabcdefabcdefabcdefabcdefabcdefabcdefabcd

`,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:        "invalid address - no 0x prefix",
			content:     `1234567890123456789012345678901234567890`,
			wantErr:     true,
			errContains: "invalid address",
		},
		{
			name:        "invalid address - wrong length",
			content:     `0x1234`,
			wantErr:     true,
			errContains: "invalid address",
		},
		{
			name:        "invalid address - bad characters",
			content:     `0x123456789012345678901234567890123456zzzz`,
			wantErr:     true,
			errContains: "invalid address",
		},
		{
			name:        "invalid starting block",
			content:     `0x1234567890123456789012345678901234567890, abc`,
			wantErr:     true,
			errContains: "invalid starting block",
		},
		{
			name:        "empty file",
			content:     ``,
			wantErr:     true,
			errContains: "no valid addresses found",
		},
		{
			name: "only comments",
			content: `# Comment 1
# Comment 2`,
			wantErr:     true,
			errContains: "no valid addresses found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(tmpDir, tt.name+".txt")
			if err := os.WriteFile(path, []byte(tt.content), 0o644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			entries, err := ParseWatchlist(path)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errContains)
				} else if !containsString(err.Error(), tt.errContains) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if len(entries) != tt.wantCount {
					t.Errorf("Expected %d entries, got %d", tt.wantCount, len(entries))
				}
			}
		})
	}
}

func TestParseWatchlist_StartingBlock(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.txt")

	content := `0x1234567890123456789012345678901234567890
0xabcdefabcdefabcdefabcdefabcdefabcdefabcd, 15000000`

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	entries, err := ParseWatchlist(path)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(entries))
	}

	if entries[0].StartingBlock != 0 {
		t.Errorf("Expected first entry starting block 0, got %d", entries[0].StartingBlock)
	}

	if entries[1].StartingBlock != 15000000 {
		t.Errorf("Expected second entry starting block 15000000, got %d", entries[1].StartingBlock)
	}
}

func TestParseWatchlist_FileNotFound(t *testing.T) {
	_, err := ParseWatchlist("/nonexistent/path/watchlist.txt")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestParseCommands(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		content     string
		wantCount   int
		wantErr     bool
		errContains string
	}{
		{
			name: "valid commands",
			content: `commands:
  - id: export-data
    command: chifra
    arguments:
      - export
      - "{address}"
      - --cache
    output: ""
  - id: export-logs
    command: chifra
    arguments:
      - export
      - "{address}"
      - --logs
    output: "/tmp/logs.json"`,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "missing command field",
			content: `commands:
  - id: test
    arguments:
      - arg1
    output: ""`,
			wantErr:     true,
			errContains: "missing 'command' field",
		},
		{
			name: "missing arguments",
			content: `commands:
  - id: test
    command: chifra
    output: ""`,
			wantErr:     true,
			errContains: "has no arguments",
		},
		{
			name: "invalid yaml",
			content: `commands:
  - id: test
    command: chifra
  arguments:
    - this is invalid indentation`,
			wantErr:     true,
			errContains: "failed to parse YAML",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(tmpDir, tt.name+".yaml")
			if err := os.WriteFile(path, []byte(tt.content), 0o644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			commands, err := ParseCommands(path)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errContains)
				} else if !containsString(err.Error(), tt.errContains) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if len(commands) != tt.wantCount {
					t.Errorf("Expected %d commands, got %d", tt.wantCount, len(commands))
				}
			}
		})
	}
}

func TestParseCommands_FileNotFound(t *testing.T) {
	commands, err := ParseCommands("/nonexistent/path/commands.yaml")
	if err != nil {
		t.Errorf("Expected no error for missing commands file, got %v", err)
	}
	if len(commands) != 0 {
		t.Errorf("Expected empty commands list, got %d commands", len(commands))
	}
}

func TestIsValidAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{
			name:    "valid address lowercase",
			address: "0x1234567890123456789012345678901234567890",
			want:    true,
		},
		{
			name:    "valid address uppercase",
			address: "0xABCDEFABCDEFABCDEFABCDEFABCDEFABCDEFABCD",
			want:    true,
		},
		{
			name:    "valid address mixed case",
			address: "0x1234567890aBcDeF1234567890aBcDeF12345678",
			want:    true,
		},
		{
			name:    "missing 0x prefix",
			address: "1234567890123456789012345678901234567890",
			want:    false,
		},
		{
			name:    "too short",
			address: "0x1234",
			want:    false,
		},
		{
			name:    "too long",
			address: "0x123456789012345678901234567890123456789012",
			want:    false,
		},
		{
			name:    "invalid characters",
			address: "0x123456789012345678901234567890123456zzzz",
			want:    false,
		},
		{
			name:    "empty string",
			address: "",
			want:    false,
		},
		{
			name:    "only 0x",
			address: "0x",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidAddress(tt.address)
			if result != tt.want {
				t.Errorf("isValidAddress(%s) = %v, want %v", tt.address, result, tt.want)
			}
		})
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsSubstring(s, substr)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
