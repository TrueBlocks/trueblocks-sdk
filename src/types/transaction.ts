import {
  address, blknum, bytes, date, Function, gas, hash, Receipt, Reconciliation, timestamp, uint64, wei,
} from '.';

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
  date: date
}
