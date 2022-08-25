package commands_duels

import (
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var DuelCommand = &discord.Command{
	Name:        "duel",
	Description: "To duel your friends.",
	Args: []discord.Arg{
		{
			Name:        "opponent",
			Description: "The opponent.",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    true,
		},
	},
	Menu: discord.GamesMenu,
	Call: DuelCreate,
}
