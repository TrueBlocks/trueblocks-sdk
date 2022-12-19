import { address, blknum, Function } from '.';

export type Result = {
  blockNumber: blknum
  address: address
  signature: string
  encoding: string
  bytes: string
  callResult: Function
  compressedResult: string
  deployed: blknum
}
