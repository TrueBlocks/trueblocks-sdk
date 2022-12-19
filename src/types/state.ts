import {
  address, blknum, bytes, uint64, wei,
} from '.';

export type State = {
  blockNumber: blknum
  address: address
  proxy: address
  balance: wei
  nonce: uint64
  code: bytes
  deployed: blknum
  accttype: string
}
