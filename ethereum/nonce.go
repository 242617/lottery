package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Nonce(ctx context.Context, address string) (uint64, error) {

	var result hexutil.Uint64
	err := client.CallContext(ctx, &result, "eth_getTransactionCount", address, "latest")
	if err != nil {
		return 0, err
	}

	return uint64(result), nil
}
