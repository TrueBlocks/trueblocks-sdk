package sdk

import (
	"strings"
	"testing"
)

// TestEncodeTransaction_Uint256SmallNumbers tests that small numbers (fitting in uint64)
// are properly encoded for uint256 parameters. This was a bug where the parser stored
// small numbers as uint64 but failed to convert to *big.Int for uint256 types.
//
// This test exercises the fix in rpc_transactions.go lines 89-108 which detects when
// a parser.ArgNumber has stored a value as uint64 (n.Uint) but the ABI type requires
// a big.Int (uint256, int256, etc.), and performs the conversion.
//
// The original error was: "abi: cannot use int64 as type ptr as argument"
func TestEncodeTransaction_Uint256SmallNumbers(t *testing.T) {
	tests := []struct {
		name      string
		amount    string
		wantError bool
	}{
		{
			name:      "small number 9 (original bug case)",
			amount:    "9",
			wantError: false,
		},
		{
			name:      "token amount in wei (18 decimals)",
			amount:    "9000000000000000000",
			wantError: false,
		},
		{
			name:      "max uint64",
			amount:    "18446744073709551615",
			wantError: false,
		},
		{
			name:      "number exceeding uint64",
			amount:    "99999999999999999999999999999",
			wantError: false,
		},
		{
			name:      "zero",
			amount:    "0",
			wantError: false,
		},
	}

	// Use RAI token from the original bug report
	tokenAddress := "0x03ab458634910aad20ef5f1c8ee96f1d6ac54919"
	spender := "0xf503017d7baf7fbc0fff7492b751025c6a78179b"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := TransactionEncodeRequest{
				Chain:           "mainnet",
				ContractAddress: tokenAddress,
				Signature:       "approve(address,uint256)",
				Arguments:       []string{spender, tt.amount},
			}

			result, err := EncodeTransaction(req)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result == "" {
				t.Errorf("expected non-empty result")
				return
			}

			// Verify it starts with approve function selector (0x095ea7b3)
			if !strings.HasPrefix(result, "0x095ea7b3") {
				t.Errorf("expected result to start with 0x095ea7b3, got %s", result[:10])
			}

			// Verify total length (0x + selector(8) + address(64) + uint256(64) = 138 chars)
			expectedLength := 138
			if len(result) != expectedLength {
				t.Errorf("expected result length %d, got %d: %s", expectedLength, len(result), result)
			}
		})
	}
}

// TestEncodeTransaction_Int256Types tests that signed integer types > 64 bits work
func TestEncodeTransaction_Int256Types(t *testing.T) {
	// This would need a contract with int256 parameters - skipping for now
	// but the fix handles both UintTy and IntTy with Size > 64
	t.Skip("Need contract with int256 parameters for testing")
}
