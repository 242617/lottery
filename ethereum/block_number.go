package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

func BlockNumber(ctx context.Context) (uint64, error) {

	var header *types.Header
	err := client.CallContext(ctx, &header, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		return 0, err
	}
	if err == nil && header == nil {
		return 0, ethereum.NotFound
	}

	return header.Number.Uint64(), nil
}
