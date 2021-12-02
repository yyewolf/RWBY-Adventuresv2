package commands_temporary

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"

	uuid "github.com/satori/go.uuid"
)

func Addchar(ctx *discord.CmdContext) {
	ID := uuid.NewV4().String()
	config.Database.Save(&models.Character{
		CharID: ID,
		Level:  1,
		Rarity: 3,
		XP:     400,
		XPCap:  5000,
		Name:   "Weiss Schnee",
		UserID: ctx.Author.ID,
		Stats: models.CharacterStats{
			CharID: ID,
			Value:  52.32,
		},
	})
	ctx.Reply(discord.ReplyParams{
		Content: "Ok boomer.",
	})
}
