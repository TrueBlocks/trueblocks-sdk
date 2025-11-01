package sdk

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/v6/pkg/rpc"

func PingRpc(providerUrl string) (*rpc.PingResult, error) {
	result, err := rpc.PingRpc(providerUrl)
	return result, err
}
