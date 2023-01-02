/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { address, blknum, Block, topic } from '../types';

export function getBlocks(
  parameters?: {
    blocks: blknum[],
    hashes?: boolean,
    uncles?: boolean,
    trace?: boolean,
    apps?: boolean,
    uniq?: boolean,
    flow?: 'from' | 'to' | 'reward',
    logs?: boolean,
    emitter?: address[],
    topic?: topic[],
    count?: boolean,
    cache?: boolean,
    list?: blknum,
    listCount?: blknum,
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
  return ApiCallers.fetch<Block[]>(
    { endpoint: '/blocks', method: 'get', parameters, options },
  );
}
