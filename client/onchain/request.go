package onchain

import (
	"fmt"
	"net/http"

	"github.com/oddsandprayers/polymarket/request"
	"github.com/xh3b4sd/tracer"
)

type Request struct {
	cli *http.Client
	key string
	url string
}

func (r *Request) Endpoint() string {
	return r.url
}

func (r *Request) Request(req *http.Request, bod any) error {
	var err error

	// Ensure our requests are authenticated with a bearer token.
	//
	//     Authorization: Bearer ${SUBGRAPH_API_KEY}
	//

	{
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.key))
	}

	{
		err = request.Request(r.cli, req, bod)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
