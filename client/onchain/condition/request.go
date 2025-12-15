package condition

//	{
//	  "query": "query GetConditions($ids: [Bytes!]!) { conditions(where: { id_in: $ids }) { id outcomeSlotCount payoutNumerators payoutDenominator } }",
//	  "variables": {
//	    "ids": [
//	      "0xe3b423dfad8c22ff75c9899c4e8176f628cf4ad4caa00481764d320e7415f7a9",
//	      "0x9b946f54f3428aafc308c33aa04a943fe13a011bdac9a9b66e1ba16c416ca256",
//	      "0x44f10d1cd5aaed4b7ae0b5edb76790f54f45dc0bcaa86831c83d865c774fbb90",
//	      "0x3e0524de013d9dc359f5eb370773f25de2f03d3200294cfd0fa7dac2f399d101",
//	      "0x5d1a1ab716fd06943441fe27cde0089651ce769bec55e191b6953468a0e9f0d0"
//	    ]
//	  }
//	}

type Request struct {
	Query     string    `json:"query"`
	Variables Variables `json:"variables"`
}

type Variables struct {
	Ids []string `json:"ids"`
}
