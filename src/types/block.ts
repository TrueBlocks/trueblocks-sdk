import {
  address, blknum, gas, hash, timestamp, Transaction, uint64, wei,
} from '.';

export type Block = {
  gasLimit: gas
  hash: hash
  blockNumber: blknum
  parentHash: hash
  miner: address
  difficulty: uint64
  timestamp: timestamp
  transactions: Transaction[]
  baseFeePerGas: wei
  finalized: boolean
  unclesCnt: uint64
}
