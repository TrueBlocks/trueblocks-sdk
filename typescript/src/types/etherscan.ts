/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, bytes, Function, gas, hash, timestamp, uint64, wei } from '.';

export type Etherscan = {
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
  gasCost: gas
  input: bytes
  hasToken: boolean
  articulatedTx: Function
  compressedTx: string
  isError: boolean
  functionName: string
  methodId: string
  gasUsed: gas
  contractAddress: address
  cumulativeGasUsed: string
  txReceiptStatus: string
  ether: string
}
