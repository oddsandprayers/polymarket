package request

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/xh3b4sd/tracer"
)

func Request(cli *http.Client, req *http.Request, bod any) error {
	var err error

	// Ensure the request does not block forever, and then make sure to call the
	// cancel function at the end of request().

	var ctx context.Context
	var can context.CancelFunc
	{
		ctx, can = context.WithTimeout(context.Background(), 10*time.Second)
	}

	{
		defer can()
	}

	{
		req = req.WithContext(ctx)
	}

	// Do the actual HTTP request to the Indexing Co API.

	var res *http.Response
	{
		res, err = cli.Do(req)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		defer res.Body.Close() // nolint:errcheck
	}

	// Handle unwanted status codes using a typed error and wrap any failures with
	// additional context.

	if res.StatusCode >= 300 {
		byt, err := io.ReadAll(res.Body)
		if err != nil {
			return tracer.Mask(err)
		}

		return tracer.Mask(statusCodeError,
			tracer.Context{Key: "url", Value: req.URL.String()},
			tracer.Context{Key: "method", Value: req.Method},
			tracer.Context{Key: "code", Value: res.Status},
			tracer.Context{Key: "body", Value: string(byt)},
		)
	}

	// At this point we got a valid response back and can simply decode the
	// response body into the provided pointer.

	{
		err = json.NewDecoder(res.Body).Decode(bod)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
