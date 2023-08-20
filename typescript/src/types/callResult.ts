/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, blknum, Outputs } from '.';

export type CallResult = {
  blockNumber: blknum
  address: address
  name: string
  encoding: string
  signature: string
  encodedArguments: string
  outputs: Outputs
}
