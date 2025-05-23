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

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	// EXISTING_CODE
)

type MonitorsOptions struct {
	Addrs     []string          `json:"addrs,omitempty"`
	Delete    bool              `json:"delete,omitempty"`
	Undelete  bool              `json:"undelete,omitempty"`
	Remove    bool              `json:"remove,omitempty"`
	Staged    bool              `json:"staged,omitempty"`
	RenderCtx *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts MonitorsOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// Monitors implements the chifra monitors command.
func (opts *MonitorsOptions) Monitors() ([]types.Message, *types.MetaData, error) {
	in := opts.toInternal()
	return queryMonitors[types.Message](in)
}

// MonitorsClean implements the chifra monitors --clean command.
func (opts *MonitorsOptions) MonitorsClean() ([]types.MonitorClean, *types.MetaData, error) {
	in := opts.toInternal()
	in.Clean = true
	return queryMonitors[types.MonitorClean](in)
}

// MonitorsList implements the chifra monitors --list command.
func (opts *MonitorsOptions) MonitorsList() ([]types.Monitor, *types.MetaData, error) {
	in := opts.toInternal()
	in.List = true
	return queryMonitors[types.Monitor](in)
}

// MonitorsCount implements the chifra monitors --count command.
func (opts *MonitorsOptions) MonitorsCount() ([]types.Count, *types.MetaData, error) {
	in := opts.toInternal()
	in.Count = true
	return queryMonitors[types.Count](in)
}

// No enums
// EXISTING_CODE
// EXISTING_CODE
