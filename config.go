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
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	// EXISTING_CODE
)

type ConfigOptions struct {
	Mode      ConfigMode        `json:"mode,omitempty"`
	RenderCtx *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts ConfigOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// ConfigPaths implements the chifra config --paths command.
func (opts *ConfigOptions) ConfigPaths() ([]types.CacheItem, *types.MetaData, error) {
	in := opts.toInternal()
	in.Paths = true
	return queryConfig[types.CacheItem](in)
}

// ConfigDump implements the chifra config --dump command.
func (opts *ConfigOptions) ConfigDump() ([]types.Config, *types.MetaData, error) {
	in := opts.toInternal()
	in.Dump = true
	return queryConfig[types.Config](in)
}

type ConfigMode int

const (
	NoCOM  ConfigMode = 0
	CMShow            = 1 << iota
	CMEdit
)

func (v ConfigMode) String() string {
	switch v {
	case NoCOM:
		return "none"
	}

	var m = map[ConfigMode]string{
		CMShow: "show",
		CMEdit: "edit",
	}

	var ret []string
	for _, val := range []ConfigMode{CMShow, CMEdit} {
		if v&val != 0 {
			ret = append(ret, m[val])
		}
	}

	return strings.Join(ret, ",")
}

func enumFromConfigMode(values []string) (ConfigMode, error) {
	if len(values) == 0 {
		return NoCOM, fmt.Errorf("no value provided for mode option")
	}

	var result ConfigMode
	for _, val := range values {
		switch val {
		case "show":
			result |= CMShow
		case "edit":
			result |= CMEdit
		default:
			return NoCOM, fmt.Errorf("unknown mode: %s", val)
		}
	}

	return result, nil
}

// EXISTING_CODE
// EXISTING_CODE
