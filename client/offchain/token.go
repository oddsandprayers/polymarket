package offchain

import "github.com/oddsandprayers/polymarket/client/offchain/token"

func (o *Offchain) Token() *token.Token {
	return o.tok
}
