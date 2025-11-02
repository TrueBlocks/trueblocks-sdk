package sdk

import "github.com/TrueBlocks/trueblocks-chifra/v6/pkg/types"

type CoreResult[T any] struct {
	Data []T            `json:"data"`
	Meta types.MetaData `json:"meta"`
}
