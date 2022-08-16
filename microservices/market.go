package microservices

import "github.com/bwmarrin/discordgo"

type MarketCreate struct {
	ID string
}

type MarketMessage struct {
	UserID  string
	Message *discordgo.MessageEmbed
}
