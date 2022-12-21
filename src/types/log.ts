/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, bytes, Function, hash, timestamp, topic } from '.';

export type Log = {
  blockNumber: blknum
  transactionIndex: blknum
  logIndex: blknum
  transactionHash: hash
  timestamp: timestamp
  address: address
  topics: topic[]
  data: bytes
  articulatedLog: Function
  compressedLog: string
}
