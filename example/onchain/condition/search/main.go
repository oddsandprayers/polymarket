package main

import (
	"encoding/json"
	"fmt"

	"github.com/oddsandprayers/polymarket/client"
	"github.com/oddsandprayers/polymarket/client/onchain/condition"
	"github.com/xh3b4sd/tracer"
)

//     curl -s -X POST \
//       -H "Content-Type: application/json" \
//       -H "Authorization: Bearer ${SUBGRAPH_API_KEY}" \
//       -d '{
//         "query": "query GetConditions($ids: [Bytes!]!) { conditions(where: { id_in: $ids }) { id outcomeSlotCount payoutNumerators payoutDenominator } }",
//         "variables": {
//           "ids": [
//             "0xe3b423dfad8c22ff75c9899c4e8176f628cf4ad4caa00481764d320e7415f7a9",
//             "0x9b946f54f3428aafc308c33aa04a943fe13a011bdac9a9b66e1ba16c416ca256",
//             "0x44f10d1cd5aaed4b7ae0b5edb76790f54f45dc0bcaa86831c83d865c774fbb90",
//             "0x3e0524de013d9dc359f5eb370773f25de2f03d3200294cfd0fa7dac2f399d101",
//             "0x5d1a1ab716fd06943441fe27cde0089651ce769bec55e191b6953468a0e9f0d0"
//           ]
//         }
//       }' \
//       https://gateway.thegraph.com/api/subgraphs/id/81Dm16JjuFSrqz813HysXoUPvzTwE7fsfPk2RTf66nyC | jq .

//
//     go run example/onchain/condition/search/main.go | jq .
//
//     [
//       {
//         "id": "0x3e0524de013d9dc359f5eb370773f25de2f03d3200294cfd0fa7dac2f399d101",
//         "outcomeSlotCount": 2,
//         "payoutDenominator": "1000000000000000000",
//         "payoutNumerators": [
//           "1000000000000000000",
//           "0"
//         ]
//       },
//       {
//         "id": "0x44f10d1cd5aaed4b7ae0b5edb76790f54f45dc0bcaa86831c83d865c774fbb90",
//         "outcomeSlotCount": 2,
//         "payoutDenominator": "1000000000000000000",
//         "payoutNumerators": [
//           "1000000000000000000",
//           "0"
//         ]
//       },
//       {
//         "id": "0x5d1a1ab716fd06943441fe27cde0089651ce769bec55e191b6953468a0e9f0d0",
//         "outcomeSlotCount": 2,
//         "payoutDenominator": "1000000000000000000",
//         "payoutNumerators": [
//           "0",
//           "1000000000000000000"
//         ]
//       },
//       {
//         "id": "0x9b946f54f3428aafc308c33aa04a943fe13a011bdac9a9b66e1ba16c416ca256",
//         "outcomeSlotCount": 2,
//         "payoutDenominator": "1000000000000000000",
//         "payoutNumerators": [
//           "0",
//           "1000000000000000000"
//         ]
//       },
//       {
//         "id": "0xe3b423dfad8c22ff75c9899c4e8176f628cf4ad4caa00481764d320e7415f7a9",
//         "outcomeSlotCount": 2,
//         "payoutDenominator": "1000000000000000000",
//         "payoutNumerators": [
//           "0",
//           "1000000000000000000"
//         ]
//       }
//     ]
//

func main() {
	var err error

	var cli *client.Client
	{
		cli = client.New(client.Config{})
	}

	var opt condition.Option
	{
		opt = condition.Option{
			Ids: []string{
				"0xe3b423dfad8c22ff75c9899c4e8176f628cf4ad4caa00481764d320e7415f7a9",
				"0x9b946f54f3428aafc308c33aa04a943fe13a011bdac9a9b66e1ba16c416ca256",
				"0x44f10d1cd5aaed4b7ae0b5edb76790f54f45dc0bcaa86831c83d865c774fbb90",
				"0x3e0524de013d9dc359f5eb370773f25de2f03d3200294cfd0fa7dac2f399d101",
				"0x5d1a1ab716fd06943441fe27cde0089651ce769bec55e191b6953468a0e9f0d0",
			},
		}
	}

	var res []condition.Object
	{
		res, err = cli.Onchain().Condition().Search(opt)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var byt []byte
	{
		byt, err = json.MarshalIndent(res, "", "  ")
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		fmt.Printf("%s\n", byt)
	}
}
