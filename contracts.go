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
	mockERC20Abi := &types.Abi{
		Functions: []types.Function{
			// Write functions (zero-parameter first)
			{
				Name:            "pause",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs:         []types.Parameter{},
			},
			{
				Name:            "unpause",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs:          []types.Parameter{},
				Outputs:         []types.Parameter{},
			},
			{
				Name:            "transfer",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "to", ParameterType: "address"},
					{Name: "amount", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool"},
				},
			},
			{
				Name:            "transferFrom",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "from", ParameterType: "address"},
					{Name: "to", ParameterType: "address"},
					{Name: "amount", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool"},
				},
			},
			{
				Name:            "approve",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "spender", ParameterType: "address"},
					{Name: "amount", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool"},
				},
			},
			{
				Name:            "mint",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "to", ParameterType: "address"},
					{Name: "amount", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "burn",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "amount", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
			// Read functions
			{
				Name:            "name",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string"},
				},
			},
			{
				Name:            "symbol",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string"},
				},
			},
			{
				Name:            "decimals",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint8"},
				},
			},
			{
				Name:            "totalSupply",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
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
					{Name: "", ParameterType: "uint256"},
				},
			},
			{
				Name:            "allowance",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "owner", ParameterType: "address"},
					{Name: "spender", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
				},
			},
		},
	}

	mockERC721Abi := &types.Abi{
		Functions: []types.Function{
			// Read functions
			{
				Name:            "name",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string"},
				},
			},
			{
				Name:            "symbol",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string"},
				},
			},
			{
				Name:            "tokenURI",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "tokenId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "string"},
				},
			},
			{
				Name:            "ownerOf",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "tokenId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "address"},
				},
			},
			{
				Name:            "balanceOf",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "owner", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
				},
			},
			{
				Name:            "getApproved",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "tokenId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "address"},
				},
			},
			{
				Name:            "isApprovedForAll",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "owner", ParameterType: "address"},
					{Name: "operator", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool"},
				},
			},
			// Write functions
			{
				Name:            "approve",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "to", ParameterType: "address"},
					{Name: "tokenId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "setApprovalForAll",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "operator", ParameterType: "address"},
					{Name: "approved", ParameterType: "bool"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "transferFrom",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "from", ParameterType: "address"},
					{Name: "to", ParameterType: "address"},
					{Name: "tokenId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "safeTransferFrom",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "from", ParameterType: "address"},
					{Name: "to", ParameterType: "address"},
					{Name: "tokenId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "safeTransferFrom",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "from", ParameterType: "address"},
					{Name: "to", ParameterType: "address"},
					{Name: "tokenId", ParameterType: "uint256"},
					{Name: "data", ParameterType: "bytes"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "mint",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "to", ParameterType: "address"},
					{Name: "tokenId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "burn",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "tokenId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{},
			},
		},
	}

	mockGovernanceAbi := &types.Abi{
		Functions: []types.Function{
			// Read functions
			{
				Name:            "proposalCount",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs:          []types.Parameter{},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
				},
			},
			{
				Name:            "getProposal",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "proposalId", ParameterType: "uint256"},
				},
				Outputs: []types.Parameter{
					{Name: "id", ParameterType: "uint256"},
					{Name: "proposer", ParameterType: "address"},
					{Name: "description", ParameterType: "string"},
					{Name: "startTime", ParameterType: "uint256"},
					{Name: "endTime", ParameterType: "uint256"},
					{Name: "executed", ParameterType: "bool"},
				},
			},
			{
				Name:            "hasVoted",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "proposalId", ParameterType: "uint256"},
					{Name: "voter", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "bool"},
				},
			},
			{
				Name:            "getVotingPower",
				FunctionType:    "function",
				StateMutability: "view",
				Inputs: []types.Parameter{
					{Name: "account", ParameterType: "address"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
				},
			},
			// Write functions
			{
				Name:            "propose",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "targets", ParameterType: "address[]"},
					{Name: "values", ParameterType: "uint256[]"},
					{Name: "calldatas", ParameterType: "bytes[]"},
					{Name: "description", ParameterType: "string"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
				},
			},
			{
				Name:            "castVote",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "proposalId", ParameterType: "uint256"},
					{Name: "support", ParameterType: "uint8"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
				},
			},
			{
				Name:            "castVoteWithReason",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "proposalId", ParameterType: "uint256"},
					{Name: "support", ParameterType: "uint8"},
					{Name: "reason", ParameterType: "string"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
				},
			},
			{
				Name:            "execute",
				FunctionType:    "function",
				StateMutability: "payable",
				Inputs: []types.Parameter{
					{Name: "targets", ParameterType: "address[]"},
					{Name: "values", ParameterType: "uint256[]"},
					{Name: "calldatas", ParameterType: "bytes[]"},
					{Name: "descriptionHash", ParameterType: "bytes32"},
				},
				Outputs: []types.Parameter{
					{Name: "", ParameterType: "uint256"},
				},
			},
			{
				Name:            "delegate",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "delegatee", ParameterType: "address"},
				},
				Outputs: []types.Parameter{},
			},
			{
				Name:            "delegateBySig",
				FunctionType:    "function",
				StateMutability: "nonpayable",
				Inputs: []types.Parameter{
					{Name: "delegatee", ParameterType: "address"},
					{Name: "nonce", ParameterType: "uint256"},
					{Name: "expiry", ParameterType: "uint256"},
					{Name: "v", ParameterType: "uint8"},
					{Name: "r", ParameterType: "bytes32"},
					{Name: "s", ParameterType: "bytes32"},
				},
				Outputs: []types.Parameter{},
			},
		},
	}

	return []*types.Contract{
		{
			Address:     base.HexToAddress("0x52df6e4d9989e7cf4739d687c765e75323a1b14c"),
			Name:        "MockToken (MTK)",
			Abi:         mockERC20Abi,
			Date:        "2024-01-15",
			LastUpdated: 1705401600, // 2024-01-15 timestamp
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{
				"name":        "MockToken",
				"symbol":      "MTK",
				"decimals":    18,
				"totalSupply": "1000000000000000000000000", // 1M tokens with 18 decimals
			},
		},
		{
			Address:     base.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			Name:        "MockNFT Collection",
			Abi:         mockERC721Abi,
			Date:        "2024-02-20",
			LastUpdated: 1708473600, // 2024-02-20 timestamp
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{
				"name":   "MockNFT Collection",
				"symbol": "MNFT",
			},
		},
		{
			Address:     base.HexToAddress("0x9876543210fedcba9876543210fedcba98765432"),
			Name:        "MockDAO Governance",
			Abi:         mockGovernanceAbi,
			Date:        "2024-03-10",
			LastUpdated: 1710115200, // 2024-03-10 timestamp
			ErrorCount:  0,
			LastError:   "",
			ReadResults: map[string]interface{}{
				"proposalCount": "5",
			},
		},
	}
}

