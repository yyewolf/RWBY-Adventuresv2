package cache

import (
	"rwby-adventures/models"
)

var Auctions []*models.Auction
var Listings []*models.Listing

func Init() {
	LoadAuctions()
	LoadListings()
}

func LoadAuctions() {
	var err error
	Auctions, err = models.GetAuctions()
	if err != nil {
		panic(err)
	}
}

func LoadListings() {
	var err error
	Listings, err = models.GetListings()
	if err != nil {
		panic(err)
	}
}
