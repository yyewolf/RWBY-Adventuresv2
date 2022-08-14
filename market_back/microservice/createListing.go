package microservice

import (
	"rwby-adventures/market_back/cache"
	"rwby-adventures/microservices"
	"rwby-adventures/models"

	"github.com/ambelovsky/gosf"
)

func createListing(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// Create a new listing

	// Convert to our nicer format
	var req microservices.MarketCreate
	gosf.MapToStruct(request.Message.Body, &req)

	l, err := models.GetListing(req.ID)
	if err != nil {
		return gosf.NewFailureMessage(err.Error())
	}

	// prepend to list
	cache.Listings = append([]*models.Listing{l}, cache.Listings...)

	return gosf.NewSuccessMessage()
}

func createAuction(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// Create a new listing

	// Convert to our nicer format
	var req microservices.MarketCreate
	gosf.MapToStruct(request.Message.Body, &req)

	a, err := models.GetAuction(req.ID)
	if err != nil {
		return gosf.NewFailureMessage(err.Error())
	}

	// prepend to list
	cache.Auctions = append([]*models.Auction{a}, cache.Auctions...)

	return gosf.NewSuccessMessage()
}
