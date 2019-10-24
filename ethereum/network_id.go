package ethereum

import (
	"context"
	"fmt"
	"math/big"
)

func NetworkID(ctx context.Context) (*big.Int, error) {

	var ver string
	err := client.CallContext(ctx, &ver, "net_version")
	if err != nil {
		return nil, err
	}

	version := big.NewInt(0)
	_, ok := version.SetString(ver, 10)
	if !ok {
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}

	return version, nil
}
