package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var Config struct {
	LogPrefix     string `yaml:"log_prefix"`
	NodeAddress   string `yaml:"node_address"`
	NodeSecret    string `yaml:"node_secret"`
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
