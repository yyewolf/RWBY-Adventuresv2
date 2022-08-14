package websocket

import (
	"encoding/json"
	"rwby-adventures/market_back/cache"
	"rwby-adventures/models"
	"strconv"

	"github.com/ambelovsky/gosf"
)

type AuctionBidReq struct {
	AuctionID string `json:"auction_id"`
	Bid       string `json:"amount"`
}

func auctionBid(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req AuctionBidReq
	// unmarshal body :
	data, _ := json.Marshal(request.Message.Body)
	err := json.Unmarshal(data, &req)
	if err != nil {
		return gosf.NewFailureMessage(errorUnmarshalAuction)
	}

	bid, err := strconv.ParseInt(req.Bid, 10, 64)
	if err != nil {
		return gosf.NewFailureMessage(errorParseAuction)
	}

	auction, err := models.GetAuction(req.AuctionID)
	if err != nil {
		return gosf.NewFailureMessage("Auction has not been found.")
	}
	if auction.Bid()+100 >= bid {
		return gosf.NewFailureMessage("Bid must be at least 100 Liens greater than current bid.")
	}
	p := models.GetPlayer("144472011924570113")
	if p.TotalBalance() < bid {
		return gosf.NewFailureMessage("You do not have enough Liens.")
	}

	b := &models.AuctionBidders{
		AuctionID: req.AuctionID,
		UserID:    "144472011924570113",
		Bid:       bid,
	}
	b.Save()

	p.BiddedBalance += bid
	p.Save()

	if len(auction.Bidders) > 0 {
		b := auction.Bidders[len(auction.Bidders)-1]
		p := models.GetPlayer(b.UserID)
		p.BiddedBalance -= b.Bid
		p.Save()
	}

	msg := gosf.NewSuccessMessage()
	msg.Body = make(map[string]interface{})
	msg.Body["amount"] = bid
	msg.Body["ends_at"] = auction.EndsAt

	go gosf.Broadcast("*", req.AuctionID, msg)
	go cache.SaveAuction(auction)
	return gosf.NewSuccessMessage("You placed your bid!")
}
