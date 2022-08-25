package commands_dungeon

import "rwby-adventures/main/discord"

var DungeonCommand = &discord.Command{
	Name:        "dungeon",
	Description: "To join in on a dungeon.",
	Menu:        discord.GamesMenu,
	Call:        createDungeon,
}
