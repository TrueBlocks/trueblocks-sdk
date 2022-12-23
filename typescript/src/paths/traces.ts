/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { Trace, txId } from '../types';

export function getTraces(
  parameters?: {
    transactions: txId[],
    articulate?: boolean,
    filter?: string,
    statediff?: boolean,
    count?: boolean,
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
  return ApiCallers.fetch<Trace[]>(
    { endpoint: '/traces', method: 'get', parameters, options },
  );
}
