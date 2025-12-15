package onchain

import (
	"github.com/oddsandprayers/polymarket/client/onchain/condition"
)

func (o *Onchain) Condition() *condition.Condition {
	return o.con
}
