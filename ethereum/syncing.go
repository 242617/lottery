package ethereum

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

func (c *client) Syncing(ctx context.Context) (bool, error) {
	var result interface{}
	err := c.client.CallContext(ctx, &result, "eth_syncing")
	if err != nil {
		log.Error().Err(err).Msg("cannot call syncing")
		return false, err
	}

	switch res := result.(type) {
	case map[string]interface{}:
		// struct {
		// 	CurrentBlock  string `json:"currentBlock"`
		// 	HighestBlock  string `json:"highestBlock"`
		// 	KnownStates   string `json:"knownStates"`
		// 	PulledStates  string `json:"pulledStates"`
		// 	StartingBlock string `json:"startingBlock"`
		// }
		return true, nil
	case bool:
		return res, nil
	default:
		return false, errors.New("unknown state")
	}
}

func (c *client) SyncingCh() chan struct{} {
	ch := make(chan struct{})
	go func() {
		var syncing bool
		var err error
		for syncing = true; syncing; syncing, _ = c.Syncing(context.Background()) {
			if err != nil {
				log.Error().Err(err).Msg("cannot get syncing")
			}
			time.Sleep(5 * time.Second)
		}
		ch <- struct{}{}
	}()
	return ch
}
