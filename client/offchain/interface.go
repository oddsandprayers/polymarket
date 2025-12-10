package offchain

import "github.com/oddsandprayers/polymarket/client/offchain/event"

type Interface interface {
	Event() *event.Event
}
