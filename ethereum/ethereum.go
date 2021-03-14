package ethereum

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rs/zerolog/log"
)

func New(cfg Config) (*client, error) {
	c := client{
		cfg: cfg,
	}
	var err error
	c.client, err = rpc.Dial(cfg.NodeAddress)
	if err != nil {
		log.Error().Str("cfg.NodeAddress", cfg.NodeAddress).Err(err).Msg("cannot dial")
		return nil, err
	}
	return &c, nil
}

type client struct {
	cfg    Config
	client *rpc.Client
}

func (c *client) Close() { c.client.Close() }
