package sdk

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/v5/pkg/types"

type CoreResult[T any] struct {
	Data []T            `json:"data"`
	Meta types.MetaData `json:"meta"`
}
