/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
 * Use of this source code is governed by a license that can
 * be found in the LICENSE file.
 *
 * This file was auto generated. DO NOT EDIT.
 */

import { Abi, address, datetime, int64, timestamp } from '.';

export type Contract = {
  abi: Abi;
  address: address;
  date?: datetime;
  errorCount: int64;
  lastError: string;
  lastUpdated: timestamp;
  name: string;
};
