package commands_gambles

import "rwby-adventures/main/discord"

var GambleCommand = &discord.Command{
	Name:        "gamble",
	Description: "To do your daily gambling.",
	Menu:        discord.GamesMenu,
	Call:        gamble,
}
