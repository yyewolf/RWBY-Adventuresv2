package commands_temporary

import (
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var TemporaryCommand = &discord.Command{
	Name:        "temp",
	Description: "All beta related commands.",
	SubCommands: []*discord.Command{
		{
			Name:        "arena",
			Description: "Create arenas.",
			Menu:        discord.GeneralMenu,
			Call:        createArena,
		},
		{
			Name:        "addliens",
			Description: "Add liens.",
			Menu:        discord.GeneralMenu,
			Call:        Addliens,
		},
		{
			Name:        "debug",
			Description: "Debugging commands.",
			Menu:        discord.GeneralMenu,
			Call:        Debug,
		},
		{
			Name:        "addchar",
			Description: "Adds a character to your inventory.",
			Aliases:     discord.CmdAlias{"ac"},
			Menu:        discord.GeneralMenu,
			Call:        Addchar,
		},
		{
			Name:        "delete",
			Description: "Delete all your data.",
			Aliases:     discord.CmdAlias{"del"},
			Menu:        discord.GeneralMenu,
			Call:        Delete,
		},
		{
			Name:        "resetdungeon",
			Description: "Resets your dungeon.",
			Menu:        discord.GeneralMenu,
			Call:        ResetDungeon,
		},
		{
			Name:        "addbox",
			Description: "Adds boxes to your inventory.",
			Menu:        discord.GeneralMenu,
			Call:        AddBox,
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
			},
		},
	},
}
