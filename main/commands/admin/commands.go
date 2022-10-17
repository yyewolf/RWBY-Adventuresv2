package admin

import (
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var AdminCommands = &discord.Command{
	Name:        "admin",
	Description: "Admin commands.",
	Restricted:  true,
	SubCommands: []*discord.Command{
		{
			Name:        "update",
			Description: "Triggers the update message, don't forget to modify the embed in the code.",
			Restricted:  true,
			Call:        isAdmin(Update),
		},
		{
			Name:        "addbox",
			Description: "Gives n boxes.",
			Restricted:  true,
			Call:        isAdmin(AddBox),
			Args: []discord.Arg{
				{
					Name:        "type",
					Description: "The type of box to add.",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Choices: []*discord.Choice{
						{
							Value: 0,
							Name:  "Character Boxes",
						},
						{
							Value: 1,
							Name:  "Rare Character Boxes",
						},
						{
							Value: 2,
							Name:  "Grimm Boxes",
						},
						{
							Value: 3,
							Name:  "Rare Grimm Boxes",
						},
					},
					Required: true,
				},
				{
					Name:        "amount",
					Description: "The amount of boxes to add.",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Required:    true,
				},
				{
					Name:        "user",
					Description: "The user to give the boxes to.",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
			},
		},
	},
}
