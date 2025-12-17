package token

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/xh3b4sd/tracer"
)

// Option defines the URL query parameters for searching prices according to
// https://docs.polymarket.com/api-reference/pricing/get-price-history-for-a-traded-token.
type Option struct {
	// Tok is the market URL query parameter.
	Tok string
	// Sta is the startTs URL query parameter.
	Sta time.Time
	// End is the endTs URL query parameter.
	End time.Time
	// Fid is the fidelity URL query parameter.
	Fid fidelity
}

func (o Option) Query() url.Values {
	// Note that we have to manually initialize url.Values, because that type is
	// actually a map.

	var qry url.Values
	{
		qry = url.Values{}
	}

	{
		qry.Add("market", o.Tok)
	}

	{
		qry.Add("startTs", strconv.Itoa(int(o.Sta.Unix())))
	}

	{
		qry.Add("endTs", strconv.Itoa(int(o.End.Unix())))
	}

	{
		qry.Add("fidelity", strconv.Itoa(int(o.Fid)))
	}

	return qry
}

func (o Option) Verify() error {
	if o.Tok == "" {
		return tracer.Mask(fmt.Errorf("token ID must not be empty"))
	}

	if o.Sta.IsZero() {
		return tracer.Mask(fmt.Errorf("start time must not be empty"))
	}

	if o.End.IsZero() {
		return tracer.Mask(fmt.Errorf("end time must not be empty"))
	}

	if o.Fid == 0 {
		return tracer.Mask(fmt.Errorf("fidelity must not be empty"))
	}

	return nil
}
