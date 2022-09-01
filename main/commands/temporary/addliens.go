package commands_temporary

import (
	"rwby-adventures/main/discord"
)

func Addliens(ctx *discord.CmdContext) {
	ctx.Player.Balance += 10000
	ctx.Player.Save()

	ctx.Reply(discord.ReplyParams{
		Content: "Added 10k.",
	})
}
