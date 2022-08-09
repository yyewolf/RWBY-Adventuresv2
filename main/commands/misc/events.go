package commands_misc

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var Eventcommand = &discord.Command{
	Name:        "event",
	Description: "Subscribe to notifications for future events.",
	Menu:        discord.ConfigurationMenu,
	Call:        event,
}

func event(ctx *discord.CmdContext) {
	ctx.Player.Settings.SubscribedToEvent = !ctx.Player.Settings.SubscribedToEvent

	if ctx.Player.Settings.SubscribedToEvent {
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageEmbed{
				Title:       "Event :",
				Description: "You are now **subscribed** to future events.",
			},
		})
	} else {
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageEmbed{
				Title:       "Event :",
				Description: "You are **no longer subscribed** to future events.",
			},
		})
	}

	config.Database.Save(ctx.Player)
}
