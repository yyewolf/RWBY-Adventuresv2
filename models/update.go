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
			Title:       "RWBY Adventures - End of the line",
			Description: "Hello there, adventurer!\n\nI regret to inform you that the service for RWBY Adventures Discord Bot has come to an end. It has been an amazing journey and I'm grateful for all the support and feedback I've received from the community.\n\nI apologize for any inconvenience this may have caused. If you have any questions or concerns, please feel free to reach out to me.\n\nAlso, for those who are interested, I'm willing to share the source code of the bot. If you'd like to get a copy of it, please send me a direct message.\n\nThank you for your understanding and for being part of this adventure.\nIt was truly a great experience, and I hope to see you around!\n\n- Yewolf\n\nP.S. I will be active on the support server if you have any questions: https://discord.gg/wHhXefjDK4.",
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