// CreateLogs returns mock log data for testing and development
func CreateLogs() []*types.Log {
	return []*types.Log{
		{
			BlockNumber:      18700000,
			TransactionIndex: 42,
			LogIndex:         0,
			TransactionHash:  base.HexToHash("0x789def456abcdef123456789abcdef123456789abcdef123456789abcdef1234"),
			Address:          base.HexToAddress("0x52df6e4d9989e7cf4739d687c765e75323a1b14c"),
			Data:             "0x000000000000000000000000000000000000000000000000000000000000007b",
			Topics: []base.Hash{
				base.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"), // Transfer event
				base.HexToHash("0x000000000000000000000000f503017d7baf7fbc0fff7492b751025c6a78179b"), // from
				base.HexToHash("0x00000000000000000000000052df6e4d9989e7cf4739d687c765e75323a1b14c"), // to
			},
		},
		{
			BlockNumber:      18700001,
			TransactionIndex: 55,
			LogIndex:         1,
			TransactionHash:  base.HexToHash("0x789abc789def123456789abcdef123456789abcdef123456789abcdef123456"),
			Address:          base.HexToAddress("0xf503017d7baf7fbc0fff7492b751025c6a78179b"),
			Data:             "0x00000000000000000000000000000000000000000000000000000000000003e8",
			Topics: []base.Hash{
				base.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"), // Approval event
				base.HexToHash("0x000000000000000000000000f503017d7baf7fbc0fff7492b751025c6a78179b"), // owner
				base.HexToHash("0x00000000000000000000000052df6e4d9989e7cf4739d687c765e75323a1b14c"), // spender
			},
		},
		{
			BlockNumber:      18700002,
			TransactionIndex: 123,
			LogIndex:         2,
			TransactionHash:  base.HexToHash("0x789ghi012345678abcdef123456789abcdef123456789abcdef123456789abc"),
			Address:          base.HexToAddress("0x52df6e4d9989e7cf4739d687c765e75323a1b14c"),
			Data:             "0x0000000000000000000000000000000000000000000000000000000000000001",
			Topics: []base.Hash{
				base.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"), // Transfer event
				base.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"), // from (mint)
				base.HexToHash("0x000000000000000000000000f503017d7baf7fbc0fff7492b751025c6a78179b"), // to
			},
		},
	}
}
