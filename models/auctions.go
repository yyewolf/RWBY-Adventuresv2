package models

import (
	"rwby-adventures/config"
)

type AuctionBidders struct {
	AuctionID string `gorm:"primary_key;column:auction_id" json:"-"`
	UserID    string `gorm:"primary_key;column:user_id" json:"-"`
	Bid       int64  `gorm:"column:bid;not null" json:"amount"`
}

type Auction struct {
	ID             string `gorm:"primary_key;column:auction_id"`
	SellerID       string `gorm:"column:seller_id;not null" json:"-"`
	SellerName     string `gorm:"column:seller_name;not null" json:"seller_name"`
	StartedAt      int64  `gorm:"column:started_at;not null" json:"started_at"`
	EndsAt         int64  `gorm:"column:ends_at;not null" json:"ends_at"`
	TimeExtensions int64  `gorm:"column:time_extension;not null" json:"-"`
	Type           int    `gorm:"column:type;not null" json:"type"`

	// Foreign keys
	Char    *Character        `gorm:"foreignkey:UserID" json:"char"`
	Grimm   *Grimm            `gorm:"foreignkey:UserID" json:"grimm"`
	Bidders []*AuctionBidders `gorm:"foreignkey:AuctionID" json:"bidders"`
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
	e := config.Database.Find(a)
	config.Database.Joins("Stats").Find(&a.Char, "user_id = ?", a.ID)
	config.Database.Joins("Stats").Find(&a.Grimm, "user_id = ?", a.ID)
	config.Database.Order("bid desc").Find(&a.Bidders, "auction_id = ?", a.ID)
	err = e.Error
	return
}

func GetAuctions() (m []*Auction, err error) {
	e := config.Database.Order("auction_id desc").Find(&m)
	err = e.Error
	for _, a := range m {
		config.Database.Joins("Stats").Find(&a.Char, "user_id = ?", a.ID)
		config.Database.Joins("Stats").Find(&a.Grimm, "user_id = ?", a.ID)
		config.Database.Order("bid desc").Find(&a.Bidders, "auction_id = ?", a.ID)
	}
	return
}

func (a *Auction) Save() (err error) {
	return config.Database.Save(a).Error
}

func (a *Auction) Delete() (err error) {
	return config.Database.Delete(a).Error
}

func (a *Auction) Bid() int64 {
	if len(a.Bidders) == 0 {
		return 0
	}
	return a.Bidders[len(a.Bidders)-1].Bid
}

func (b *AuctionBidders) Save() (err error) {
	return config.Database.Save(b).Error
}

func (b *AuctionBidders) Delete() (err error) {
	return config.Database.Delete(b).Error
}
