package websocket

import (
	"rwby-adventures/market_back/cache"
	"rwby-adventures/models"

	"github.com/ambelovsky/gosf"
)

func getLatestAuctions(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// Get 10 auctions from cache :
	var auctions []*models.Auction
	for i, a := range cache.Auctions {
		if i >= 10 {
			break
		}
		auctions = append(auctions, a)
	}

	var icons []string
	for _, a := range auctions {
		if a.Char != nil {
			icons = append(icons, a.Char.ToRealChar().IconURL)
		} else if a.Grimm != nil {
			icons = append(icons, a.Grimm.ToRealGrimm().IconURL)
		}
	}

	msg := gosf.NewSuccessMessage()
	msg.Body = make(map[string]interface{})
	msg.Body["auctions"] = auctions
	msg.Body["icons"] = icons
	return msg
}
