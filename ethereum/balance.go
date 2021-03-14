package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rs/zerolog/log"
)

func (c *client) Balance(ctx context.Context, address string) (*big.Int, error) {
	var result hexutil.Big
	err := c.client.CallContext(ctx, &result, "eth_getBalance", address, "latest")
	if err != nil {
		log.Error().Err(err).Msg("cannot get balance")
		return nil, err
	}

	return result.ToInt(), nil
}
