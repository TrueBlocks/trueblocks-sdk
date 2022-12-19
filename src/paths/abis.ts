import * as ApiCallers from '../lib/api_callers';
import { address, Function } from '../types';

export function getAbis(
  parameters?: {
    addrs: address[],
    known?: boolean,
    sol?: boolean,
    find?: string[],
    hint?: string[],
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
  return ApiCallers.fetch<Function[]>(
    {
      endpoint: '/abis', method: 'get', parameters, options,
    },
  );
}
