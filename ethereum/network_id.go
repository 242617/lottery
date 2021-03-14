package ethereum

import (
	"context"
	"fmt"
	"math/big"

	"github.com/rs/zerolog/log"
)

func (c *client) NetworkID(ctx context.Context) (*big.Int, error) {
	var ver string
	err := c.client.CallContext(ctx, &ver, "net_version")
	if err != nil {
		log.Error().Err(err).Msg("cannot get net version")
		return nil, err
	}

	version := big.NewInt(0)
	_, ok := version.SetString(ver, 10)
	if !ok {
		log.Error().Str("ver", ver).Msg("cannot set net version")
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}

	return version, nil
}
