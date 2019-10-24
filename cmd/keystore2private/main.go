package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

var ErrNotEnoughArguments = errors.New("not enough arguments")

func main() {
	if len(os.Args) < 2 {
		log.Fatal(ErrNotEnoughArguments)
	}

	filename := os.Args[1]

	barr, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var password string
	if len(os.Args) > 2 {
		password = os.Args[2]
	}

	key, err := keystore.DecryptKey(barr, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(common.BigToHash(key.PrivateKey.D).Hex())
}
