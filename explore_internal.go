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
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	explore "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/sdk"
	// EXISTING_CODE
)

type exploreOptionsInternal struct {
	Terms     []string          `json:"terms,omitempty"`
	NoOpen    bool              `json:"noOpen,omitempty"`
	Local     bool              `json:"local,omitempty"`
	Google    bool              `json:"google,omitempty"`
	Dalle     bool              `json:"dalle,omitempty"`
	RenderCtx *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts *exploreOptionsInternal) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// ExploreBytes implements the chifra explore command for the SDK.
func (opts *exploreOptionsInternal) ExploreBytes(w io.Writer) error {
	values, err := structToValues(*opts)
	if err != nil {
		return fmt.Errorf("error converting explore struct to URL values: %v", err)
	}

	if opts.RenderCtx == nil {
		opts.RenderCtx = output.NewRenderContext()
	}
	return explore.Explore(opts.RenderCtx, w, values)
}

// exploreParseFunc handles special cases such as structs and enums (if any).
func exploreParseFunc(target any, key, value string) (bool, error) {
	var found bool
	_, ok := target.(*exploreOptionsInternal)
	if !ok {
		return false, fmt.Errorf("parseFunc(explore): target is not of correct type")
	}

	// No enums
	// EXISTING_CODE
	// EXISTING_CODE

	return found, nil
}

// GetExploreOptions returns a filled-in options instance given a string array of arguments.
func GetExploreOptions(args []string) (*exploreOptionsInternal, error) {
	var opts exploreOptionsInternal
	if err := assignValuesFromArgs(args, exploreParseFunc, &opts, &opts.Globals); err != nil {
		return nil, err
	}

	return &opts, nil
}

type exploreGeneric interface {
	types.Destination
}

func queryExplore[T exploreGeneric](opts *exploreOptionsInternal) ([]T, *types.MetaData, error) {
	// EXISTING_CODE
	// EXISTING_CODE

	buffer := bytes.Buffer{}
	if err := opts.ExploreBytes(&buffer); err != nil {
		return nil, nil, err
	}

	str := buffer.String()
	// EXISTING_CODE
	// EXISTING_CODE

	var result Result[T]
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		debugPrint(str, result, err)
		return nil, nil, err
	} else {
		return result.Data, &result.Meta, nil
	}
}

// toInternal converts the SDK options to the internal options format.
func (opts *ExploreOptions) toInternal() *exploreOptionsInternal {
	return &exploreOptionsInternal{
		Terms:     opts.Terms,
		NoOpen:    opts.NoOpen,
		Local:     opts.Local,
		Google:    opts.Google,
		Dalle:     opts.Dalle,
		RenderCtx: opts.RenderCtx,
		Globals:   opts.Globals,
	}
}

// EXISTING_CODE
// EXISTING_CODE
