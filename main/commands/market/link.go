package commands_market

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

func MarketLink(ctx *discord.CmdContext) {
	ctx.Reply(discord.ReplyParams{
		Content:   "Click the button to go to the market.",
		Ephemeral: true,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Market",
						Style: discordgo.SuccessButton,
						URL:   config.MarketFront,
					},
				},
			},
		},
	})
}
