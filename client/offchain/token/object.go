package token

import (
	"github.com/oddsandprayers/polymarket/encoding"
)

type Object struct {
	ID      string    `json:"id"`
	History []History `json:"history"`
}

type History struct {
	Time  encoding.Time `json:"t"`
	Price float64       `json:"p"`
}
