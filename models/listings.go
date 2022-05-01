package models

import (
	"rwby-adventures/config"
)

type Listing struct {
	ID         string `gorm:"primary_key;column:listing_id"`
	SellerID   string `gorm:"column:seller_id;not null"`
	SellerName string `gorm:"column:seller_name"`
	Price      int    `gorm:"column:price;not null"`
	Note       string `gorm:"column:note"`
	Type       int    `gorm:"column:type;not null"`

	// Foreign keys
	Char  Character `gorm:"foreignkey:UserID"`
	Grimm Grimm     `gorm:"foreignkey:UserID"`
}

type PlayerMarket struct {
	DiscordID string

	Listings []*Listing `gorm:"foreignkey:SellerID"`
	Auctions []*Auction `gorm:"foreignkey:SellerID"`
}

func GetListing(id string) (m *Listing, err error) {
	m = &Listing{
		ID: id,
	}
	e := config.Database.Find(m, id)
	config.Database.Joins("Stats").Find(&m.Char, "user_id = ?", m.ID)
	config.Database.Joins("Stats").Find(&m.Grimm, "user_id = ?", m.ID)
	err = e.Error
	return
}

func (p *Player) FillPlayerMarket() {
	m := &PlayerMarket{
		DiscordID: p.DiscordID,
	}
	config.Database.Find(&m.Listings, "seller_id = ?", m.DiscordID)
	for i := range m.Listings {
		config.Database.Joins("Stats").Find(&m.Listings[i].Char, "user_id = ?", m.Listings[i].ID)
		config.Database.Joins("Stats").Find(&m.Listings[i].Grimm, "user_id = ?", m.Listings[i].ID)
	}
	config.Database.Find(&m.Auctions, "seller_id = ?", m.DiscordID)
	for i := range m.Auctions {
		config.Database.Joins("Stats").Find(&m.Auctions[i].Char, "user_id = ?", m.Auctions[i].ID)
		config.Database.Joins("Stats").Find(&m.Auctions[i].Char, "user_id = ?", m.Auctions[i].ID)
		config.Database.Find(&m.Auctions[i].Bidders, "auction_id = ?", m.Auctions[i].ID)
	}
	p.Market = m
}
