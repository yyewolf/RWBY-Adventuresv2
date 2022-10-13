package models

import (
	"rwby-adventures/config"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Update struct {
	ID   int       `gorm:"primary_key;column:id"`
	Time time.Time `gorm:"column:time;not null"`
}

var message = &discordgo.MessageSend{
	Embeds: []*discordgo.MessageEmbed{
		{
			Title: "Update - RWBY Adventures V2",
			Description: "Hello there, adventurer!\n" +
				"I'm here to tell you about the new update of RWBY Adventures V2 because I haven't had the time to tell you about it!\n" +
				"You probably noticed most of it but here are the changes :\n" +
				" - **Slash Commands** (ty Discord)\n" +
				" - **Complete rewrite**\n" +
				" - **New market looks**\n" +
				" - **New arenas looks**\n" +
				" - **New trade looks**\n" +
				" - **New dungeons**\n" +
				" - **New [OC Contest](https://oc.rwbyadventures.com/)**\n\n" +
				"And here are some upcoming features :\n" +
				" - **New arenas mechanics**\n" +
				" - **New dungeons loots**\n" +
				" - **Trade update**\n" +
				" - **New characters/grimms**\n\n" +
				"If you have any **suggestions**, please join the [support server](https://discord.gg/3Z5Y4Z5) and tell us!\n" +
				"I hope you like the news changes.",
			Color: config.Botcolor,
		},
	},
}

// EDIT ON UPDATE
func (u *Update) GetMessage() *discordgo.MessageSend {
	return message
}

func (u *Update) Save() {
	config.Database.Save(u)
}

func GetLastUpdate() *Update {
	u := &Update{}
	config.Database.Order("id desc").Limit(1).Find(u)
	return u
}
