package discord

import (
	"rwby-adventures/config"
	"rwby-adventures/microservices"
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/bwmarrin/discordgo"
	"github.com/pmylund/go-cache"
	"github.com/yyewolf/gosf"
)

var sent = cache.New(1*time.Minute, 10*time.Second)

func endArena(channel *gosocketio.Channel, request *gosf.Message) {
	_, found := sent.Get(request.Text)
	sent.Set(request.Text, true, 0)

	var req microservices.EndArena
	gosf.MapToStruct(request.Body, &req)

	if !found {
		Session.ChannelMessageSendEmbed(req.ChannelID, &discordgo.MessageEmbed{
			Title:       "Arena rewards",
			Color:       config.Botcolor,
			Description: req.Message,
			Footer:      DefaultFooter,
		})
	}
}
