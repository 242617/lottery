package ethereum

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/242617/lottery/config"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

func SendTransaction(ctx context.Context, account Account, targetAddress string, nonce uint64, amount *big.Int, gasPrice *big.Int) (string, error) {

	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(targetAddress),
		amount,
		config.Config.GasLimit,
		gasPrice,
		nil,
	)

	privateKey, err := crypto.HexToECDSA(account.PrivateKey[2:])
	if err != nil {
		return "", err
	}

	chainID, err := NetworkID(ctx)
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", err
	}

	err = client.CallContext(ctx, nil, "eth_sendRawTransaction", common.ToHex(data))
	if err != nil {
		if err.Error() == "insufficient funds for gas * price + value" {
			return "", ErrInsufficientFunds
		}
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
