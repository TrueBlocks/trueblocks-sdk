/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
 * Use of this source code is governed by a license that can
 * be found in the LICENSE file.
 *
 * This file was auto generated. DO NOT EDIT.
 */

import { blkrange, hash, RangeDates, uint64 } from '.';

export type ChunkBloom = {
  byteWidth: uint64;
  fileSize: uint64;
  hash: hash;
  magic: string;
  nBlooms: uint64;
  nInserted: uint64;
  range: blkrange;
  rangeDates?: RangeDates;
};
