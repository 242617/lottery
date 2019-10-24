package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func GasPrice(ctx context.Context) (*big.Int, error) {

	var result hexutil.Big
	err := client.CallContext(ctx, &result, "eth_gasPrice")
	if err != nil {
		return nil, err
	}

	return result.ToInt(), nil
}
