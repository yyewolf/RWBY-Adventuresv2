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

func SaveAuction(a *models.Auction) {
	for i, a2 := range Auctions {
		if a2.ID == a.ID {
			a, err := models.GetAuction(a.ID)
			if err != nil {
				return
			}
			Auctions[i] = a
			return
		}
	}
}

func RemoveListing(l *models.Listing) {
	for i, l2 := range Listings {
		if l2.ID == l.ID {
			Listings = append(Listings[:i], Listings[i+1:]...)
			return
		}
	}
}
