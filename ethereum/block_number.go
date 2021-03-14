package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

func (c *client) BlockNumber(ctx context.Context) (uint64, error) {
	var header *types.Header
	err := c.client.CallContext(ctx, &header, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		log.Error().Err(err).Msg("cannot get block by number")
		return 0, err
	}
	if err == nil && header == nil {
		return 0, ethereum.NotFound
	}

	return header.Number.Uint64(), nil
}
