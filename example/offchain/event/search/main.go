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
//     go run example/offchain/event/search/main.go | jq '.[].id'
//
//     "12137"
//     "12138"
//     "12139"
//     "12140"
//     "12141"
//     "12142"
//     <repeat>
//     "12137"
//     "12138"
//     "12139"
//     "12140"
//     "12141"
//     "12142"
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
			Ord: "id",
			Asc: true,
			Lim: 3,
			Off: 0,
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

	// Fetch the next page of events again using the sorted offset as cursor.

	{
		opt.Off = 3
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

	// Fetch the the above results in one go to verify that we get the same result.

	{
		opt.Lim = 6
		opt.Off = 0
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
