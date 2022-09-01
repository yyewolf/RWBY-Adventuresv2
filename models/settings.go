package models

import "rwby-adventures/config"

type PlayerSettings struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	SubscribedToEvent bool `gorm:"column:event;not null"`
	DMable            bool `gorm:"column:dmable;not null"`

	MarketPublic  bool   `gorm:"column:market_public;not null;default:true"`
	ProfilePublic bool   `gorm:"column:profile_public;not null;default:true"`
	OrderBy       string `gorm:"column:order_by;not null;default:''"`
}

func (p *PlayerSettings) Save() {
	config.Database.Save(p)
}
