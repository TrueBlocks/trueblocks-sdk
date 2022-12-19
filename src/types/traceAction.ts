import { address, bytes, gas } from '.';

export type TraceAction = {
  from: address
  to: address
  gas: gas
  input: bytes
  callType: string
  refundAddress: address
}
