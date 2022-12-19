import { address, blknum, timestamp } from '.';

export type Appearance = {
  blockNumber: blknum
  transactionIndex: blknum
  address: address
  name: string
  timestamp: timestamp
  date: string
}
