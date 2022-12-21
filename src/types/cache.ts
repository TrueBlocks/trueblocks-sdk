/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { CacheEntry, uint64 } from '.';

export type Cache = {
  type: string
  path: string
  nFiles: uint64
  nFolders: uint64
  sizeInBytes: uint64
  items: CacheEntry[]
}
