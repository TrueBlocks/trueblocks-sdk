package sdk

// EXISTING_CODE
import (
	"fmt"
	"slices"
	"sort"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type SortOrder = types.SortOrder

const (
	Asc SortOrder = types.Ascending
	Dec SortOrder = types.Descending
)

type SortSpec struct {
	Fields []string    `json:"fields"`
	Order  []SortOrder `json:"orders"`
}

func SortMonitors(monitors []types.Monitor, sortSpec SortSpec) error {
	if len(sortSpec.Fields) != len(sortSpec.Order) {
		return fmt.Errorf("fields and order must have the same length")
	}

	sorts := make([]func(p1, p2 types.Monitor) bool, len(sortSpec.Fields))
	for i, field := range sortSpec.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(types.GetSortFieldsMonitor(), field) {
			return fmt.Errorf("%s is not an Monitor sort field", field)
		}
		sorts[i] = types.MonitorBy(types.MonitorField(field), types.SortOrder(sortSpec.Order[i]))
	}

	if len(sorts) > 0 {
		sort.SliceStable(monitors, types.MonitorCmp(monitors, sorts...))
	}
	return nil
}

func SortNames(names []types.Name, sortSpec SortSpec) error {
	if len(sortSpec.Fields) != len(sortSpec.Order) {
		return fmt.Errorf("fields and order must have the same length")
	}

	sorts := make([]func(p1, p2 types.Name) bool, len(sortSpec.Fields))
	for i, field := range sortSpec.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(types.GetSortFieldsName(), field) {
			return fmt.Errorf("%s is not an Name sort field", field)
		}
		sorts[i] = types.NameBy(types.NameField(field), types.SortOrder(sortSpec.Order[i]))
	}

	if len(sorts) > 0 {
		sort.SliceStable(names, types.NameCmp(names, sorts...))
	}
	return nil
}

func SortChunkRecords(chunkrecords []types.ChunkRecord, sortSpec SortSpec) error {
	if len(sortSpec.Fields) != len(sortSpec.Order) {
		return fmt.Errorf("fields and order must have the same length")
	}

	sorts := make([]func(p1, p2 types.ChunkRecord) bool, len(sortSpec.Fields))
	for i, field := range sortSpec.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(types.GetSortFieldsChunkRecord(), field) {
			return fmt.Errorf("%s is not an ChunkRecord sort field", field)
		}
		sorts[i] = types.ChunkRecordBy(types.ChunkRecordField(field), types.SortOrder(sortSpec.Order[i]))
	}

	if len(sorts) > 0 {
		sort.SliceStable(chunkrecords, types.ChunkRecordCmp(chunkrecords, sorts...))
	}
	return nil
}

func SortChunkStats(chunkstats []types.ChunkStats, sortSpec SortSpec) error {
	if len(sortSpec.Fields) != len(sortSpec.Order) {
		return fmt.Errorf("fields and order must have the same length")
	}

	sorts := make([]func(p1, p2 types.ChunkStats) bool, len(sortSpec.Fields))
	for i, field := range sortSpec.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(types.GetSortFieldsChunkStats(), field) {
			return fmt.Errorf("%s is not an ChunkStats sort field", field)
		}
		sorts[i] = types.ChunkStatsBy(types.ChunkStatsField(field), types.SortOrder(sortSpec.Order[i]))
	}

	if len(sorts) > 0 {
		sort.SliceStable(chunkstats, types.ChunkStatsCmp(chunkstats, sorts...))
	}
	return nil
}

func SortCacheItems(cacheitems []types.CacheItem, sortSpec SortSpec) error {
	if len(sortSpec.Fields) != len(sortSpec.Order) {
		return fmt.Errorf("fields and order must have the same length")
	}

	sorts := make([]func(p1, p2 types.CacheItem) bool, len(sortSpec.Fields))
	for i, field := range sortSpec.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(types.GetSortFieldsCacheItem(), field) {
			return fmt.Errorf("%s is not an CacheItem sort field", field)
		}
		sorts[i] = types.CacheItemBy(types.CacheItemField(field), types.SortOrder(sortSpec.Order[i]))
	}

	if len(sorts) > 0 {
		sort.SliceStable(cacheitems, types.CacheItemCmp(cacheitems, sorts...))
	}
	return nil
}

func SortChains(chains []types.Chain, sortSpec SortSpec) error {
	if len(sortSpec.Fields) != len(sortSpec.Order) {
		return fmt.Errorf("fields and order must have the same length")
	}

	sorts := make([]func(p1, p2 types.Chain) bool, len(sortSpec.Fields))
	for i, field := range sortSpec.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(types.GetSortFieldsChain(), field) {
			return fmt.Errorf("%s is not an Chain sort field", field)
		}
		sorts[i] = types.ChainBy(types.ChainField(field), types.SortOrder(sortSpec.Order[i]))
	}

	if len(sorts) > 0 {
		sort.SliceStable(chains, types.ChainCmp(chains, sorts...))
	}
	return nil
}

func SortAbis(abis []types.Abi, sortSpec SortSpec) error {
	if len(sortSpec.Fields) != len(sortSpec.Order) {
		return fmt.Errorf("fields and order must have the same length")
	}

	sorts := make([]func(p1, p2 types.Abi) bool, len(sortSpec.Fields))
	for i, field := range sortSpec.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(types.GetSortFieldsAbi(), field) {
			return fmt.Errorf("%s is not an Abi sort field", field)
		}
		sorts[i] = types.AbiBy(types.AbiField(field), types.SortOrder(sortSpec.Order[i]))
	}

	if len(sorts) > 0 {
		sort.SliceStable(abis, types.AbiCmp(abis, sorts...))
	}
	return nil
}

func SortFunctions(functions []types.Function, sortSpec SortSpec) error {
	if len(sortSpec.Fields) != len(sortSpec.Order) {
		return fmt.Errorf("fields and order must have the same length")
	}

	sorts := make([]func(p1, p2 types.Function) bool, len(sortSpec.Fields))
	for i, field := range sortSpec.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(types.GetSortFieldsFunction(), field) {
			return fmt.Errorf("%s is not an Function sort field", field)
		}
		sorts[i] = types.FunctionBy(types.FunctionField(field), types.SortOrder(sortSpec.Order[i]))
	}

	if len(sorts) > 0 {
		sort.SliceStable(functions, types.FunctionCmp(functions, sorts...))
	}
	return nil
}

// EXISTING_CODE
// EXISTING_CODE
