package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var Config struct {
	IPC           string `yaml:"ipc"`
	Streams       int    `yaml:"streams"`
	TargetAddress string `yaml:"target_address"`
	GasLimit      uint64 `yaml:"gas_limit"`
}

func Init(filename string) error {

	barr, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(barr, &Config)
	if err != nil {
		return err
	}

	return nil
}
