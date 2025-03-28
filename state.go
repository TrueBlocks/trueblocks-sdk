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

type StateOptions struct {
	Addrs      []string          `json:"addrs,omitempty"`
	BlockIds   []string          `json:"blocks,omitempty"`
	Parts      StateParts        `json:"parts,omitempty"`
	Changes    bool              `json:"changes,omitempty"`
	NoZero     bool              `json:"noZero,omitempty"`
	Calldata   string            `json:"calldata,omitempty"`
	Articulate bool              `json:"articulate,omitempty"`
	ProxyFor   base.Address      `json:"proxyFor,omitempty"`
	RenderCtx  *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts StateOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

// State implements the chifra state command.
func (opts *StateOptions) State() ([]types.State, *types.MetaData, error) {
	in := opts.toInternal()
	return queryState[types.State](in)
}

// StateCall implements the chifra state --call command.
func (opts *StateOptions) StateCall() ([]types.Result, *types.MetaData, error) {
	in := opts.toInternal()
	in.Call = true
	return queryState[types.Result](in)
}

// StateSend implements the chifra state --send command.
func (opts *StateOptions) StateSend() ([]types.Result, *types.MetaData, error) {
	in := opts.toInternal()
	in.Send = true
	return queryState[types.Result](in)
}

type StateParts int

const (
	NoSTP     StateParts = 0
	SPBalance            = 1 << iota
	SPNonce
	SPCode
	SPProxy
	SPDeployed
	SPAccttype
	STPSome = SPBalance | SPProxy | SPDeployed | SPAccttype
	STPAll  = SPBalance | SPNonce | SPCode | SPProxy | SPDeployed | SPAccttype
)

func (v StateParts) String() string {
	switch v {
	case NoSTP:
		return "none"
	case STPSome:
		return "some"
	case STPAll:
		return "all"
	}

	var m = map[StateParts]string{
		SPBalance:  "balance",
		SPNonce:    "nonce",
		SPCode:     "code",
		SPProxy:    "proxy",
		SPDeployed: "deployed",
		SPAccttype: "accttype",
	}

	var ret []string
	for _, val := range []StateParts{SPBalance, SPNonce, SPCode, SPProxy, SPDeployed, SPAccttype} {
		if v&val != 0 {
			ret = append(ret, m[val])
		}
	}

	return strings.Join(ret, ",")
}

func enumFromStateParts(values []string) (StateParts, error) {
	if len(values) == 0 {
		return NoSTP, fmt.Errorf("no value provided for parts option")
	}

	if len(values) == 1 && values[0] == "all" {
		return STPAll, nil
	} else if len(values) == 1 && values[0] == "some" {
		return STPSome, nil
	}

	var result StateParts
	for _, val := range values {
		switch val {
		case "balance":
			result |= SPBalance
		case "nonce":
			result |= SPNonce
		case "code":
			result |= SPCode
		case "proxy":
			result |= SPProxy
		case "deployed":
			result |= SPDeployed
		case "accttype":
			result |= SPAccttype
		default:
			return NoSTP, fmt.Errorf("unknown parts: %s", val)
		}
	}

	return result, nil
}

// EXISTING_CODE
// EXISTING_CODE
