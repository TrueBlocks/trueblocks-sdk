/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { address, Appearance, blknum, ChunkAddress, ChunkBlooms, ChunkIndex, ChunkStats, double, Manifest, PinnedChunk, ReportCheck } from '../types';

export function getChunks(
  parameters?: {
    mode: 'status' | 'manifest' | 'index' | 'blooms' | 'addresses' | 'appearances' | 'stats',
    blocks?: blknum[],
    check?: boolean,
    pin?: boolean,
    publish?: boolean,
    remote?: boolean,
    belongs?: address[],
    sleep?: double,
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
  return ApiCallers.fetch<Appearance[] | ChunkAddress[] | ChunkBlooms[] | ChunkIndex[] | ChunkStats[] | Manifest[] | PinnedChunk[] | ReportCheck[]>(
    { endpoint: '/chunks', method: 'get', parameters, options },
  );
}
