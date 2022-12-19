import * as ApiCallers from '../lib/api_callers';
import { Manifest, PinnedChunk } from '../types';

export function getScrape(
  parameters?: {
    blockCnt?: number,
    pin?: boolean,
    remote?: boolean,
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
      endpoint: '/scrape', method: 'get', parameters, options,
    },
  );
}
