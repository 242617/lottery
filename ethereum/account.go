package ethereum

import (
	"encoding/hex"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	Address    string
	PrivateKey string
}

func generateAccount() Account {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Println("err", err)
		return generateAccount()
	}
	return Account{
		Address:    strings.ToLower(crypto.PubkeyToAddress(key.PublicKey).Hex()),
		PrivateKey: "0x" + strings.ToLower(hex.EncodeToString(key.D.Bytes())),
	}
}

func NewGenerator() chan Account {
	ch := make(chan Account)
	go func() {
		for {
			ch <- generateAccount()
		}
	}()
	return ch
}
