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
	"errors"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/crud"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	// EXISTING_CODE
)

type NamesOptions struct {
	Terms     []string          `json:"terms,omitempty"`
	Expand    bool              `json:"expand,omitempty"`
	MatchCase bool              `json:"matchCase,omitempty"`
	All       bool              `json:"all,omitempty"`
	Custom    bool              `json:"custom,omitempty"`
	Prefund   bool              `json:"prefund,omitempty"`
	Regular   bool              `json:"regular,omitempty"`
	DryRun    bool              `json:"dryRun,omitempty"`
	RenderCtx *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts NamesOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// Names implements the chifra names command.
func (opts *NamesOptions) Names() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	return queryNames[types.Name](in)
}

// NamesAddr implements the chifra names --addr command.
func (opts *NamesOptions) NamesAddr() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	in.Addr = true
	return queryNames[types.Name](in)
}

// NamesTags implements the chifra names --tags command.
func (opts *NamesOptions) NamesTags() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	in.Tags = true
	return queryNames[types.Name](in)
}

// NamesClean implements the chifra names --clean command.
func (opts *NamesOptions) NamesClean() ([]types.Message, *types.MetaData, error) {
	in := opts.toInternal()
	in.Clean = true
	return queryNames[types.Message](in)
}

// NamesCount implements the chifra names --count command.
func (opts *NamesOptions) NamesCount() ([]types.Count, *types.MetaData, error) {
	in := opts.toInternal()
	in.Count = true
	return queryNames[types.Count](in)
}

// NamesAutoname implements the chifra names --autoname command.
func (opts *NamesOptions) NamesAutoname(val base.Address) ([]types.Message, *types.MetaData, error) {
	in := opts.toInternal()
	in.Autoname = val
	return queryNames[types.Message](in)
}

// NamesCreate implements the chifra names --create command.
func (opts *NamesOptions) NamesCreate() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	in.Create = true
	return queryNames[types.Name](in)
}

// NamesUpdate implements the chifra names --update command.
func (opts *NamesOptions) NamesUpdate() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	in.Update = true
	return queryNames[types.Name](in)
}

// NamesDelete implements the chifra names --delete command.
func (opts *NamesOptions) NamesDelete() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	in.Delete = true
	return queryNames[types.Name](in)
}

// NamesUndelete implements the chifra names --undelete command.
func (opts *NamesOptions) NamesUndelete() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	in.Undelete = true
	return queryNames[types.Name](in)
}

// NamesRemove implements the chifra names --remove command.
func (opts *NamesOptions) NamesRemove() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	in.Remove = true
	return queryNames[types.Name](in)
}

// No enums
// EXISTING_CODE
func (opts *NamesOptions) ModifyName(op crud.Operation, cd *crud.NameCrud) ([]types.Name, *types.MetaData, error) {
	defer func() {
		cd.Unsetenv()
	}()

	opts.Terms = []string{cd.Address.Value}
	cd.SetEnv()

	switch op {
	case crud.Create:
		return opts.NamesCreate()
	case crud.Update:
		return opts.NamesUpdate()
	case crud.Delete:
		return opts.NamesDelete()
	case crud.Undelete:
		return opts.NamesUndelete()
	case crud.Remove:
		return opts.NamesRemove()
	case crud.Autoname:
		addr := base.HexToAddress(cd.Address.Value)
		// TODO: NamesAutoname should return the names array like everything else
		_, meta, err := opts.NamesAutoname(addr)
		return nil, meta, err
	}

	return nil, nil, errors.New("invalid operation " + string(op))
}

// EXISTING_CODE
