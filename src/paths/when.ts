import * as ApiCallers from '../lib/api_callers';
import { DatedBlock } from '../types';

export function getWhen(
  parameters?: {
    blocks?: string[],
    list?: boolean,
    timestamps?: boolean,
    count?: boolean,
    repair?: boolean,
    check?: boolean,
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
  return ApiCallers.fetch<DatedBlock[]>(
    {
      endpoint: '/when', method: 'get', parameters, options,
    },
  );
}
