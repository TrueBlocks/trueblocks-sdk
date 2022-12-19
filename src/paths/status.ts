import * as ApiCallers from '../lib/api_callers';
import { Status } from '../types';

export function getConfig(
  parameters?: {
    modes?: string[],
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
  return ApiCallers.fetch<Status[]>(
    {
      endpoint: '/status', method: 'get', parameters, options,
    },
  );
}
