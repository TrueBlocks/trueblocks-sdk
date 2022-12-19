import { ipfshash, PinnedChunk } from '.';

export type Manifest = {
  version: string
  chain: string
  schemas: ipfshash
  databases: ipfshash
  chunks: PinnedChunk[]
}
