/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, bytes32, gas, hash, Log, uint32, wei } from '.';

export type Receipt = {
  blockHash: hash
  blockNumber: blknum
  contractAddress: address
  cumulativeGasUsed: wei
  from: address
  gasUsed: gas
  effectiveGasPrice: gas
  logs: Log[]
  root: bytes32
  status: uint32
  to: address
  transactionHash: hash
  transactionIndex: blknum
  hash: string
  isError: boolean
}
