package commands_temporary

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
)

func ResetGamble(ctx *discord.CmdContext) {
	ctx.Player.Gamble.Time = 0
	ctx.Player.Gamble.Amount = 0
	config.Database.Save(ctx.Player.Gamble)
	ctx.Reply(discord.ReplyParams{
		Content:   "Your gamble should be available now.",
		Ephemeral: true,
	})
}
