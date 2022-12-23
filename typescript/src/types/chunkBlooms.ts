/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { blockRange, hash, uint64 } from '.';

export type ChunkBlooms = {
  range: blockRange
  magic: string
  hash: hash
  count: uint64
  nInserted: uint64
  size: uint64
  width: uint64
}
