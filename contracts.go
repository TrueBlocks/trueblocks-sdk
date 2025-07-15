// Copyright 2016, 2026 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package sdk

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// CreateContracts returns mock contract data for testing and development
func CreateContracts() []*types.Contract {
	// Panvala PAN - ERC20 Token
	panvalaAbi := &types.Abi{
		Functions: []types.Function{
			{
				Name:            "name",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string", Value: "Panvala pan"},
				},
			},
			{
				Name:            "symbol",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string", Value: "PAN"},
				},
			},
			{
				Name:            "decimals",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint8", Value: "18"},
				},
			},
			{
				Name:            "totalSupply",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
			{
				Name:            "balanceOf",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "account", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
			{
				Name:            "approve",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "spender", ParameterType: "address"},
					{Name: "value", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool", Value: nil},
				},
			},
			{
				Name:            "transfer",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "recipient", ParameterType: "address"},
					{Name: "amount", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool", Value: nil},
				},
			},
			{
				Name:            "transferFrom",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "sender", ParameterType: "address"},
					{Name: "recipient", ParameterType: "address"},
					{Name: "amount", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool", Value: nil},
				},
			},
			{
				Name:            "allowance",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "owner", ParameterType: "address"},
					{Name: "spender", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
		},
	}

	// zkSync - Simple contract with AbiNotFound
	zkSyncAbi := &types.Abi{
		Functions: []types.Function{
			{
				Name:            "AbiNotFound",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs:         []types.Parameter{},
			},
		},
	}

	// USD Coin - Proxy contract
	usdcAbi := &types.Abi{
		Functions: []types.Function{
			{
				Name:            "implementation",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "address", Value: nil},
				},
			},
			{
				Name:            "admin",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "address", Value: nil},
				},
			},
			{
				Name:            "upgradeTo",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "newImplementation", ParameterType: "address"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "upgradeToAndCall",
				FunctionType:    "function",
				StateMutability: "payable",
				Inputs: []types.Parameter{
					{Name: "newImplementation", ParameterType: "address"},
					{Name: "data", ParameterType: "bytes"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "changeAdmin",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "newAdmin", ParameterType: "address"},
				},
				Outputs: []types.Parameter{},
			},
		},
	}

	// Tether USDT - Complex token with admin functions
	usdtAbi := &types.Abi{
		Functions: []types.Function{
			{
				Name:            "name",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string", Value: "Tether USD"},
				},
			},
			{
				Name:            "symbol",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string", Value: "USDT"},
				},
			},
			{
				Name:            "decimals",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: "6"},
				},
			},
			{
				Name:            "totalSupply",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
			{
				Name:            "balanceOf",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "who", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
			{
				Name:            "owner",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "address", Value: nil},
				},
			},
			{
				Name:            "paused",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool", Value: nil},
				},
			},
			{
				Name:            "approve",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "_spender", ParameterType: "address"},
					{Name: "_value", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "transfer",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "_to", ParameterType: "address"},
					{Name: "_value", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "transferFrom",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "_from", ParameterType: "address"},
					{Name: "_to", ParameterType: "address"},
					{Name: "_value", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
		},
	}

	// Unchained Index
	unchainedIndexAbi := &types.Abi{
		Functions: []types.Function{
			{
				Name:            "publishHash",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "chain", ParameterType: "string"},
					{Name: "hash", ParameterType: "string"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "changeOwner",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "newOwner", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "oldOwner", ParameterType: "address"},
				},
			},
			{
				Name:            "donate",
				FunctionType:    "function",
				StateMutability: "payable",
				Inputs:          []types.Parameter{},
				Outputs:         []types.Parameter{},
			},
			{
				Name:            "owner",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "address", Value: nil},
				},
			},
		},
	}

	// Gitcoin GTC - ERC20 Token
	gitcoinAbi := &types.Abi{
		Functions: []types.Function{
			{
				Name:            "name",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string", Value: "Gitcoin"},
				},
			},
			{
				Name:            "symbol",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string", Value: "GTC"},
				},
			},
			{
				Name:            "decimals",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint8", Value: "18"},
				},
			},
			{
				Name:            "totalSupply",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
			{
				Name:            "balanceOf",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "account", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
		},
	}

	// DAI Stablecoin - ERC20 Token
	daiAbi := &types.Abi{
		Functions: []types.Function{
			{
				Name:            "name",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string", Value: "Dai Stablecoin"},
				},
			},
			{
				Name:            "symbol",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string", Value: "DAI"},
				},
			},
			{
				Name:            "decimals",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: "18"},
				},
			},
			{
				Name:            "totalSupply",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
			{
				Name:            "balanceOf",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "src", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256", Value: nil},
				},
			},
		},
	}

	// GitCoin Grant - Simple contract with AbiNotFound
	gitcoinGrantAbi := &types.Abi{
		Functions: []types.Function{
			{
				Name:            "AbiNotFound",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs:         []types.Parameter{},
			},
		},
	}

	return []*types.Contract{
		{
			Address:     base.HexToAddress("0xd56dac73a4d6766464b38ec6d91eb45ce7457c44"),
			Name:        "Panvala pan PAN",
			Abi:         panvalaAbi,
			Date:        "2024-01-15",
			LastUpdated: 1705401600,
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{
				"name":     "Panvala pan",
				"symbol":   "PAN",
				"decimals": "18",
			},
		},
		{
			Address:     base.HexToAddress("0xabea9132b05a70803a4e85094fd0e1800777fbef"),
			Name:        "zkSync",
			Abi:         zkSyncAbi,
			Date:        "2024-01-16",
			LastUpdated: 1705488000,
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{},
		},
		{
			Address:     base.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
			Name:        "USD Coin (L1) USDC",
			Abi:         usdcAbi,
			Date:        "2024-01-17",
			LastUpdated: 1705574400,
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{},
		},
		{
			Address:     base.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
			Name:        "Tether USD (L1) USDT",
			Abi:         usdtAbi,
			Date:        "2024-01-18",
			LastUpdated: 1705660800,
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{
				"name":     "Tether USD",
				"symbol":   "USDT",
				"decimals": "6",
			},
		},
		{
			Address:     base.HexToAddress("0x0c316b7042b419d07d343f2f4f5bd54ff731183d"),
			Name:        "Unchained Index (v1.0)",
			Abi:         unchainedIndexAbi,
			Date:        "2024-01-19",
			LastUpdated: 1705747200,
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{},
		},
		{
			Address:     base.HexToAddress("0xde30da39c46104798bb5aa3fe8b9e0e1f348163f"),
			Name:        "Gitcoin (L1) GTC",
			Abi:         gitcoinAbi,
			Date:        "2024-01-20",
			LastUpdated: 1705833600,
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{
				"name":     "Gitcoin",
				"symbol":   "GTC",
				"decimals": "18",
			},
		},
		{
			Address:     base.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
			Name:        "Dai Stablecoin (L1) DAI",
			Abi:         daiAbi,
			Date:        "2024-01-21",
			LastUpdated: 1705920000,
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{
				"name":     "Dai Stablecoin",
				"symbol":   "DAI",
				"decimals": "18",
			},
		},
		{
			Address:     base.HexToAddress("0x7d655c57f71464b6f83811c55d84009cd9f5221c"),
			Name:        "GitCoin Grant 6,7,8",
			Abi:         gitcoinGrantAbi,
			Date:        "2024-01-22",
			LastUpdated: 1706006400,
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{},
		},
	}
}
