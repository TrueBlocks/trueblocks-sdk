import { blknum, date, timestamp } from '.';

export type DatedBlock = {
  blockNumber: blknum
  timestamp: timestamp
  date: date
}
