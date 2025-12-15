package onchain

import (
	"github.com/oddsandprayers/polymarket/client/onchain/condition"
)

type Interface interface {
	Condition() *condition.Condition
}
