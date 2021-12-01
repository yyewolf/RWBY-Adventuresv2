package commands

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

func Addchar(ctx *discord.CmdContext) {
	config.Database.Save(&models.Character{
		CharID: "0",
		UserID: ctx.Message.Author.ID,
		Stats: models.CharacterStats{
			CharID: "0",
			Value:  52.32,
		},
	})
	ctx.Reply(discord.ReplyParams{
		Content: "Ok boomer.",
	})
}
