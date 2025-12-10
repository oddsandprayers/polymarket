package main

import (
	"encoding/json"
	"fmt"

	"github.com/oddsandprayers/polymarket/client"
	"github.com/oddsandprayers/polymarket/client/offchain/event"
	"github.com/xh3b4sd/tracer"
)

//
//     curl --request GET \
//          --url 'https://gamma-api.polymarket.com/events?tag_id=100350&limit=2&closed=true' \
//          | jq .
//

//
//     go run example/offchain/event/search/main.go
//

func main() {
	var err error

	var cli *client.Client
	{
		cli = client.New(client.Config{})
	}

	var opt event.Option
	{
		opt = event.Option{
			Tag: "100350",
			Clo: true,
			Lim: 2,
		}
	}

	var res []event.Object
	{
		res, err = cli.Offchain().Event().Search(opt)
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
