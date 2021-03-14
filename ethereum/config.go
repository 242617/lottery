package ethereum

type Config struct {
	NodeAddress string `yaml:"node_address"`
	GasLimit    uint64 `yaml:"gas_limit"`
}
