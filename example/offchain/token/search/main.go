package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/oddsandprayers/polymarket/client"
	"github.com/oddsandprayers/polymarket/client/offchain/token"
	"github.com/xh3b4sd/tracer"
)

//
//     curl -s --request GET \
//          --url 'https://clob.polymarket.com/prices-history?market=51028631785140982248590109905860777819766751100592776255708383993117903047997&startTs=1765587600&endTs=1765674000&fidelity=60' \
//          | jq .
//

//
//     go run example/offchain/token/search/main.go | jq .
//
//     {
//       "id": "51028631785140982248590109905860777819766751100592776255708383993117903047997",
//       "history": [
//         {
//           "t": "2025-12-13T01:00:13Z",
//           "p": 0.205
//         },
//         {
//           "t": "2025-12-13T02:00:13Z",
//           "p": 0.205
//         },
//         {
//           "t": "2025-12-13T03:00:15Z",
//           "p": 0.205
//         },
//         {
//           "t": "2025-12-13T04:00:14Z",
//           "p": 0.205
//         },
//         {
//           "t": "2025-12-13T05:00:13Z",
//           "p": 0.195
//         },
//         {
//           "t": "2025-12-13T06:00:28Z",
//           "p": 0.2
//         },
//         {
//           "t": "2025-12-13T07:00:13Z",
//           "p": 0.2
//         },
//         {
//           "t": "2025-12-13T08:00:13Z",
//           "p": 0.2
//         },
//         {
//           "t": "2025-12-13T09:00:16Z",
//           "p": 0.2
//         },
//         {
//           "t": "2025-12-13T10:00:13Z",
//           "p": 0.2
//         },
//         {
//           "t": "2025-12-13T11:00:13Z",
//           "p": 0.2
//         },
//         {
//           "t": "2025-12-13T12:00:20Z",
//           "p": 0.2
//         },
//         {
//           "t": "2025-12-13T13:00:13Z",
//           "p": 0.2
//         },
//         {
//           "t": "2025-12-13T14:00:27Z",
//           "p": 0.195
//         },
//         {
//           "t": "2025-12-13T15:00:16Z",
//           "p": 0.185
//         },
//         {
//           "t": "2025-12-13T16:00:13Z",
//           "p": 0.185
//         },
//         {
//           "t": "2025-12-13T17:00:17Z",
//           "p": 0.185
//         },
//         {
//           "t": "2025-12-13T18:00:17Z",
//           "p": 0.185
//         },
//         {
//           "t": "2025-12-13T19:00:13Z",
//           "p": 0.185
//         },
//         {
//           "t": "2025-12-13T20:00:14Z",
//           "p": 0.18
//         },
//         {
//           "t": "2025-12-13T21:00:16Z",
//           "p": 0.18
//         },
//         {
//           "t": "2025-12-13T22:00:15Z",
//           "p": 0.18
//         },
//         {
//           "t": "2025-12-13T23:00:13Z",
//           "p": 0.175
//         },
//         {
//           "t": "2025-12-14T00:00:19Z",
//           "p": 0.175
//         }
//       ]
//     }
//

func main() {
	var err error

	var cli *client.Client
	{
		cli = client.New(client.Config{})
	}

	var opt token.Option
	{
		opt = token.Option{
			Tok: "51028631785140982248590109905860777819766751100592776255708383993117903047997",
			Sta: time.Unix(1765587600, 0),
			End: time.Unix(1765674000, 0),
			Fid: token.FidelityHour,
		}
	}

	var res token.Object
	{
		res, err = cli.Offchain().Token().Search(opt)
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
