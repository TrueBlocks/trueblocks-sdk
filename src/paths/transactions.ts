import * as ApiCallers from '../lib/api_callers';
import { Transaction, Transfer, txId } from '../types';

export function getTransactions(
  parameters?: {
    transactions: txId[],
    articulate?: boolean,
    trace?: boolean,
    uniq?: boolean,
    flow?: 'from' | 'to',
    accountFor?: string,
    cache?: boolean,
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
  return ApiCallers.fetch<Transaction[] | Transfer[]>(
    {
      endpoint: '/transactions', method: 'get', parameters, options,
    },
  );
}
