package admin

import "rwby-adventures/main/discord"

func isAdmin(callback func(ctx *discord.CmdContext)) func(ctx *discord.CmdContext) {
	return func(ctx *discord.CmdContext) {
		if ctx.Player.DiscordID != "144472011924570113" {
			ctx.Reply(discord.ReplyParams{
				Content:   "You are not allowed to use this command.",
				Ephemeral: true,
			})
			return
		}
		callback(ctx)
	}
}
