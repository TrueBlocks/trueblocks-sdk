/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
 * Use of this source code is governed by a license that can
 * be found in the LICENSE file.
 *
 * This file was auto generated. DO NOT EDIT.
 */

import * as ApiCallers from '../lib/api_callers';
import { Abi, address, Count, Function } from '../types';

export function getAbis(
  parameters?: {
    addrs: address[],
    known?: boolean,
    proxyFor?: address,
    list?: boolean,
    count?: boolean,
    find?: string[],
    hint?: string[],
    encode?: string,
    fmt?: string,
    chain: string,
    noHeader?: boolean,
    cache?: boolean,
    decache?: boolean,
  },
  options?: RequestInit,
) {
  return ApiCallers.fetch<Abi[] | Count[] | Function[]>(
    { endpoint: '/abis', method: 'get', parameters, options },
  );
}
