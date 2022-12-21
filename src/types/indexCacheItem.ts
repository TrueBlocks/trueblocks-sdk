/* eslint object-curly-newline: ["error", "never"] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { ipfshash } from '.';

export type IndexCacheItem = {
  nAddrs: number,
  nApps: number,
  firstApp: number,
  latestApp: number,
  firstTs: number,
  latestTs: number,
  filename: string,
  fileDate: string,
  indexSizeBytes: number,
  indexHash: ipfshash,
  bloomSizeBytes: number,
  bloomHash: ipfshash,
}
