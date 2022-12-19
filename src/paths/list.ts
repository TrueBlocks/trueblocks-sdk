import * as ApiCallers from '../lib/api_callers';
import { address, Appearance, ListStats } from '../types';

export function getList(
  parameters?: {
    addrs: address[],
    count?: boolean,
    appearances?: boolean,
    silent?: boolean,
    noZero?: boolean,
    firstRecord?: number,
    maxRecords?: number,
    firstBlock?: number,
    lastBlock?: number,
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
  return ApiCallers.fetch<Appearance[] | ListStats[]>(
    {
      endpoint: '/list', method: 'get', parameters, options,
    },
  );
}
