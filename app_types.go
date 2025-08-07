package sdk

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"

// Modeler interface exposes the Model method for SDK consumers
type Modeler interface {
	Model(chain, format string, verbose bool, extraOpts map[string]any) types.Model
}

// Model type alias for the return type of Model()
type Model = types.Model

type Abi = types.Abi
type Function = types.Function
type Parameter = types.Parameter
type Bloom = types.ChunkBloom
type Index = types.ChunkIndex
type Manifest = types.ChunkManifest
type Stats = types.ChunkStats
type Balance = types.Token
type Statement = types.Statement
type Transaction = types.Transaction
type Transfer = types.Transfer
type Withdrawal = types.Withdrawal
type Monitor = types.Monitor
type Name = types.Name
type Asset = types.Name
type Log = types.Log
type Receipt = types.Receipt
type Trace = types.Trace
type Cache = types.CacheItem
type Status = types.Status
type Chain = types.Chain
type Contract = types.Contract
