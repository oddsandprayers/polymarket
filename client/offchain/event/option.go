package event

import (
	"net/url"
	"strconv"
)

type Option struct {
	// Tag is the tag_id URL query parameter.
	Tag string
	// Clo is the closed URL query parameter.
	Clo bool
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

	if o.Lim != 0 {
		qry.Add("limit", strconv.Itoa(o.Lim))
	}

	if o.Off != 0 {
		qry.Add("offset", strconv.Itoa(o.Off))
	}

	return qry
}
