package commands_shop

import (
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var ShopCommand = &discord.Command{
	Name:        "shop",
	Description: "View the shop menu.",
	SubCommands: []*discord.Command{
		{
			Name:        "view",
			Description: "View the shop menu.",
			Call:        Shop,
		},
		{
			Name:        "buy",
			Description: "Buy an item from the shop.",
			Call:        Buy,
			Args: []discord.Arg{
				{
					Name:        "item",
					Description: "The item to buy.",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Size:        1,
					Choices: []*discord.Choice{
						{
							Name:  "XP Boost",
							Value: 1,
						},
						{
							Name:  "Luck Boost",
							Value: 2,
						},
						{
							Name:  "Loot box",
							Value: 3,
						},
						{
							Name:  "Grimm box",
							Value: 4,
						},
						{
							Name:  "Rare loot box",
							Value: 5,
						},
						{
							Name:  "Rare grimm box",
							Value: 6,
						},
						{
							Name:  "Exchange for rare loot box",
							Value: 7,
						},
						{
							Name:  "Exchange for rare grimm box",
							Value: 8,
						},
						{
							Name:  "Backpack",
							Value: 9,
						},
					},
					Required: true,
				},
			},
		},
	},
}
