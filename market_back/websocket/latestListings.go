package websocket

import (
	"rwby-adventures/market_back/cache"
	"rwby-adventures/models"

	"github.com/ambelovsky/gosf"
)

func getLatestListings(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// Get 10 listings from cache :
	var listings []*models.Listing
	for i, l := range cache.Listings {
		if i >= 10 {
			break
		}
		listings = append(listings, l)
	}

	var icons []string
	for _, l := range listings {
		if l.Char != nil {
			icons = append(icons, l.Char.ToRealChar().IconURL)
		} else if l.Grimm != nil {
			icons = append(icons, l.Grimm.ToRealGrimm().IconURL)
		}
	}

	msg := gosf.NewSuccessMessage()
	msg.Body = make(map[string]interface{})
	msg.Body["listings"] = listings
	msg.Body["icons"] = icons
	return msg
}
