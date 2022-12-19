import * as ApiCallers from '../lib/api_callers';
import { address, blknum, Transaction } from '../types';

export function getSlurp(
  parameters?: {
    addrs: address[],
    blocks?: blknum[],
    types?: string[],
    appearances?: boolean,
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
  return ApiCallers.fetch<Transaction[]>(
    {
      endpoint: '/slurp', method: 'get', parameters, options,
    },
  );
}
