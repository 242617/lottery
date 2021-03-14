package main

import (
	"context"
	"math/big"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"

	"github.com/242617/lottery/ethereum"
)

const (
	TimeFormat  = "15:04:05"
	ServiceName = "lottery"
)

var (
	n uint64
	l sync.RWMutex
)
var config struct {
	Ethereum      ethereum.Config `yaml:"ethereum"`
	Streams       int             `yaml:"streams"`
	TargetAddress string          `yaml:"target_address"`
	StatPeriod    time.Duration   `yaml:"stat_period"`
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{
		TimeFormat: TimeFormat,
		Out:        os.Stderr,
	})
}

func main() {
	if os.Getenv("CONSUL_HTTP_TOKEN") == "" {
		log.Fatal().Msg("empty token")
	}

	client, err := api.NewClient(&api.Config{Address: "0.0.0.0:8500"})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create client")
	}

	service, err := connect.NewService(ServiceName, client)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new service")
	}
	defer service.Close()

	<-service.ReadyWait()

	log.Info().Msg("start lottery")

	pair, _, err := client.KV().Get(ServiceName, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get kv")
	}

	if err := yaml.Unmarshal(pair.Value, &config); err != nil {
		log.Fatal().Err(err).Msg("cannot unmarshal config")
	}

	eth, err := ethereum.New(config.Ethereum)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot init ethereum")
	}
	defer eth.Close()

	<-eth.SyncingCh()

	balance, err := eth.Balance(context.TODO(), config.TargetAddress)
	if err != nil {
		log.Error().Err(err).Msg("cannot get balance")
		return
	}
	log.Info().
		Str("balance.String()", balance.String()).
		Str("config.TargetAddress", config.TargetAddress).
		Msg("check")

	generator := ethereum.NewGenerator()
	for s := 0; s <= config.Streams; s++ {
		go func() {
			for {
				account := <-generator

				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				defer cancel()

				balance, err := eth.Balance(ctx, account.Address)
				if err != nil {
					log.Error().Err(err).Msg("cannot get balance")
					time.Sleep(time.Second)
					continue
				}
				if balance.Cmp(big.NewInt(0)) != 1 {
					l.Lock()
					n++
					l.Unlock()
					continue
				}

				log.Info().Str("account.Address", account.Address).Msg("bingo!")

				nonce, err := eth.Nonce(ctx, account.Address)
				if err != nil {
					log.Error().Err(err).Msg("cannot get nonce")
					continue
				}

				gasPrice, err := eth.GasPrice(ctx)
				if err != nil {
					log.Error().Err(err).Msg("cannot get gas price")
					continue
				}

				amount := big.NewInt(0)
				amount.Sub(
					balance,
					big.NewInt(0).Mul(
						gasPrice,
						big.NewInt(int64(config.Ethereum.GasLimit)),
					),
				)

				hash, err := eth.SendTransaction(ctx, config.Ethereum, account, config.TargetAddress, nonce, amount, gasPrice)
				if err != nil {
					log.Error().Err(err).Msg("cannot send transaction")
					continue
				}

				log.Info().
					Str("hash", hash).
					Int64("balance", balance.Int64()).
					Str("address", account.Address).
					Int64("amount", amount.Int64()).
					Str("private key", account.PrivateKey).
					Msg("success!")

			}
		}()
	}

	start := time.Now()
	go func() {
		for {
			time.Sleep(config.StatPeriod)
			l.RLock()
			log.Debug().
				Dur("time", time.Since(start)).
				Uint64("tries", n).
				Msg("stat")
			l.RUnlock()
		}
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	log.Info().Msg("stop lottery")
}
