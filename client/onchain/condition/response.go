package condition

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Conditions []Object `json:"conditions"`
}

type Object struct {
	Id       string   `json:"id"`
	Slots    int      `json:"outcomeSlotCount"`
	Payout   string   `json:"payoutDenominator"`
	Outcomes []string `json:"payoutNumerators"`
}
