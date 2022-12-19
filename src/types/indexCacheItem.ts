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
