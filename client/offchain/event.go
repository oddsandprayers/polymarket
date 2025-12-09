package offchain

import (
	"github.com/oddsandprayers/polymarket/client/offchain/event"
)

func (o *Offchain) Event() *event.Event {
	return o.eve
}
