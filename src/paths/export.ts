/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { address, Appearance, AppearanceCount, blknum, fourbyte, Log, Receipt, Reconciliation, topic, Trace, Transaction, Transfer, uint64 } from '../types';

export function getExport(
  parameters?: {
    addrs: address[],
    topics?: topic[],
    fourbytes?: fourbyte[],
    appearances?: boolean,
    receipts?: boolean,
    logs?: boolean,
    traces?: boolean,
    neighbors?: boolean,
    accounting?: boolean,
    statements?: boolean,
    articulate?: boolean,
    cache?: boolean,
    cacheTraces?: boolean,
    count?: boolean,
    firstRecord?: uint64,
    maxRecords?: uint64,
    relevant?: boolean,
    emitter?: address[],
    topic?: topic[],
    asset?: address[],
    flow?: 'in' | 'out' | 'zero',
    unripe?: boolean,
    firstBlock?: blknum,
    lastBlock?: blknum,
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
  return ApiCallers.fetch<AppearanceCount[] | Appearance[] | Log[] | Receipt[] | Reconciliation[] | Trace[] | Transaction[] | Transfer[]>(
    { endpoint: '/export', method: 'get', parameters, options },
  );
}
