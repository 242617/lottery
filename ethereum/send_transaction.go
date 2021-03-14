package ethereum

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/rs/zerolog/log"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

func (c *client) SendTransaction(ctx context.Context, cfg Config, account Account, targetAddress string, nonce uint64, amount *big.Int, gasPrice *big.Int) (string, error) {

	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(targetAddress),
		amount,
		cfg.GasLimit,
		gasPrice,
		nil,
	)

	privateKey, err := crypto.HexToECDSA(account.PrivateKey[2:])
	if err != nil {
		log.Error().Err(err).Msg("cannot convert hex to ecdsa")
		return "", err
	}

	chainID, err := c.NetworkID(ctx)
	if err != nil {
		log.Error().Err(err).Msg("cannot get network id")
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error().Err(err).Msg("cannot sign transaction")
		return "", err
	}

	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Error().Err(err).Msg("cannot encode to bytes")
		return "", err
	}

	err = c.client.CallContext(ctx, nil, "eth_sendRawTransaction", common.ToHex(data))
	if err != nil {
		log.Error().Err(err).Msg("cannot send raw transaction")
		if err.Error() == "insufficient funds for gas * price + value" {
			return "", ErrInsufficientFunds
		}
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
