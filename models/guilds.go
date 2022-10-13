package models

import (
	"rwby-adventures/config"
	"time"
)

type Guild struct {
	GuildID                    string `gorm:"primary_key;column:guild_id"`
	LastArena                  int64  `gorm:"column:last_arena;not null"`
	AutomatedMessagesChannelID string `gorm:"column:channel_id;not null"`
	AutomatedMessagesEnabled   bool   `gorm:"column:channel_on;not null"`
	PingRoles                  string `gorm:"column:pingable;not null"`
	Prefix                     string `gorm:"column:prefix;not null"`

	LastUpdateMessage time.Time `gorm:"column:last_update_message"`
}

func GetGuild(id string) *Guild {
	g := &Guild{
		GuildID: id,
		Prefix:  "r!",
	}
	e := config.Database.
		Find(g, id)
	if e.Error != nil {
		config.Database.Create(g)
	}
	return g
}

func (g *Guild) Save() {
	config.Database.Save(g)
}
