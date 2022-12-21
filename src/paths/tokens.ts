/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { address, blknum, TokenBalanceRecord } from '../types';

export function getTokens(
  parameters?: {
    addrs: address[],
    blocks?: blknum[],
    parts?: string[],
    byAcct?: boolean,
    noZero?: boolean,
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
  return ApiCallers.fetch<TokenBalanceRecord[]>(
    { endpoint: '/tokens', method: 'get', parameters, options },
  );
}
