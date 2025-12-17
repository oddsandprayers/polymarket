package event

import (
	"net/url"
	"strconv"
	"time"

	"github.com/oddsandprayers/polymarket/request"
)

// Option defines the URL query parameters for searching events according to
// https://docs.polymarket.com/api-reference/events/list-events.
type Option struct {
	// Tag is the tag_id URL query parameter.
	Tag string
	// Clo is the closed URL query parameter.
	Clo bool
	// Ord is the order URL query parameter.
	Ord string
	// Asc is the ascending URL query parameter.
	Asc bool
	// Sta is the start_date_min URL query parameter.
	Sta time.Time
	// Lim is the limit URL query parameter.
	Lim int
	// Off is the offset URL query parameter.
	Off int
}

func (o Option) Query() url.Values {
	// Note that we have to manually initialize url.Values, because that type is
	// actually a map.

	var qry url.Values
	{
		qry = url.Values{}
	}

	if o.Tag != "" {
		qry.Add("tag_id", o.Tag)
	}

	{
		qry.Add("closed", strconv.FormatBool(o.Clo))
	}

	if o.Ord != "" {
		qry.Add("order", o.Ord)
	}

	{
		qry.Add("ascending", strconv.FormatBool(o.Asc))
	}

	if !o.Sta.IsZero() {
		qry.Add("start_date_min", o.Sta.Format(request.Layout))
	}

	if o.Lim != 0 {
		qry.Add("limit", strconv.Itoa(o.Lim))
	}

	if o.Off != 0 {
		qry.Add("offset", strconv.Itoa(o.Off))
	}

	return qry
}
