import { uint64 } from '.';

export type Chain = {
  chain: string
  chainId: uint64
  symbol: string
  rpcProvider: string
  apiProvider: string
  remoteExplorer: string
  localExplorer: string
  ipfsGateway: string
}
