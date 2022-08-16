package models

import (
	"rwby-adventures/config"
)

type Listing struct {
	ID         string `gorm:"primary_key;column:listing_id"`
	SellerID   string `gorm:"column:seller_id;not null" json:"-"`
	SellerName string `gorm:"column:seller_name" json:"seller_name"`
	Price      int64  `gorm:"column:price;not null" json:"price"`
	Note       string `gorm:"column:note" json:"note"`
	Type       int    `gorm:"column:type;not null" json:"type"`

	// Foreign keys
	Char  *Character `gorm:"foreignkey:UserID" json:"char"`
	Grimm *Grimm     `gorm:"foreignkey:UserID" json:"grimm"`
}

type PlayerMarket struct {
	DiscordID string

	Listings []*Listing `gorm:"foreignkey:SellerID"`
	Auctions []*Auction `gorm:"foreignkey:SellerID"`
}

func CreateListing(l *Listing) (err error) {
	d := config.Database.Create(l)
	if l.Char != nil {
		l.Char.UserID = l.ID
		config.Database.Save(l.Char)
	} else if l.Grimm != nil {
		l.Grimm.UserID = l.ID
		config.Database.Save(l.Grimm)
	}
	return d.Error
}

func GetListing(id string) (m *Listing, err error) {
	m = &Listing{
		ID: id,
	}
	e := config.Database.
		Preload("Grimm.Stats").
		Preload("Char.Stats").
		Joins("Char").
		Joins("Grimm").
		Find(m)
	err = e.Error
	return
}

func GetListings() (m []*Listing, err error) {
	e := config.Database.
		Order("listing_id desc").
		Preload("Grimm.Stats").
		Preload("Char.Stats").
		Joins("Char").
		Joins("Grimm").
		Find(&m)
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

func (l *Listing) Save() (err error) {
	return config.Database.Save(l).Error
}

func (l *Listing) Delete() (err error) {
	return config.Database.Delete(l).Error
}
