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
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	// EXISTING_CODE
)

type WhenOptions struct {
	BlockIds  []string          `json:"blocks,omitempty"`
	Truncate  base.Blknum       `json:"truncate,omitempty"`
	Repair    bool              `json:"repair,omitempty"`
	Update    bool              `json:"update,omitempty"`
	Deep      uint64            `json:"deep,omitempty"`
	RenderCtx *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts WhenOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// When implements the chifra when command.
func (opts *WhenOptions) When() ([]types.NamedBlock, *types.MetaData, error) {
	in := opts.toInternal()
	return queryWhen[types.NamedBlock](in)
}

// WhenList implements the chifra when --list command.
func (opts *WhenOptions) WhenList() ([]types.NamedBlock, *types.MetaData, error) {
	in := opts.toInternal()
	in.List = true
	return queryWhen[types.NamedBlock](in)
}

// WhenTimestamps implements the chifra when --timestamps command.
func (opts *WhenOptions) WhenTimestamps() ([]types.Timestamp, *types.MetaData, error) {
	in := opts.toInternal()
	in.Timestamps = true
	return queryWhen[types.Timestamp](in)
}

// WhenCount implements the chifra when --count command.
func (opts *WhenOptions) WhenCount() ([]types.Count, *types.MetaData, error) {
	in := opts.toInternal()
	in.Count = true
	return queryWhen[types.Count](in)
}

// WhenCheck implements the chifra when --check command.
func (opts *WhenOptions) WhenCheck() ([]types.ReportCheck, *types.MetaData, error) {
	in := opts.toInternal()
	in.Timestamps = true
	in.Check = true
	return queryWhen[types.ReportCheck](in)
}

// No enums
// EXISTING_CODE
func TsFromBlock(chain string, blockNum base.Blknum) (base.Timestamp, error) {
	whenOpts := WhenOptions{
		BlockIds: []string{fmt.Sprintf("%d", blockNum)},
		Globals: Globals{
			Chain: chain,
		},
	}

	var err error
	var when []types.NamedBlock
	if when, _, err = whenOpts.When(); err != nil {
		return 0, fmt.Errorf("error getting timestamp on chain %s: %w", chain, err)
	}

	return when[0].Timestamp, nil
}

// EXISTING_CODE
