package event

import (
	"fmt"

	"github.com/oddsandprayers/polymarket/request"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Req request.Interface
}

type Event struct {
	req request.Interface
}

func New(c Config) *Event {
	if c.Req == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Req must not be empty", c)))
	}

	return &Event{
		req: c.Req,
	}
}
