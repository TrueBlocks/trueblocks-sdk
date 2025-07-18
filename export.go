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

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	// EXISTING_CODE
)

type ExportOptions struct {
	Addrs       []string          `json:"addrs,omitempty"`
	Topics      []string          `json:"topics,omitempty"`
	Fourbytes   []string          `json:"fourbytes,omitempty"`
	Articulate  bool              `json:"articulate,omitempty"`
	CacheTraces bool              `json:"cacheTraces,omitempty"`
	FirstRecord uint64            `json:"firstRecord,omitempty"`
	MaxRecords  uint64            `json:"maxRecords,omitempty"`
	Relevant    bool              `json:"relevant,omitempty"`
	Emitter     []string          `json:"emitter,omitempty"`
	Topic       []string          `json:"topic,omitempty"`
	Nfts        bool              `json:"nfts,omitempty"`
	Reverted    bool              `json:"reverted,omitempty"`
	Asset       []string          `json:"asset,omitempty"`
	Flow        ExportFlow        `json:"flow,omitempty"`
	Factory     bool              `json:"factory,omitempty"`
	Unripe      bool              `json:"unripe,omitempty"`
	Reversed    bool              `json:"reversed,omitempty"`
	NoZero      bool              `json:"noZero,omitempty"`
	FirstBlock  base.Blknum       `json:"firstBlock,omitempty"`
	LastBlock   base.Blknum       `json:"lastBlock,omitempty"`
	Accounting  bool              `json:"accounting,omitempty"`
	RenderCtx   *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts ExportOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// Export implements the chifra export command.
func (opts *ExportOptions) Export() ([]types.Transaction, *types.MetaData, error) {
	in := opts.toInternal()
	return queryExport[types.Transaction](in)
}

// ExportAppearances implements the chifra export --appearances command.
func (opts *ExportOptions) ExportAppearances() ([]types.Appearance, *types.MetaData, error) {
	in := opts.toInternal()
	in.Appearances = true
	return queryExport[types.Appearance](in)
}

// ExportReceipts implements the chifra export --receipts command.
func (opts *ExportOptions) ExportReceipts() ([]types.Receipt, *types.MetaData, error) {
	in := opts.toInternal()
	in.Receipts = true
	return queryExport[types.Receipt](in)
}

// ExportLogs implements the chifra export --logs command.
func (opts *ExportOptions) ExportLogs() ([]types.Log, *types.MetaData, error) {
	in := opts.toInternal()
	in.Logs = true
	return queryExport[types.Log](in)
}

// ExportTraces implements the chifra export --traces command.
func (opts *ExportOptions) ExportTraces() ([]types.Trace, *types.MetaData, error) {
	in := opts.toInternal()
	in.Traces = true
	return queryExport[types.Trace](in)
}

// ExportNeighbors implements the chifra export --neighbors command.
func (opts *ExportOptions) ExportNeighbors() ([]types.Message, *types.MetaData, error) {
	in := opts.toInternal()
	in.Neighbors = true
	return queryExport[types.Message](in)
}

// ExportStatements implements the chifra export --statements command.
func (opts *ExportOptions) ExportStatements() ([]types.Statement, *types.MetaData, error) {
	in := opts.toInternal()
	in.Statements = true
	return queryExport[types.Statement](in)
}

// ExportTransfers implements the chifra export --transfers command.
func (opts *ExportOptions) ExportTransfers() ([]types.Transfer, *types.MetaData, error) {
	in := opts.toInternal()
	in.Transfers = true
	return queryExport[types.Transfer](in)
}

// ExportAssets implements the chifra export --assets command.
func (opts *ExportOptions) ExportAssets() ([]types.Name, *types.MetaData, error) {
	in := opts.toInternal()
	in.Assets = true
	return queryExport[types.Name](in)
}

// ExportBalances implements the chifra export --balances command.
func (opts *ExportOptions) ExportBalances() ([]types.Token, *types.MetaData, error) {
	in := opts.toInternal()
	in.Balances = true
	return queryExport[types.Token](in)
}

// ExportWithdrawals implements the chifra export --withdrawals command.
func (opts *ExportOptions) ExportWithdrawals() ([]types.Withdrawal, *types.MetaData, error) {
	in := opts.toInternal()
	in.Withdrawals = true
	return queryExport[types.Withdrawal](in)
}

// ExportCount implements the chifra export --count command.
func (opts *ExportOptions) ExportCount() ([]types.Count, *types.MetaData, error) {
	in := opts.toInternal()
	in.Count = true
	return queryExport[types.Count](in)
}

type ExportFlow int

const (
	NoEF ExportFlow = 0
	EFIn            = 1 << iota
	EFOut
	EFZero
)

func (v ExportFlow) String() string {
	switch v {
	case NoEF:
		return "none"
	}

	var m = map[ExportFlow]string{
		EFIn:   "in",
		EFOut:  "out",
		EFZero: "zero",
	}

	var ret []string
	for _, val := range []ExportFlow{EFIn, EFOut, EFZero} {
		if v&val != 0 {
			ret = append(ret, m[val])
		}
	}

	return strings.Join(ret, ",")
}

func enumFromExportFlow(values []string) (ExportFlow, error) {
	if len(values) == 0 {
		return NoEF, fmt.Errorf("no value provided for flow option")
	}

	var result ExportFlow
	for _, val := range values {
		switch val {
		case "in":
			result |= EFIn
		case "out":
			result |= EFOut
		case "zero":
			result |= EFZero
		default:
			return NoEF, fmt.Errorf("unknown flow: %s", val)
		}
	}

	return result, nil
}

// EXISTING_CODE
func SortStatements(statuses []types.Statement, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}

func SortBalances(statuses []types.Token, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}

func SortTransfers(statuses []types.Transfer, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}

func SortTransactions(statuses []types.Transaction, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}

func SortWithdrawals(statuses []types.Withdrawal, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}

func SortAssets(statuses []types.Name, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}

func SortLogs(statuses []types.Log, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}
func SortTraces(statuses []types.Trace, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}
func SortReceipts(statuses []types.Receipt, sortSpec SortSpec) error {
	_ = statuses
	_ = sortSpec
	return nil
}

// EXISTING_CODE
