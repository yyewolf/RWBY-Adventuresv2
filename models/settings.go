package models

import "rwby-adventures/config"

type PlayerSettings struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	SubscribedToEvent bool `gorm:"column:event;not null"`
	DMable            bool `gorm:"column:dmable;not null"`
}

func (p *PlayerSettings) Save() {
	config.Database.Save(p)
}
