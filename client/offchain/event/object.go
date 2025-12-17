package event

import (
	"github.com/oddsandprayers/polymarket/encoding"
)

type Object struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Description string `json:"description"`

	Volume float64 `json:"volume"`

	Markets []Market `json:"markets"`
}

type Market struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Question    string `json:"question"`
	Condition   string `json:"conditionId"`
	Description string `json:"description"`
	Title       string `json:"groupItemTitle"`
	Status      string `json:"umaResolutionStatus"`

	Start  encoding.Time `json:"gameStartTime"`
	Volume float64       `json:"volumeNum"`

	Outcomes encoding.Strings `json:"outcomes"`
	Prices   encoding.Strings `json:"outcomePrices"`
	Tokens   encoding.Strings `json:"clobTokenIds"`
}
