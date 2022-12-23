/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, gas, hash, timestamp, Transaction, uint64, wei } from '.';

export type Block = {
  gasLimit: gas
  gasUsed: gas
  hash: hash
  blockNumber: blknum
  parentHash: hash
  miner: address
  difficulty: uint64
  finalized: boolean
  timestamp: timestamp
  baseFeePerGas: wei
  transactions: Transaction[]
  tx_hashes: string[]
  name: string
  unclesCnt: uint64
  transactionsCnt: uint64
}
