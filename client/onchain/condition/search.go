package condition

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

func (e *Condition) Search(opt Option) ([]Object, error) {
	var err error

	var bod Request
	{
		bod = Request{
			Query:     e.qry,
			Variables: Variables(opt),
		}
	}

	var pat string
	{
		pat = e.req.Endpoint()
	}

	var byt []byte
	{
		byt, err = json.Marshal(bod)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var req *http.Request
	{
		req, err = http.NewRequest(http.MethodPost, pat, bytes.NewReader(byt))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		req.Header.Set("content-type", "application/json")
	}

	var res Response
	{
		err = e.req.Request(req, &res)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range res.Data.Conditions {
		if x.ID == "" {
			return nil, tracer.Mask(searchError, tracer.Context{Key: "reason", Value: "condition ID must not be empty"})
		}
	}

	return res.Data.Conditions, nil
}
