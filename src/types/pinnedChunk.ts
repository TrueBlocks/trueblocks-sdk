/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { blknum, ipfshash } from '.';

export type PinnedChunk = {
  range: string
  bloomHash: ipfshash
  indexHash: ipfshash
  firstApp: blknum
  latestApp: blknum
}
