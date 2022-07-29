package commands_misc

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var InviteCommand = &discord.Command{
	Name:        "invite",
	Description: "Invite me somewhere else.",
	Menu:        discord.ConfigurationMenu,
	Call:        invite,
}

func invite(ctx *discord.CmdContext) {
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       "Invite :",
			Description: "Help me visit more server!",
			Color:       config.Botcolor,
		},
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.Button{
						Label: "Invite",
						Style: discordgo.LinkButton,
						URL:   "https://discord.com/oauth2/authorize?client_id=602105650968920094&scope=bot%20applications.commands&permissions=322624",
					},
				},
			},
		},
	})
}
