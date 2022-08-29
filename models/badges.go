package models

import "rwby-adventures/config"

type Badges struct {
	BadgeID     int                `gorm:"primary_key;column:badge_id"`
	Name        string             `gorm:"column:name;not null"`
	Emoji       string             `gorm:"column:emoji;not null"`
	Description string             `gorm:"column:description;not null"`
	Check       func(*Player) bool `gorm:"-" json:"-"`
}

func (b *Badges) Save() {
	config.Database.Save(b)
}

type PlayerBadges struct {
	BadgeID   int    `gorm:"primary_key;column:badge_id"`
	DiscordID string `gorm:"primary_key;column:discord_id"`

	Badge *Badges `gorm:"foreignkey:BadgeID"`
}

func (b *PlayerBadges) Save() {
	config.Database.Save(b)
}
