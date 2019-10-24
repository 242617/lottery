package ethereum

import (
	"context"
	"log"
	"time"
)

func AfterBlocks(ctx context.Context, blocks uint64) chan struct{} {
	ch := make(chan struct{})
	go func() {
		blockNumber, err := BlockNumber(ctx)
		if err != nil {
			log.Println("err", err)
			return
		}
		for {
			currentBlockNumber, err := BlockNumber(ctx)
			if err != nil {
				log.Println("err", err)
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
