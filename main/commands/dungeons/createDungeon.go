package commands_dungeon

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/main/dungeons"
	"rwby-adventures/microservices"

	"github.com/ambelovsky/gosf"
	"github.com/bwmarrin/discordgo"
	uuid "github.com/satori/go.uuid"
)

var DungeonCommand = &discord.Command{
	Name:        "dungeon",
	Description: "To join in on a dungeon.",
	Menu:        discord.InventoryMenu,
	Call:        createDungeon,
}

func createDungeon(ctx *discord.CmdContext) {
	if !dungeons.DungeonsMicroservice.Connected() {
		ctx.Reply(discord.ReplyParams{
			Content:   "Cannot contact dungeons at the moment.",
			Ephemeral: true,
		})
		return
	}

	ID := uuid.NewV4().String()
	req := &microservices.DungeonCreateRequest{
		ID: ID,
	}

	ctx.Reply(discord.ReplyParams{
		Content: "Here you go",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Join!",
						Style: discordgo.LinkButton,
						URL:   fmt.Sprintf("http://%s%s/d/%s", config.DungeonHost, config.DungeonPort, ID),
					},
				},
			},
		},
	})

	response, err := dungeons.CreateDungeon(req)
	if err != nil {
		return
	}

	resp := &microservices.DungeonEndResponse{}
	gosf.MapToStruct(response.Body, resp)

	// HERE THE DUNGEON FINISHED

	ctx.Reply(discord.ReplyParams{
		Content: "It finished !",
	})
}
