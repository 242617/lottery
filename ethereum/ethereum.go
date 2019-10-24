package ethereum

import (
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/242617/lottery/config"
)

var client *rpc.Client

func Init() error {

	var err error
	client, err = rpc.DialHTTPWithClient(config.Config.NodeAddress, &http.Client{Transport: &transport{}})
	if err != nil {
		return err
	}

	return nil
}

type transport struct{}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth("", config.Config.NodeSecret)
	return http.DefaultTransport.RoundTrip(req)
}

func Close() {
	client.Close()
}
