package commands_trade

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

var TradesCommand = &discord.Command{
	Name:        "trade",
	Description: "All commands regarding trades.",
	SubCommands: []*discord.Command{
		{
			Name:        "accept",
			Description: "Accept a trade.",
			Menu:        discord.InventoryMenu,
			Call:        AcceptTrade,
			Args: []discord.Arg{
				{
					Name:        "id",
					Description: "Identification number of your persona.",
					Size:        1,
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "refuse",
			Description: "Refuse a trade.",
			Menu:        discord.InventoryMenu,
			Call:        RefuseTrade,
			Args: []discord.Arg{
				{
					Name:        "id",
					Description: "Identification number of your persona.",
					Size:        1,
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "delete",
			Description: "Delete a trade.",
			Menu:        discord.InventoryMenu,
			Call:        DeleteTrade,
			Args: []discord.Arg{
				{
					Name:        "id",
					Description: "Identification number of your persona.",
					Size:        1,
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "start",
			Description: "Start a trade.",
			Menu:        discord.InventoryMenu,
			Call:        StartTrade,
			Args: []discord.Arg{
				{
					Name:        "user",
					Description: "Person you'll be trading with.",
					Size:        1,
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionUser,
				},
			},
		},
	},
}

func StartTrade(ctx *discord.CmdContext) {
	arg := ctx.Arguments.GetArg("user", 0, "")
	if !arg.Found {
		ctx.Reply(discord.ReplyParams{
			Content: "You need to specify a user.",
		})
		return
	}
	TargetID := fmt.Sprint(arg.Value)
	target := models.GetPlayer(TargetID)
	if target.IsNew {
		ctx.Reply(discord.ReplyParams{
			Content: "You cannot trade with that person.",
		})
		return
	}
	ctx.Reply(discord.ReplyParams{
		Content:   "Click this to continue !",
		Ephemeral: true,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Here.",
						Style: discordgo.LinkButton,
						URL:   fmt.Sprintf("http://%s%s/t/%s", config.TradeHost, config.TradePort, TargetID),
					},
				},
			},
		},
	})
}
