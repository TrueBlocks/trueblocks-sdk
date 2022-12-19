import { address, bytes, gas } from '.';

export type TraceResult = {
  newContract: address
  code: bytes
  gasUsed: gas
  output: bytes
}
