package commands_inventory

import (
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var FavoritesCommand = &discord.Command{
	Name:        "favorites",
	Description: "All commands regarding favorites",
	SubCommands: []*discord.Command{
		{
			Name:        "add",
			Description: "Adds a persona to your favorites.",
			Menu:        discord.InventoryMenu,
			Call:        AddFavoritePersona,
			Args: []discord.Arg{
				{
					Name:        "id",
					Description: "Identification number of your persona.",
					Size:        1,
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "latest",
					Description: "Whether or not you want to view the infos of your latest persona.",
					Size:        1,
					Required:    false,
					Type:        discordgo.ApplicationCommandOptionBoolean,
				},
			},
		},
		{
			Name:        "remove",
			Description: "Removes a persona from your favorites.",
			Menu:        discord.InventoryMenu,
			Call:        RemoveFavoritePersona,
			Args: []discord.Arg{
				{
					Name:        "id",
					Description: "Identification number of your persona.",
					Size:        1,
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "latest",
					Description: "Whether or not you want to view the infos of your latest persona.",
					Size:        1,
					Required:    false,
					Type:        discordgo.ApplicationCommandOptionBoolean,
				},
			},
		},
		{

			Name:        "chars",
			Description: "View informations about your favorites characters.",
			Menu:        discord.InventoryMenu,
			Call:        FavoriteChars,
			Args: []discord.Arg{
				{
					Name:        "page",
					Description: "The page of your inventory you want to check.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "name",
					Description: "Filter your inventory by name.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "level",
					Description: "Filter your inventory by level.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "rarity",
					Description: "Filter your inventory by rarity.",
					Size:        1,
					Choices: []*discord.Choice{
						{
							Name:  "Common",
							Value: 0,
						},
						{
							Name:  "Uncommon",
							Value: 1,
						},
						{
							Name:  "Rare",
							Value: 2,
						},
						{
							Name:  "Very_Rare",
							Value: 3,
						},
						{
							Name:  "Legendary",
							Value: 4,
						},
						{
							Name:  "Collector",
							Value: 5,
						},
					},
					Type: discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "arms",
					Description: "Filter your inventory by amount of arm.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "value_above",
					Description: "Filter your inventory by value.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "value_below",
					Description: "Filter your inventory by value.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{

			Name:        "grimms",
			Description: "View informations about your favorites grimms.",
			Menu:        discord.InventoryMenu,
			Call:        FavoriteGrimms,
			Args: []discord.Arg{
				{
					Name:        "page",
					Description: "The page of your inventory you want to check.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "name",
					Description: "Filter your inventory by name.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "level",
					Description: "Filter your inventory by level.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "rarity",
					Description: "Filter your inventory by rarity.",
					Size:        1,
					Choices: []*discord.Choice{
						{
							Name:  "Common",
							Value: 0,
						},
						{
							Name:  "Uncommon",
							Value: 1,
						},
						{
							Name:  "Rare",
							Value: 2,
						},
						{
							Name:  "Very_Rare",
							Value: 3,
						},
						{
							Name:  "Legendary",
							Value: 4,
						},
						{
							Name:  "Collector",
							Value: 5,
						},
					},
					Type: discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "arms",
					Description: "Filter your inventory by amount of arm.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "value_above",
					Description: "Filter your inventory by value.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionString,
				},
				{
					Name:        "value_below",
					Description: "Filter your inventory by value.",
					Size:        1,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
	},
}

func FavoriteChars(ctx *discord.CmdContext) {
	ctx.Arguments = append(ctx.Arguments, &discord.CommandArg{
		Name:  "favorites",
		Value: true,
		Found: true,
	})

	Inventory(ctx)
}

func FavoriteGrimms(ctx *discord.CmdContext) {
	ctx.Arguments = append(ctx.Arguments, &discord.CommandArg{
		Name:  "favorites",
		Value: true,
		Found: true,
	})

	Pack(ctx)
}
