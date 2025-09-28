package sdk

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"

// Modeler interface exposes the Model method for SDK consumers
type Modeler interface {
	Model(chain, format string, verbose bool, extraOpts map[string]any) types.Model
}

// Model type alias for the return type of Model()
type Model = types.Model

type Abi = types.Abi
type Appearance = types.Appearance
type Block = types.Block
type BlockCount = types.BlockCount
type Bounds = types.Bounds
type CacheItem = types.CacheItem
type Chain = types.Chain
type ChunkAddress = types.ChunkAddress
type ChunkAppearance = types.ChunkAppearance
type ChunkBloom = types.ChunkBloom
type ChunkIndex = types.ChunkIndex
type ChunkManifest = types.ChunkManifest
type ChunkPin = types.ChunkPin
type ChunkRecord = types.ChunkRecord
type ChunkStats = types.ChunkStats
type Config = types.Config
type Count = types.Count
type Destination = types.Destination
type Function = types.Function
type LightBlock = types.LightBlock
type Log = types.Log
type Message = types.Message
type Monitor = types.Monitor
type MonitorClean = types.MonitorClean
type Name = types.Name
type NamedBlock = types.NamedBlock
type Receipt = types.Receipt
type ReportCheck = types.ReportCheck
type Slurp = types.Slurp
type State = types.State
type Statement = types.Statement
type Status = types.Status
type Timestamp = types.Timestamp
type Token = types.Token
type Trace = types.Trace
type TraceCount = types.TraceCount
type Transaction = types.Transaction
type Transfer = types.Transfer
type Withdrawal = types.Withdrawal
