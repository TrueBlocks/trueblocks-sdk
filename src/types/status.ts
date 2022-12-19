import {
  Cache, Chain, Key, timestamp,
} from '.';

export type Status = {
  clientVersion: string
  clientIds: string
  trueblocksVersion: string
  rpcProvider: string
  configPath: string
  cachePath: string
  indexPath: string
  host: string
  isTesting: boolean
  isApi: boolean
  isScraping: boolean
  isArchive: boolean
  isTracing: boolean
  hasEskey: boolean
  hasPinkey: boolean
  ts: timestamp
  chains: Chain[]
  caches: Cache[]
  keys: Key[]
}
