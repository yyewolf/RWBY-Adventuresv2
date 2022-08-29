package commands_boxes

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"gorm.io/gorm"
)

func OpenRareGrimm(ctx *discord.CmdContext) {
	if ctx.Player.Boxes.RareGrimmBoxes <= 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have enough **Rare Grimm Boxes** to do that.",
			Ephemeral: true,
		})
		return
	}
	b := &BoxFilter{
		Box:        "Rare Grimm Box",
		ValStd:     15,
		ValMean:    55,
		RarityRate: 1.175,
	}
	if OpenGrimm(ctx, b) {
		config.Database.Model(ctx.Player.Boxes).Update("rare_char_boxes", gorm.Expr("rare_char_boxes - ?", 1))
	}
}
