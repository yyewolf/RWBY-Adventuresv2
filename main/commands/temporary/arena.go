package commands_temporary

import (
	"fmt"
	"rwby-adventures/main/arenas"
	"rwby-adventures/main/discord"
	"rwby-adventures/microservices"

	uuid "github.com/satori/go.uuid"
)

var ArenaCommand = &discord.Command{
	Name:        "arena",
	Description: "Create arenas.",
	Menu:        discord.GeneralMenu,
	Call:        createArena,
}

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
