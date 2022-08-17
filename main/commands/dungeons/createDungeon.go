package commands_dungeon

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/main/dungeons"
	"rwby-adventures/microservices"
	"time"

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
	if ctx.Player.CanDungeon() {
		t := ctx.Player.DungeonCooldown()
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("Sorry but you still have to wait **%.0fh %.0fm and %.0fs** before you can join in on a dungeon.", t.Hours(), t.Minutes(), t.Seconds()),
			Ephemeral: true,
		})
		return
	}

	if !dungeons.DungeonsMicroservice.Connected() {
		_, err := dungeons.DungeonsMicroservice.Connect()
		if err != nil {
			ctx.Reply(discord.ReplyParams{
				Content:   "Cannot contact dungeons at the moment.",
				Ephemeral: true,
			})
			return
		}
	}

	ctx.Player.Status.LastDungeon = time.Now().Unix()
	config.Database.Save(ctx.Player)

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
