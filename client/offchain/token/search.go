package token

import (
	"fmt"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

func (t *Token) Search(opt Option) (Object, error) {
	var err error

	{
		err = opt.Verify()
		if err != nil {
			return Object{}, tracer.Mask(err)
		}
	}

	var pat string
	{
		pat = fmt.Sprintf(
			t.req.Endpoint(),
			"clob",
			"prices-history?"+opt.Query().Encode(),
		)
	}

	var req *http.Request
	{
		req, err = http.NewRequest(http.MethodGet, pat, nil)
		if err != nil {
			return Object{}, tracer.Mask(err)
		}
	}

	{
		req.Header.Set("content-type", "application/json")
	}

	var res Object
	{
		err = t.req.Request(req, &res)
		if err != nil {
			return Object{}, tracer.Mask(err)
		}
	}

	{
		res.ID = opt.Tok
	}

	return res, nil
}
