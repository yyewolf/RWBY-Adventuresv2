package commands

import (
	"rwby-adventures/config"
	"rwby-adventures/discord"

	"gorm.io/gorm/clause"
)

func Delete(ctx *discord.CmdContext) {
	config.Database.Select(clause.Associations).Delete(ctx.Player, ctx.Author.ID)
	config.Database.Delete(ctx.Player, ctx.Author.ID)
	ctx.Reply(discord.ReplyParams{
		Content: "Ok Boomer.",
	})
}
