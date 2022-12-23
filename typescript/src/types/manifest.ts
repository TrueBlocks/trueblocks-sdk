/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { ipfshash, PinnedChunk } from '.';

export type Manifest = {
  version: string
  chain: string
  schemas: ipfshash
  databases: ipfshash
  chunks: PinnedChunk[]
}
