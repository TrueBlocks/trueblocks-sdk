package sdk

import (
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/base"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/rpc"
)

func PingRpc(providerUrl string) (*rpc.PingResult, error) {
	result, err := rpc.PingRpc(providerUrl)
	return result, err
}

// EstimateGasAndPrice estimates gas and retrieves current gas price in a single batch RPC call
func EstimateGasAndPrice(chain string, from base.Address, to base.Address, data string, value *base.Wei) (estimatedGas base.Gas, gasPrice base.Gas, err error) {
	return rpc.EstimateGasAndPrice(chain, from, to, data, value)
}
