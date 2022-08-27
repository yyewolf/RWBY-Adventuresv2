package websocket

import (
	"encoding/json"
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/market_back/cache"
	"rwby-adventures/market_back/microservice"
	"rwby-adventures/microservices"
	"rwby-adventures/models"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/yyewolf/gosf"
)

type AuctionBidReq struct {
	AuctionID string `json:"auction_id"`
	Bid       string `json:"amount"`
}

func auctionBid(client *gosf.Client, request *gosf.Request) *gosf.Message {
	token, exists := GetToken(request)
	if !exists {
		return gosf.NewFailureMessage("Invalid token")
	}
	t := token.(*Token)

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
	p := models.GetPlayer(t.UserID)
	if p.TotalBalance() < bid {
		return gosf.NewFailureMessage("You do not have enough Liens.")
	}

	b := &models.AuctionBidders{
		AuctionID: req.AuctionID,
		UserID:    t.UserID,
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

	var personaString string
	if auction.Type == models.CharType {
		personaString = auction.Char.FullString()
	} else {
		personaString = auction.Grimm.FullString()
	}

	// Bidders's message
	go microservice.SendMessageToBot(&microservices.MarketMessage{
		UserID: p.DiscordID,
		Message: &discordgo.MessageEmbed{
			Title:       "Auction Bid",
			Color:       config.Botcolor,
			Description: fmt.Sprintf("You have successfully bidded %d Liens on `%s`.", b.Bid, personaString),
		},
	})

	go gosf.Broadcast("*", req.AuctionID, msg)
	go cache.SaveAuction(auction)
	return gosf.NewSuccessMessage("You placed your bid!")
}
