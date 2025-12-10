package main

import (
	"encoding/json"
	"fmt"
	"time"

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
//     go run example/offchain/event/search/main.go | jq '.[].id'
//
//     "16483"
//     "16503"
//     "16502"
//     "16502"
//     "16516"
//     "16519"
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
			Tag: "100350", // soccer
			Clo: true,
			Ord: "startDate",
			Asc: true,
			Sta: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
			Lim: 3,
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

	// Fetch the next page of events again using the start date as cursor.

	{
		opt.Sta = res[len(res)-1].Start
	}

	{
		res, err = cli.Offchain().Event().Search(opt)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

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
