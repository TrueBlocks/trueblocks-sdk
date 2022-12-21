/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, gas, Log, uint32 } from '.';

export type Receipt = {
  status: uint32
  contractAddress: address
  gasUsed: gas
  logs: Log[]
}
