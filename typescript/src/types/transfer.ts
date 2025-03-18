/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
 * Use of this source code is governed by a license that can
 * be found in the LICENSE file.
 *
 * This file was auto generated. DO NOT EDIT.
 */

import { address, blknum, int256, lognum, txnum, uint64 } from '.';

export type Transfer = {
  amount: int256;
  asset: address;
  blockNumber: blknum;
  decimals: uint64;
  holder: address;
  logIndex: lognum;
  transactionIndex: txnum;
};
