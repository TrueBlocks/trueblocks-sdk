// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package sdk

import (
	// EXISTING_CODE

	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	// EXISTING_CODE
)

type AbisOptions struct {
	Addrs     []string          `json:"addrs,omitempty"`
	Known     bool              `json:"known,omitempty"`
	ProxyFor  base.Address      `json:"proxyFor,omitempty"`
	Hint      []string          `json:"hint,omitempty"`
	RenderCtx *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts AbisOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// Abis implements the chifra abis command.
func (opts *AbisOptions) Abis() ([]types.Function, *types.MetaData, error) {
	in := opts.toInternal()
	return queryAbis[types.Function](in)
}

// AbisList implements the chifra abis --list command.
func (opts *AbisOptions) AbisList() ([]types.Abi, *types.MetaData, error) {
	in := opts.toInternal()
	in.List = true
	return queryAbis[types.Abi](in)
}

// AbisDetails implements the chifra abis --details command.
func (opts *AbisOptions) AbisDetails() ([]types.Function, *types.MetaData, error) {
	in := opts.toInternal()
	in.Details = true
	return queryAbis[types.Function](in)
}

// AbisCount implements the chifra abis --count command.
func (opts *AbisOptions) AbisCount() ([]types.Count, *types.MetaData, error) {
	in := opts.toInternal()
	in.Count = true
	return queryAbis[types.Count](in)
}

// AbisFind implements the chifra abis --find command.
func (opts *AbisOptions) AbisFind(val []string) ([]types.Function, *types.MetaData, error) {
	in := opts.toInternal()
	in.Find = val
	return queryAbis[types.Function](in)
}

// AbisEncode implements the chifra abis --encode command.
func (opts *AbisOptions) AbisEncode(val string) ([]types.Function, *types.MetaData, error) {
	in := opts.toInternal()
	in.Encode = val
	return queryAbis[types.Function](in)
}

// No enums
// EXISTING_CODE
// EXISTING_CODE
