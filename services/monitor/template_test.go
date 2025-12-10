package monitor

import (
	"testing"
)

func TestExpandTemplate(t *testing.T) {
	vars := TemplateVars{
		Address:    "0x1234567890abcdef",
		Chain:      "mainnet",
		FirstBlock: 15000000,
		LastBlock:  15000100,
		BlockCount: 101,
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "all variables",
			input:    "{address} on {chain} from {first_block} to {last_block} ({block_count} blocks)",
			expected: "0x1234567890abcdef on mainnet from 15000000 to 15000100 (101 blocks)",
		},
		{
			name:     "address only",
			input:    "Address: {address}",
			expected: "Address: 0x1234567890abcdef",
		},
		{
			name:     "chain only",
			input:    "Chain: {chain}",
			expected: "Chain: mainnet",
		},
		{
			name:     "blocks only",
			input:    "{first_block}-{last_block}",
			expected: "15000000-15000100",
		},
		{
			name:     "block count only",
			input:    "Count: {block_count}",
			expected: "Count: 101",
		},
		{
			name:     "no variables",
			input:    "static string",
			expected: "static string",
		},
		{
			name:     "repeated variables",
			input:    "{address} {address} {address}",
			expected: "0x1234567890abcdef 0x1234567890abcdef 0x1234567890abcdef",
		},
		{
			name:     "file path template",
			input:    "~/exports/{chain}/transactions/{address}.csv",
			expected: "~/exports/mainnet/transactions/0x1234567890abcdef.csv",
		},
		{
			name:     "command arguments",
			input:    "--first_block {first_block} --last_block {last_block}",
			expected: "--first_block 15000000 --last_block 15000100",
		},
		{
			name:     "unknown variable preserved",
			input:    "{unknown_var} {address}",
			expected: "{unknown_var} 0x1234567890abcdef",
		},
		{
			name:     "partial braces",
			input:    "{address incomplete",
			expected: "{address incomplete",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExpandTemplate(tt.input, vars)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestExpandTemplate_ZeroValues(t *testing.T) {
	vars := TemplateVars{
		Address:    "",
		Chain:      "",
		FirstBlock: 0,
		LastBlock:  0,
		BlockCount: 0,
	}

	input := "{address}|{chain}|{first_block}|{last_block}|{block_count}"
	expected := "||0|0|0"
	result := ExpandTemplate(input, vars)

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFormatUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected string
	}{
		{
			name:     "zero",
			input:    0,
			expected: "0",
		},
		{
			name:     "single digit",
			input:    5,
			expected: "5",
		},
		{
			name:     "multiple digits",
			input:    12345,
			expected: "12345",
		},
		{
			name:     "large number",
			input:    15000000,
			expected: "15000000",
		},
		{
			name:     "max uint64",
			input:    18446744073709551615,
			expected: "18446744073709551615",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatUint64(tt.input)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
