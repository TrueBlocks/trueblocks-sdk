/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
 * Use of this source code is governed by a license that can
 * be found in the LICENSE file.
 *
 * This file was auto generated. DO NOT EDIT.
 */

import { address, blknum, int256, Log, lognum, Transaction, txnum, uint64 } from '.';

export type Transfer = {
  amountIn?: int256;
  amountOut?: int256;
  asset: address;
  blockNumber: blknum;
  decimals: uint64;
  gasOut?: int256;
  holder: address;
  internalIn?: int256;
  internalOut?: int256;
  log?: Log;
  logIndex: lognum;
  minerBaseRewardIn?: int256;
  minerNephewRewardIn?: int256;
  minerTxFeeIn?: int256;
  minerUncleRewardIn?: int256;
  prefundIn?: int256;
  recipient: address;
  selfDestructIn?: int256;
  selfDestructOut?: int256;
  sender: address;
  transaction?: Transaction;
  transactionIndex: txnum;
};
