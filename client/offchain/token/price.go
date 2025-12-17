package token

import (
	"fmt"

	"github.com/oddsandprayers/polymarket/request"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Req request.Interface
}

type Token struct {
	req request.Interface
}

func New(c Config) *Token {
	if c.Req == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Req must not be empty", c)))
	}

	return &Token{
		req: c.Req,
	}
}
