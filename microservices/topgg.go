package microservices

import "github.com/bwmarrin/discordgo"

type TopGGMessage struct {
	UserID  string
	Message *discordgo.MessageEmbed
}
