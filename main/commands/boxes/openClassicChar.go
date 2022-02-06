package commands_boxes

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"gorm.io/gorm"
)

func OpenClassicChar(ctx *discord.CmdContext) {
	if ctx.Player.Boxes.Boxes <= 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have enough **Classic Character Boxes** to do that.",
			Ephemeral: true,
		})
		return
	}
	b := &BoxFilter{
		Box:        "Classic Character Box",
		ValStd:     15,
		ValMean:    48.5,
		RarityRate: 1.25,
	}
	if OpenChar(ctx, b) {
		config.Database.Model(ctx.Player.Boxes).Update("classic_char_boxes", gorm.Expr("classic_char_boxes - ?", 1))
	}
}
