package websocket

import (
	"encoding/json"
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/market_back/cache"
	"rwby-adventures/market_back/microservice"
	"rwby-adventures/microservices"
	"rwby-adventures/models"

	"github.com/ambelovsky/gosf"
	"github.com/bwmarrin/discordgo"
)

type ListingBuyReq struct {
	ListingID string `json:"listing_id"`
}

func listingPurchase(client *gosf.Client, request *gosf.Request) *gosf.Message {
	token, exists := GetToken(request)
	if !exists {
		return gosf.NewFailureMessage("Invalid token")
	}
	t := token.(*Token)

	var req ListingBuyReq
	// unmarshal body :
	data, _ := json.Marshal(request.Message.Body)
	err := json.Unmarshal(data, &req)
	if err != nil {
		return gosf.NewFailureMessage(errorUnmarshalListing)
	}

	listing, err := models.GetListing(req.ListingID)
	if err != nil {
		return gosf.NewFailureMessage("Listing has not been found.")
	}

	p := models.GetPlayer(t.UserID)

	if p.TotalBalance() < listing.Price {
		return gosf.NewFailureMessage("You do not have enough Liens to purchase this.")
	}

	seller := models.GetPlayer(listing.SellerID)

	// Money transfer :
	if seller.DiscordID != t.UserID {
		seller.Balance += listing.Price
		p.Balance -= listing.Price
	}

	p.Save()
	seller.Save()

	// Listing transfer :
	if listing.Type == models.CharType {
		listing.Char.UserID = p.DiscordID
		listing.Char.Save()
	} else {
		listing.Grimm.UserID = p.DiscordID
		listing.Grimm.Save()
	}

	listing.Delete()

	var personaString string
	if listing.Type == models.CharType {
		personaString = listing.Char.FullString()
	} else {
		personaString = listing.Grimm.FullString()
	}

	// Buyer's message
	go microservice.SendMessageToBot(&microservices.MarketMessage{
		UserID: p.DiscordID,
		Message: &discordgo.MessageEmbed{
			Title:       "Listing Purchase",
			Color:       config.Botcolor,
			Description: fmt.Sprintf("You have successfully purchased `%s` for **%d** Liens.", personaString, listing.Price),
		},
	})

	// Seller's message
	go microservice.SendMessageToBot(&microservices.MarketMessage{
		UserID: seller.DiscordID,
		Message: &discordgo.MessageEmbed{
			Title:       "Listing Sold",
			Color:       config.Botcolor,
			Description: fmt.Sprintf("You have successfully sold `%s` for **%d** Liens.", personaString, listing.Price),
		},
	})

	go gosf.Broadcast("*", req.ListingID, gosf.NewSuccessMessage())
	go cache.RemoveListing(listing)
	return gosf.NewSuccessMessage("You successfully purchased this listing!")
}
