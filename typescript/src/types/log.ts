/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, bytes, Function, hash, timestamp, topic } from '.';

export type Log = {
  address: address
  blockHash: hash
  blockNumber: blknum
  logIndex: blknum
  topics: topic[]
  data: bytes
  articulatedLog: Function
  compressedLog: string
  transactionHash: hash
  transactionIndex: blknum
  transactionLogIndex: blknum
  timestamp: timestamp
  type: string
}
