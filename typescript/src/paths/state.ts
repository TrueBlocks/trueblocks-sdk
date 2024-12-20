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
import { address, blknum, Result, State } from '../types';

export function getState(
  parameters?: {
    addrs: address[];
    blocks?: blknum[];
    parts?: string[];
    changes?: boolean;
    noZero?: boolean;
    call?: boolean;
    calldata?: string;
    articulate?: boolean;
    proxyFor?: address;
    fmt?: string;
    chain: string;
    noHeader?: boolean;
    cache?: boolean;
    decache?: boolean;
    ether?: boolean;
  },
  options?: RequestInit,
) {
  return ApiCallers.fetch<Result[] | State[]>(
    { endpoint: '/state', method: 'get', parameters, options },
  );
}
