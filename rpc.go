package sdk

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpc"

func PingRpc(providerUrl string) error {
	return rpc.PingRpc(providerUrl)
}
