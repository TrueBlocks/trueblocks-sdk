/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { address, blknum, CallResult, EthCall } from '../types';

export function getState(
  parameters?: {
    addrs: address[],
    blocks?: blknum[],
    parts?: string[],
    changes?: boolean,
    noZero?: boolean,
    call?: string,
    proxyFor?: address,
    chain: string,
    noHeader?: boolean,
    fmt?: string,
    verbose?: boolean,
    wei?: boolean,
    ether?: boolean,
    dollars?: boolean,
    raw?: boolean,
  },
  options?: RequestInit,
) {
  return ApiCallers.fetch<CallResult[] | EthCall[]>(
    { endpoint: '/state', method: 'get', parameters, options },
  );
}
