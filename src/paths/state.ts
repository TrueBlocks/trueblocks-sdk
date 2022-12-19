import * as ApiCallers from '../lib/api_callers';
import {
  address, blknum, Result, State,
} from '../types';

export function getState(
  parameters?: {
    addrs: address[],
    blocks?: blknum[],
    parts?: string[],
    changes?: boolean,
    noZero?: boolean,
    call?: string,
    proxyFor?: string,
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
  return ApiCallers.fetch<State[] | Result[]>(
    {
      endpoint: '/state', method: 'get', parameters, options,
    },
  );
}
