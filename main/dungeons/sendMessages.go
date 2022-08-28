package dungeons

import (
	"rwby-adventures/main/discord"
	"rwby-adventures/microservices"
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/bwmarrin/discordgo"
	"github.com/pmylund/go-cache"
	"github.com/yyewolf/gosf"
)

var sent = cache.New(1*time.Minute, 10*time.Second)

func sendMessage(channel *gosocketio.Channel, request *gosf.Message) {
	_, found := sent.Get(request.Text)
	sent.Set(request.Text, true, 0)

	var req microservices.DungeonsMessage
	gosf.MapToStruct(request.Body, &req)

	var ch *discordgo.Channel
	var err error
	// Create User DM
	if ch, err = discord.Session.UserChannelCreate(req.UserID); err != nil {
		return
	}

	req.Message.Footer = discord.DefaultFooter

	if !found {
		discord.Session.ChannelMessageSendEmbed(ch.ID, req.Message)
	}
}
