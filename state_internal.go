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
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	state "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/sdk"
	// EXISTING_CODE
)

type stateOptionsInternal struct {
	Addrs      []string          `json:"addrs,omitempty"`
	BlockIds   []string          `json:"blocks,omitempty"`
	Parts      StateParts        `json:"parts,omitempty"`
	Changes    bool              `json:"changes,omitempty"`
	NoZero     bool              `json:"noZero,omitempty"`
	Call       bool              `json:"call,omitempty"`
	Send       bool              `json:"send,omitempty"`
	Calldata   string            `json:"calldata,omitempty"`
	Articulate bool              `json:"articulate,omitempty"`
	ProxyFor   base.Address      `json:"proxyFor,omitempty"`
	RenderCtx  *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts *stateOptionsInternal) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// StateBytes implements the chifra state command for the SDK.
func (opts *stateOptionsInternal) StateBytes(w io.Writer) error {
	values, err := structToValues(*opts)
	if err != nil {
		return fmt.Errorf("error converting state struct to URL values: %v", err)
	}

	if opts.RenderCtx == nil {
		opts.RenderCtx = output.NewRenderContext()
	}
	return state.State(opts.RenderCtx, w, values)
}

// stateParseFunc handles special cases such as structs and enums (if any).
func stateParseFunc(target any, key, value string) (bool, error) {
	_ = key
	_ = value
	var found bool
	opts, ok := target.(*stateOptionsInternal)
	if !ok {
		return false, fmt.Errorf("parseFunc(state): target is not of correct type")
	}

	if key == "parts" {
		var err error
		values := strings.Split(value, ",")
		if opts.Parts, err = enumFromStateParts(values); err != nil {
			return false, err
		} else {
			found = true
		}
	}

	// EXISTING_CODE
	if key == "proxyFor" {
		opts.ProxyFor = base.HexToAddress(value)
		return base.IsValidAddress(value), nil
	}
	// EXISTING_CODE

	return found, nil
}

// GetStateOptions returns a filled-in options instance given a string array of arguments.
func GetStateOptions(args []string) (*stateOptionsInternal, error) {
	var opts stateOptionsInternal
	if err := assignValuesFromArgs(args, stateParseFunc, &opts, &opts.Globals); err != nil {
		return nil, err
	}

	return &opts, nil
}

type stateGeneric interface {
	types.State |
		types.Result
}

func queryState[T stateGeneric](opts *stateOptionsInternal) ([]T, *types.MetaData, error) {
	// EXISTING_CODE
	// EXISTING_CODE

	buffer := bytes.Buffer{}
	if err := opts.StateBytes(&buffer); err != nil {
		return nil, nil, err
	}

	str := buffer.String()
	// EXISTING_CODE
	str = convertObjectToArray("inputs", str)
	str = convertObjectToArray("outputs", str)
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
func (opts *StateOptions) toInternal() *stateOptionsInternal {
	return &stateOptionsInternal{
		Addrs:      opts.Addrs,
		BlockIds:   opts.BlockIds,
		Parts:      opts.Parts,
		Changes:    opts.Changes,
		NoZero:     opts.NoZero,
		Calldata:   opts.Calldata,
		Articulate: opts.Articulate,
		ProxyFor:   opts.ProxyFor,
		RenderCtx:  opts.RenderCtx,
		Globals:    opts.Globals,
	}
}

// EXISTING_CODE
// EXISTING_CODE
