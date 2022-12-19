import * as ApiCallers from '../lib/api_callers';
import {
  address, Appearance, fourbyte, ListStats, Log, Receipt, Reconciliation, topic, Trace, Transaction, Transfer,
} from '../types';

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
    firstRecord?: number,
    maxRecords?: number,
    relevant?: boolean,
    emitter?: address[],
    topic?: topic[],
    asset?: address[],
    flow?: 'in' | 'out' | 'zero',
    unripe?: boolean,
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
  return ApiCallers.fetch<Appearance[] | Reconciliation[] | ListStats[] | Transaction[] | Receipt[] | Log[] | Trace[] | Transfer[]>(
    {
      endpoint: '/export', method: 'get', parameters, options,
    },
  );
}
