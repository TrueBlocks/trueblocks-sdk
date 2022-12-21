/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import * as ApiCallers from '../lib/api_callers';
import { address, blknum, ChunkAddresses, ChunkAppearances, ChunkBlooms, ChunkIndex, ChunkStats, double, Manifest, PinnedChunk } from '../types';

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
    toFile?: boolean,
  },
  options?: RequestInit,
) {
  return ApiCallers.fetch<ChunkAddresses[] | ChunkAppearances[] | ChunkBlooms[] | ChunkIndex[] | ChunkStats[] | Manifest[] | PinnedChunk[]>(
    { endpoint: '/chunks', method: 'get', parameters, options },
  );
}
