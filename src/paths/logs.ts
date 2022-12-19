import * as ApiCallers from '../lib/api_callers';
import { Log, txId } from '../types';

export function getLogs(
  parameters?: {
    transactions: txId[],
    articulate?: boolean,
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
  return ApiCallers.fetch<Log[]>(
    {
      endpoint: '/logs', method: 'get', parameters, options,
    },
  );
}
