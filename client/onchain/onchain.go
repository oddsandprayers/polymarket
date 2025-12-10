package onchain

import (
	"fmt"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Cli is the required HTTP client being used to handle requests and
	// responses. If this value is empty, then the client creation will panic.
	Cli *http.Client

	// Key is the required API key to work with the underlying Subgraph API. If
	// this value is empty, then the client creation will panic.
	Key string

	// Url is the required endpoint of the underlying Subgraph API. If this value
	// is empty, then the client creation will panic.
	Url string
}

type Onchain struct{}

func New(c Config) *Onchain {
	if c.Cli == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cli must not be empty", c)))
	}
	if c.Key == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Key must not be empty", c)))
	}
	if c.Url == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Url must not be empty", c)))
	}

	// TODO add e.g. Condition service
	// var req *Request
	// {
	// 	req = &Request{
	// 		cli: c.Cli,
	// 		key: c.Key,
	// 		url: strings.TrimSuffix(strings.TrimSpace(c.Url), "/"),
	// 	}
	// }

	return &Onchain{}
}
