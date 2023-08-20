/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { address, Appearance, AppearanceCount, blknum, Bounds, Monitor, uint64 } from '../types';

export function getList(
  parameters?: {
    addrs: address[],
    count?: boolean,
    noZero?: boolean,
    bounds?: boolean,
    unripe?: boolean,
    silent?: boolean,
    firstRecord?: uint64,
    maxRecords?: uint64,
    reversed?: boolean,
    firstBlock?: blknum,
    lastBlock?: blknum,
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
  return ApiCallers.fetch<Appearance[] | AppearanceCount[] | Bounds[] | Monitor[]>(
    { endpoint: '/list', method: 'get', parameters, options },
  );
}
