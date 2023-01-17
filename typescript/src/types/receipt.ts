/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, gas, hash, Log, uint32, wei } from '.';

export type Receipt = {
  blockHash?: hash
  blockNumber: blknum
  contractAddress?: address
  cumulativeGasUsed?: wei
  gasUsed: gas
  effectiveGasPrice?: gas
  from?: address
  logs: Log[]
  status: uint32
  to?: address
  transactionHash: hash
  transactionIndex: blknum
  isError?: boolean
}
