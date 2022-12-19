import { address, uint64 } from '.';

export type Name = {
  tags: string
  address: address
  name: string
  symbol: string
  source: string
  decimals: uint64
  petname: string
  deleted: boolean
  isCustom: boolean
  isPrefund: boolean
  isContract: boolean
  isErc20: boolean
  isErc721: boolean
}
