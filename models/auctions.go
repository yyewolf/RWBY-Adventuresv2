package models

import (
	"rwby-adventures/config"
)

type AuctionBidders struct {
	AuctionID string `gorm:"primary_key;column:auction_id"`
	UserID    string `gorm:"primary_key;column:user_id"`
	Bid       int64  `gorm:"column:bid;not null"`
}

type Auction struct {
	ID             string `gorm:"primary_key;column:auction_id"`
	SellerID       string `gorm:"column:seller_id;not null"`
	SellerName     string `gorm:"column:seller_name;not null"`
	StartedAt      int64  `gorm:"column:started_at;not null"`
	EndsAt         int64  `gorm:"column:ends_at;not null"`
	TimeExtensions int64  `gorm:"column:time_extension;not null"`
	Type           int    `gorm:"column:type;not null"`

	// Foreign keys
	Char    *Character        `gorm:"foreignkey:UserID"`
	Grimm   *Grimm            `gorm:"foreignkey:UserID"`
	Bidders []*AuctionBidders `gorm:"foreignkey:AuctionID"`
}

func CreateAuction(a *Auction) (err error) {
	d := config.Database.Create(a)
	if a.Char != nil {
		a.Char.UserID = a.ID
		config.Database.Save(a.Char)
	} else if a.Grimm != nil {
		a.Grimm.UserID = a.ID
		config.Database.Save(a.Grimm)
	}
	return d.Error
}

func GetAuction(id string) (a *Auction, err error) {
	a = &Auction{
		ID: id,
	}
	e := config.Database.Find(a, id)
	config.Database.Joins("Stats").Find(&a.Char, "user_id = ?", a.ID)
	config.Database.Joins("Stats").Find(&a.Grimm, "user_id = ?", a.ID)
	config.Database.Find(&a.Bidders, "auction_id = ?", a.ID)
	err = e.Error
	return
}

func GetAuctions() (m []*Auction, err error) {
	e := config.Database.Order("auction_id desc").Find(&m)
	err = e.Error
	for _, a := range m {
		config.Database.Joins("Stats").Find(&a.Char, "user_id = ?", a.ID)
		config.Database.Joins("Stats").Find(&a.Grimm, "user_id = ?", a.ID)
		config.Database.Find(&a.Bidders, "auction_id = ?", a.ID)
	}
	return
}
