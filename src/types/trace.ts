import {
  blknum, Function, hash, timestamp, TraceAction, TraceResult, uint64,
} from '.';

export type Trace = {
  blockHash: hash
  blockNumber: blknum
  timestamp: timestamp
  transactionHash: hash
  transactionIndex: blknum
  traceAddress: string[]
  subtraces: uint64
  type: string
  action: TraceAction
  result: TraceResult
  articulatedTrace: Function
  compressedTrace: string
}
