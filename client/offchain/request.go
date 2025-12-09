package offchain

import (
	"net/http"

	"github.com/oddsandprayers/polymarket/request"
	"github.com/xh3b4sd/tracer"
)

type Request struct {
	cli *http.Client
	url string
}

func (r *Request) Endpoint() string {
	return r.url
}

func (r *Request) Request(req *http.Request, bod any) error {
	var err error

	{
		err = request.Request(r.cli, req, bod)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
