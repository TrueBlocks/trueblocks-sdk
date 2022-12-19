import { Parameter } from '.';

export type Function = {
  name: string
  type: string
  signature: string
  encoding: string
  inputs: Parameter[]
  outputs: Parameter[]
}
