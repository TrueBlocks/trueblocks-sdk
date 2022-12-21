/* eslint object-curly-newline: ["error", "never"] */
/* eslint max-len: ["error", 160] */
/*
 * This file was generated with makeClass --sdk. Do not edit it.
 */
import { uint64 } from '.';

export type Parameter = {
  type: string
  name: string
  strDefault: string
  value: string
  indexed?: boolean
  internalType: string
  components: Parameter[]
  is_flags: uint64
  precision: uint64
  maxWidth: uint64
  doc?: uint64
  disp?: uint64
  example?: string
  description?: string
  is_pointer: boolean
  is_array: boolean
  is_object: boolean
  is_builtin: boolean
  is_minimal: boolean
  is_noaddfld: boolean
  is_nowrite: boolean
  is_omitempty: boolean
  is_extra: boolean
}
