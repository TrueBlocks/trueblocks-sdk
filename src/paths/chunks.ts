import * as ApiCallers from '../lib/api_callers';
import {
  address, blknum, Manifest, PinnedChunk,
} from '../types';

export function getChunks(
  parameters?: {
    mode: 'status' | 'manifest' | 'index' | 'blooms' | 'addresses' | 'appearances' | 'stats',
    blocks?: blknum[],
    check?: boolean,
    pin?: boolean,
    publish?: boolean,
    remote?: boolean,
    belongs?: address[],
    sleep?: number,
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
  return ApiCallers.fetch<PinnedChunk[] | Manifest[]>(
    {
      endpoint: '/chunks', method: 'get', parameters, options,
    },
  );
}
