package client

import (
	"fmt"
	"net/http"
	"os"

	"github.com/oddsandprayers/polymarket/client/offchain"
	"github.com/oddsandprayers/polymarket/client/onchain"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Cli is the optional HTTP client being used to handle requests and
	// responses. Defaults to http.DefaultClient.
	Cli *http.Client

	// Key is the required API key to work with the underlying Subgraph API. If
	// this value is empty, then the client will try to find an API key within the
	// process environment using the $POLYMARKET_SUBGRAPH_API_KEY environment
	// variable. If that lookup does then also not yield a non-empty string, then
	// the client creation will panic.
	//
	//     https://thegraph.com/studio/apikeys
	//
	Key string
}

type Client struct {
	onc *onchain.Onchain
	off *offchain.Offchain
}

func New(c Config) *Client {
	if c.Cli == nil {
		c.Cli = http.DefaultClient
	}
	if c.Key == "" {
		c.Key = os.Getenv(string(SubgraphApiKey))
	}
	if c.Key == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Key must not be empty", c)))
	}

	return &Client{
		onc: onchain.New(onchain.Config{
			Cli: c.Cli,
			Key: c.Key,
			Url: string(SubgraphApiUrl),
		}),
		off: offchain.New(offchain.Config{
			Cli: c.Cli,
			// NOTE Polymarket has no API auth
			Url: string(PolymarketApiUrl),
		}),
	}
}
