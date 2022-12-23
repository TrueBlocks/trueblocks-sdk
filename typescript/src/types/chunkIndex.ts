/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { blockRange, hash, uint64 } from '.';

export type ChunkIndex = {
  range: blockRange
  magic: string
  hash: hash
  addressCount: uint64
  appearanceCount: uint64
  size: uint64
}
