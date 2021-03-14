package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rs/zerolog/log"
)

func (c *client) Nonce(ctx context.Context, address string) (uint64, error) {
	var result hexutil.Uint64
	err := c.client.CallContext(ctx, &result, "eth_getTransactionCount", address, "latest")
	if err != nil {
		log.Error().Msg("cannot get transaction count")
		return 0, err
	}

	return uint64(result), nil
}
