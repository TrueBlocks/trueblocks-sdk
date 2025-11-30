package sdk

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	chifraAbi "github.com/TrueBlocks/trueblocks-chifra/v6/pkg/abi"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/base"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/parser"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/rpc"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/types"
)

// TransactionEncodeRequest represents a request to encode a contract function call
type TransactionEncodeRequest struct {
	Chain           string   // Chain name (e.g., "mainnet", "sepolia")
	ContractAddress string   // Contract address
	Signature       string   // Function signature (e.g., "approve(address,uint256)")
	Arguments       []string // String arguments to encode
}

// EncodeTransaction encodes a contract function call into calldata hex string
func EncodeTransaction(req TransactionEncodeRequest) (string, error) {
	// Create RPC connection
	conn := rpc.TempConnection(req.Chain)
	contractAddr := base.HexToAddress(req.ContractAddress)

	// Extract function name from signature
	functionName := req.Signature
	if idx := strings.Index(req.Signature, "("); idx > 0 {
		functionName = req.Signature[:idx]
	}

	// Build call string with arguments
	callStr := fmt.Sprintf("%s(%s)", functionName, strings.Join(req.Arguments, ","))

	// Parse the call using chifra's parser
	parsed, err := parser.ParseCall(callStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse call: %w", err)
	}

	// Load ABI for the contract
	abiMap := &chifraAbi.SelectorSyncMap{}
	if err := chifraAbi.LoadAbi(conn, contractAddr, abiMap); err != nil {
		return "", fmt.Errorf("failed to load ABI for contract %s: %w", req.ContractAddress, err)
	}

	// Find the function in the ABI
	var function *types.Function
	var callArguments []*parser.ContractArgument

	if parsed.FunctionNameCall != nil {
		callArguments = parsed.FunctionNameCall.Arguments

		// Find function by name
		abiMap.Range(func(k any, v any) bool {
			fn := v.(*types.Function)
			if fn.Name == functionName {
				function = fn
				return false // stop iteration
			}
			return true // continue
		})
	}

	if function == nil {
		return "", fmt.Errorf("function %s not found in ABI for contract %s", functionName, req.ContractAddress)
	}

	// Get ABI method for type checking
	abiMethod, err := function.GetAbiMethod()
	if err != nil {
		return "", fmt.Errorf("failed to get ABI method: %w", err)
	}

	// Validate argument count
	if len(abiMethod.Inputs) != len(callArguments) {
		return "", fmt.Errorf("expected %d arguments, got %d", len(abiMethod.Inputs), len(callArguments))
	}

	// Convert arguments to proper types
	args := make([]any, len(callArguments))
	for i, arg := range callArguments {
		// Workaround for parser bug: numbers requiring big.Int aren't converted properly when stored as uint64/int64
		inputType := &abiMethod.Inputs[i].Type
		if (inputType.T == abi.UintTy || inputType.T == abi.IntTy) && inputType.Size > 64 && arg.Number != nil {
			// ABI type requires *big.Int but parser may have stored as uint64/int64 - force conversion
			if arg.Number.Big != nil {
				// Already a big number - convert base.Wei to big.Int
				args[i] = arg.Number.Big.BigInt()
			} else if arg.Number.Uint != nil {
				// Small number stored as uint64 - convert to *big.Int
				bigInt := base.NewWei(0)
				bigInt.SetUint64(*arg.Number.Uint)
				args[i] = bigInt.BigInt()
			} else if arg.Number.Int != nil {
				// Signed number - convert to *big.Int
				bigInt := base.NewWei(0)
				bigInt.SetInt64(*arg.Number.Int)
				args[i] = bigInt.BigInt()
			} else {
				return "", fmt.Errorf("argument %d: number has no value", i)
			}
		} else {
			converted, err := arg.AbiType(&abiMethod.Inputs[i].Type)
			if err != nil {
				return "", fmt.Errorf("failed to convert argument %d: %w", i, err)
			}
			args[i] = converted
		}
	}

	// Pack the function call (selector + encoded arguments)
	calldata, err := function.Pack(args)
	if err != nil {
		return "", fmt.Errorf("failed to encode transaction: %w", err)
	}

	// Return hex-encoded calldata
	return hexutil.Encode(calldata), nil
}
