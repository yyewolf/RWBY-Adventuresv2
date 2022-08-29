package commands_boxes

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

func OpenLimited(ctx *discord.CmdContext) {
	if len(ctx.Player.LimitedBoxes) <= 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have enough **Limited Boxes** to do that.",
			Ephemeral: true,
		})
		return
	}
	b := &BoxFilter{
		IncludeLimited: true,
		OnlyLimited:    true,
		Box:            "Limited Box",
		ValStd:         15,
		ValMean:        58,
		RarityRate:     1.1,
	}
	if ctx.Player.LimitedBoxes[0].Type == models.CharType {
		if OpenChar(ctx, b) {
			config.Database.Delete(ctx.Player.LimitedBoxes, "for=?", ctx.Player.LimitedBoxes[0].For)
		}
	} else {
		if OpenGrimm(ctx, b) {
			config.Database.Delete(ctx.Player.LimitedBoxes, "for=?", ctx.Player.LimitedBoxes[0].For)
		}
	}
}
