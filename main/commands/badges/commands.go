package commands_badges

import (
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var BadgesCommand = &discord.Command{
	Name:        "badges",
	Description: "To view your badges.",
	SubCommands: []*discord.Command{
		{
			Name:        "view",
			Description: "To view your badges.",
			Call:        BadgesView,
			Menu:        discord.MiscMenu,
		},
		{
			Name:        "give",
			Description: "To give a badge to players.",
			Args: []discord.Arg{
				{
					Name:        "badge",
					Description: "The badge to claim.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "player",
					Description: "The player to give the badge to.",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
			},
			Call: BadgesGive,
			Menu: discord.MiscMenu,
		},
	},
}
