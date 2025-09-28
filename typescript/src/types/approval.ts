/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
 * Use of this source code is governed by a license that can
 * be found in the LICENSE file.
 *
 * This file was auto generated. DO NOT EDIT.
 */

import { address, blknum, datetime, lognum, timestamp, txnum, wei } from '.';

export type Approval = {
  allowance: wei;
  blockNumber: blknum;
  date?: datetime;
  lastAppBlock: blknum;
  lastAppLogID: lognum;
  lastAppTs: timestamp;
  lastAppTxID: txnum;
  owner: address;
  spender: address;
  timestamp: timestamp;
  token: address;
};
