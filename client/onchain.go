package client

import "github.com/oddsandprayers/polymarket/client/onchain"

func (c *Client) Onchain() *onchain.Onchain {
	return c.onc
}
