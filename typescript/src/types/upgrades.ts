/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { ipfshash, int64 } from '.';

export type ChunkRecordUp = {
  range: string
  bloomHash: ipfshash
  indexHash: ipfshash
  bloomSize: int64
  indexSize: int64
  firstApp:  int64
}
