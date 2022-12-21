/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, uint64 } from '.';

export type CacheEntry = {
  type: uint64
  extra: string
  cached: boolean
  path: string
  address: address
  name: string
}
