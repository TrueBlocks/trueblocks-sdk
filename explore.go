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

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/v6/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/v6/pkg/types"
	// EXISTING_CODE
)

type ExploreOptions struct {
	Terms     []string          `json:"terms,omitempty"`
	NoOpen    bool              `json:"noOpen,omitempty"`
	Local     bool              `json:"local,omitempty"`
	Google    bool              `json:"google,omitempty"`
	Dalle     bool              `json:"dalle,omitempty"`
	RenderCtx *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts ExploreOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// Explore implements the chifra explore command.
func (opts *ExploreOptions) Explore() ([]types.Destination, *types.MetaData, error) {
	in := opts.toInternal()
	return queryExplore[types.Destination](in)
}

// No enums
// EXISTING_CODE
// EXISTING_CODE
