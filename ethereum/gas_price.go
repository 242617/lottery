package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rs/zerolog/log"
)

func (c *client) GasPrice(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := c.client.CallContext(ctx, &result, "eth_gasPrice")
	if err != nil {
		log.Error().Err(err).Msg("cannot get gas price")
		return nil, err
	}

	return result.ToInt(), nil
}
