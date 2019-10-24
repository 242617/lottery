package ethereum

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

var client *rpc.Client

func Init(ipc string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var err error
	client, err = rpc.DialContext(ctx, ipc)
	if err != nil {
		return err
	}

	return nil
}
