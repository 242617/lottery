package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

var ErrNotEnoughArguments = errors.New("not enough arguments")

func main() {
	if len(os.Args) < 2 {
		log.Fatal(ErrNotEnoughArguments)
	}

	private := os.Args[1]

	key, err := crypto.HexToECDSA(private[2:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.ToLower(crypto.PubkeyToAddress(key.PublicKey).String()))
}
