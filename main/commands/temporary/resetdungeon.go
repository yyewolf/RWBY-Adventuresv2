package commands_temporary

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
)

func ResetDungeon(ctx *discord.CmdContext) {
	ctx.Player.Status.LastDungeon = 0
	config.Database.Save(ctx.Player.Status)
	ctx.Reply(discord.ReplyParams{
		Content:   "Your dungeon should be available now.",
		Ephemeral: true,
	})
}
