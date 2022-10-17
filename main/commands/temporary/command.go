package commands_temporary

import (
	"rwby-adventures/main/discord"
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
	},
}
