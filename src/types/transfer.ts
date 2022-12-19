import {
  address, blknum, Datetime, double, hash, timestamp, uint64, uint256,
} from '.';

export type Transfer = {
  blockNumber: blknum
  transactionIndex: blknum
  logIndex: blknum
  transactionHash: hash
  timestamp: timestamp
  date: Datetime
  sender: address
  recipient: address
  assetAddr: address
  assetSymbol: string
  decimals: uint64
  amount: uint256
  spotPrice: double
  priceSource: string
  encoding: string
}
