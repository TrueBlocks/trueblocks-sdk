/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, uint64 } from '.';

export type Monitor = {
  nApps: blknum
  firstApp: blknum
  latestApp: blknum
  sizeInBytes: uint64
  tags: string
  address: address
  name: string
  isCustom: boolean
  deleted: boolean
  symbol: string
  source: string
  decimals: uint64
  isContract: boolean
}
