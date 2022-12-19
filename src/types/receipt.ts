import {
  address, gas, Log, uint32,
} from '.';

export type Receipt = {
  status: uint32
  contractAddress: address
  gasUsed: gas
  logs: Log[]
}
