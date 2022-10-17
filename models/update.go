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
			Title: "OC Contest - RWBY Adventures V2",
			Description: "Hello there, adventurer!\n" +
				"I'm here to tell you about the OC Contest currently going on!\n" +
				"**What is an OC Contest?**\n" +
				"An OC Contest is a contest where you can create your own character and submit it to the contest.\n" +
				"**How do I submit my OC?**\n" +
				"First, you need to create your OC. You can do this by going to the [OC Creation Page](https://oc.rwbyadventures.com/) and filling out the form.\n" +
				"Once you've done that, you can vote for other OCs on the website.\n\n" +
				"Thank you for reading this :D.",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://media.discordapp.net/attachments/717814612892057610/1026930137884807278/Poster.png",
			},
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
