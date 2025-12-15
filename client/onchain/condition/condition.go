package condition

import (
	"embed"
	"fmt"

	"github.com/oddsandprayers/polymarket/request"
	"github.com/xh3b4sd/tracer"
)

//go:embed query.graphql
var emb embed.FS

type Config struct {
	Req request.Interface
}

type Condition struct {
	qry string
	req request.Interface
}

func New(c Config) *Condition {
	if c.Req == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Req must not be empty", c)))
	}

	byt, err := emb.ReadFile("query.graphql")
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return &Condition{
		qry: string(byt),
		req: c.Req,
	}
}
