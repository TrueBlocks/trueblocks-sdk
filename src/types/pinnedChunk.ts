import { blknum, ipfshash } from '.';

export type PinnedChunk = {
  range: string
  bloomHash: ipfshash
  indexHash: ipfshash
  firstApp: blknum
  latestApp: blknum
}
