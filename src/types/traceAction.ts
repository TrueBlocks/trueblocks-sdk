/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { address, bytes, gas } from '.';

export type TraceAction = {
  from: address
  to: address
  gas: gas
  input: bytes
  callType: string
  refundAddress: address
}
