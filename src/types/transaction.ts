/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, bytes, datetime, Function, gas, hash, Receipt, Reconciliation, timestamp, uint64, wei } from '.';

export type Transaction = {
  hash: hash
  blockHash: hash
  blockNumber: blknum
  transactionIndex: blknum
  nonce: uint64
  timestamp: timestamp
  from: address
  to: address
  value: wei
  gas: gas
  gasPrice: gas
  input: bytes
  receipt: Receipt
  statements: Reconciliation[]
  articulatedTx: Function
  compressedTx: string
  hasToken: boolean
  finalized: boolean
  extraData: string
  isError: boolean
  date: datetime
}
