package commands_temporary

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/arenas"
	"rwby-adventures/main/discord"
	"rwby-adventures/microservices"

	"github.com/bwmarrin/discordgo"
	uuid "github.com/satori/go.uuid"
)

func createArena(ctx *discord.CmdContext) {
	if !arenas.ArenaMicroservice.Connected() {
		_, err := arenas.ArenaMicroservice.Connect()
		if err != nil {
			ctx.Reply(discord.ReplyParams{
				Content:   "Cannot contact arenas at the moment.",
				Ephemeral: true,
			})
			return
		}
	}

	ID := uuid.NewV4().String()
	in := &microservices.CreateArena{
		ID:        ID,
		ChannelID: ctx.ChannelID,
	}

	ctx.Reply(discord.ReplyParams{
		Content: fmt.Sprintf("Creating arena with ID %s", ID),
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Click here!",
						Style: discordgo.LinkButton,
						URL:   fmt.Sprintf("%sa/%s", config.ArenaHost, ID),
					},
				},
			},
		},
	})

	rep, err := arenas.CreateArena(in)
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content: fmt.Sprint(err),
		})
		return
	}

	if !rep.Success {
		ctx.Reply(discord.ReplyParams{
			Content: rep.Text,
		})
		return
	}
}
