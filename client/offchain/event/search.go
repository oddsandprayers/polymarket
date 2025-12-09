package event

import (
	"fmt"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

func (e *Event) Search(opt Option) ([]Object, error) {
	var err error

	var pat string
	{
		pat = fmt.Sprintf(
			e.req.Endpoint(),
			"gamma-api",
			"events?"+opt.Query().Encode(),
		)
	}

	var req *http.Request
	{
		req, err = http.NewRequest(http.MethodGet, pat, nil)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		req.Header.Set("content-type", "application/json")
	}

	var res []Object
	{
		err = e.req.Request(req, &res)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return res, nil
}
