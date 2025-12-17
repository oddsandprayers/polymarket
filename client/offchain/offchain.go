package offchain

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/oddsandprayers/polymarket/client/offchain/event"
	"github.com/oddsandprayers/polymarket/client/offchain/token"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Cli is the required HTTP client being used to handle requests and
	// responses. If this value is empty, then the client creation will panic.
	Cli *http.Client

	// Url is the required endpoint of the underlying Polymarket API. If this
	// value is empty, then the client creation will panic.
	Url string
}

type Offchain struct {
	eve *event.Event
	tok *token.Token
}

func New(c Config) *Offchain {
	if c.Cli == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cli must not be empty", c)))
	}
	if c.Url == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Url must not be empty", c)))
	}

	var req *Request
	{
		req = &Request{
			cli: c.Cli,
			url: strings.TrimSuffix(strings.TrimSpace(c.Url), "/"),
		}
	}

	return &Offchain{
		eve: event.New(event.Config{Req: req}),
		tok: token.New(token.Config{Req: req}),
	}
}
