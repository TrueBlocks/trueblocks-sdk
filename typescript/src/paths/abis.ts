/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { Abi, address, Function, Parameter } from '../types';

export function getAbis(
  parameters?: {
    addrs: address[],
    known?: boolean,
    sol?: boolean,
    find?: string[],
    hint?: string[],
    chain: string,
    noHeader?: boolean,
    fmt?: string,
    verbose?: boolean,
    logLevel?: number,
    wei?: boolean,
    ether?: boolean,
    dollars?: boolean,
    raw?: boolean,
  },
  options?: RequestInit,
) {
  return ApiCallers.fetch<Abi[] | Function[] | Parameter[]>(
    { endpoint: '/abis', method: 'get', parameters, options },
  );
}
