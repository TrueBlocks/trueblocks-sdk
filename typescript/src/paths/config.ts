/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { Cache, CacheEntry, Chain, Config, IndexCacheItem, Key, Monitor } from '../types';

export function getConfig(
  parameters?: {
    modes?: string[],
    module?: string[],
    details?: boolean,
    terse?: boolean,
    chain: string,
    noHeader?: boolean,
    fmt?: string,
    verbose?: boolean,
    logLevel?: number,
    wei?: boolean,
    ether?: boolean,
    dollars?: boolean,
    raw?: boolean,
    toFile?: boolean,
  },
  options?: RequestInit,
) {
  return ApiCallers.fetch<Cache[] | CacheEntry[] | Chain[] | IndexCacheItem[] | Key[] | Monitor[] | Config[]>(
    { endpoint: '/config', method: 'get', parameters, options },
  );
}
