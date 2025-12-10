package client

import "github.com/oddsandprayers/polymarket/client/offchain"

func (c *Client) Offchain() *offchain.Offchain {
	return c.off
}
