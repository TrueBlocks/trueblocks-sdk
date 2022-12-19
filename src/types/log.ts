import {
  address, blknum, bytes, Function, hash, timestamp, topic,
} from '.';

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
