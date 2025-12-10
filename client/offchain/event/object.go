package event

import (
	"time"

	"github.com/oddsandprayers/polymarket/encoding"
)

type Object struct {
	ID          string    `json:"id"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Start       time.Time `json:"startDate"`
	End         time.Time `json:"endDate"`

	Volume float64 `json:"volume"`

	Markets []Market `json:"markets"`
}

type Market struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Question    string `json:"question"`
	Condition   string `json:"conditionId"`
	Description string `json:"description"`
	Option      string `json:"groupItemTitle"`
	Status      string `json:"umaResolutionStatus"`

	Volume float64 `json:"volumeNum"`

	Outcomes encoding.Strings `json:"outcomes"`
	Prices   encoding.Strings `json:"outcomePrices"`
	Tokens   encoding.Strings `json:"clobTokenIds"`
}
