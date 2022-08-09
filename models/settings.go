package models

type PlayerSettings struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	SubscribedToEvent bool `gorm:"column:event;not null"`
}
