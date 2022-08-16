package market

import (
	"rwby-adventures/microservices"

	"github.com/yyewolf/gosf"
)

func UpdateListings(req *microservices.MarketCreate) (*gosf.Message, error) {
	body := gosf.StructToMap(req)
	message := &gosf.Message{
		Body: body,
	}

	if response, err := MarketMicroservice.Call("createListing", message); err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
