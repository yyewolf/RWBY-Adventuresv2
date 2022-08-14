package websocket

import (
	"encoding/json"
	"rwby-adventures/market_back/cache"
	"rwby-adventures/models"

	"github.com/ambelovsky/gosf"
)

type ListingBuyReq struct {
	ListingID string `json:"listing_id"`
}

func listingPurchase(client *gosf.Client, request *gosf.Request) *gosf.Message {
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

	p := models.GetPlayer("144472011924570113")

	if p.TotalBalance() < listing.Price {
		return gosf.NewFailureMessage("You do not have enough Liens to purchase this.")
	}

	seller := models.GetPlayer(listing.SellerID)
	// Money transfer :
	seller.Balance += listing.Price
	p.Balance -= listing.Price

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

	go gosf.Broadcast("*", req.ListingID, gosf.NewSuccessMessage())
	go cache.RemoveListing(listing)
	return gosf.NewSuccessMessage("You successfully purchased this listing!")
}
