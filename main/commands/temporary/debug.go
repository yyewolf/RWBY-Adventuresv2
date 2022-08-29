package commands_temporary

import (
	"bytes"
	"encoding/json"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

func Debug(ctx *discord.CmdContext) {
	data, _ := json.Marshal(ctx.Player)

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageSend{
			Files: []*discordgo.File{
				{
					Name:        "player.json",
					ContentType: "application/json",
					Reader:      bytes.NewReader(data),
				},
			},
			Embed: &discordgo.MessageEmbed{
				Title: "Player Debug",
			},
		},
	})
}
