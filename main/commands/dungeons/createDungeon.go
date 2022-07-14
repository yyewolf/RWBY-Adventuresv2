package commands_dungeon

import (
	"context"
	"fmt"
	"rwby-adventures/config"
	dungeonpc "rwby-adventures/dungeons_rpc"
	"rwby-adventures/main/discord"
	rwby_grpc "rwby-adventures/main/grpc"

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
	_, err := rwby_grpc.DungeonServer.Ping(context.Background(), &dungeonpc.PingReq{})
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content:   "There has been an error : " + err.Error(),
			Ephemeral: true,
		})
		return
	}

	ID := uuid.NewV4().String()
	in := &dungeonpc.CreateDungeonReq{
		Id: ID,
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

	rep, err := rwby_grpc.DungeonServer.CreateDungeon(context.Background(), in)
	if err != nil {
		return
	}
	if rep.Status == 1 {
		return
	}

	// HERE THE DUNGEON FINISHED

	ctx.Reply(discord.ReplyParams{
		Content: "It finished !",
	})
}
