/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { blknum, datetime, ipfshash, timestamp, uint32 } from '.';

export type IndexCacheItem = {
  type: string
  nAddrs: uint32
  nApps: uint32
  firstApp: blknum
  latestApp: blknum
  firstTs: timestamp
  latestTs: timestamp
  filename: string
  fileDate: datetime
  indexSizeBytes: uint32
  indexHash: ipfshash
  bloomSizeBytes: uint32
  bloomHash: ipfshash
}
