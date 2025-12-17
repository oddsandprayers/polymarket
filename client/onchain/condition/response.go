package condition

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Conditions []Object `json:"conditions"`
}

type Object struct {
	ID       string   `json:"id"`
	Slots    int      `json:"outcomeSlotCount"`
	Payout   string   `json:"payoutDenominator"`
	Outcomes []string `json:"payoutNumerators"`
}

// Winner returns the slot index of a condition object that was resolved to be
// true. In the example below, Winner returns 0, because the payout denominator
// "1" has the index 0 in the payout numerators array.
//
//	{
//	    "id": "0x6153e5def1788e35bc42be5120358ff3d440b655b25410148f53f2b4283a4977",
//	    "outcomeSlotCount": 2,
//	    "payoutDenominator": "1",
//	    "payoutNumerators": [
//	        "1",
//	        "0"
//	    ]
//	}
func (o Object) Winner() int {
	if o.Payout == "" {
		return -1
	}

	if o.Slots != len(o.Outcomes) {
		return -1
	}

	for i, x := range o.Outcomes {
		if x == o.Payout {
			return i
		}
	}

	return -1
}
