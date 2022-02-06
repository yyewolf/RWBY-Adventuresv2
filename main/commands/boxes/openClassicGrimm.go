package commands_boxes

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"gorm.io/gorm"
)

func OpenClassicGrimm(ctx *discord.CmdContext) {
	if ctx.Player.Boxes.GrimmBoxes <= 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have enough **Classic Grimm Boxes** to do that.",
			Ephemeral: true,
		})
		return
	}
	b := &BoxFilter{
		Box:        "Classic Grimm Box",
		ValStd:     15,
		ValMean:    48.5,
		RarityRate: 1.25,
	}
	if OpenGrimm(ctx, b) {
		config.Database.Model(ctx.Player.Boxes).Update("classic_grimm_boxes", gorm.Expr("classic_grimm_boxes - ?", 1))
	}
}
