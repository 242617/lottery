package ethereum

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

func (c *client) AfterBlocks(ctx context.Context, blocks uint64) chan struct{} {
	ch := make(chan struct{})
	go func() {
		blockNumber, err := c.BlockNumber(ctx)
		if err != nil {
			log.Error().Err(err).Msg("cannot get block number")
			return
		}
		for {
			currentBlockNumber, err := c.BlockNumber(ctx)
			if err != nil {
				log.Error().Err(err).Msg("cannot get block number")
				break
			}
			if currentBlockNumber >= blockNumber+blocks {
				ch <- struct{}{}
				break
			}
			time.Sleep(time.Second)
		}
	}()
	return ch
}
