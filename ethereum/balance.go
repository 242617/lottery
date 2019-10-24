package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Balance(ctx context.Context, address string) (*big.Int, error) {

	var result hexutil.Big
	err := client.CallContext(ctx, &result, "eth_getBalance", address, "latest")
	if err != nil {
		return nil, err
	}

	return result.ToInt(), nil
}
