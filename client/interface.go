package client

import (
	"github.com/oddsandprayers/polymarket/client/offchain"
	"github.com/oddsandprayers/polymarket/client/onchain"
)

type Interface interface {
	Offchain() *offchain.Offchain
	Onchain() *onchain.Onchain
}
