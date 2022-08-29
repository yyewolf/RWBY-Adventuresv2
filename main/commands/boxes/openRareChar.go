package commands_boxes

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"gorm.io/gorm"
)

func OpenRareChar(ctx *discord.CmdContext) {
	if ctx.Player.Boxes.RareBoxes <= 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have enough **Rare Character Boxes** to do that.",
			Ephemeral: true,
		})
		return
	}
	b := &BoxFilter{
		Box:        "Rare Character Box",
		ValStd:     15,
		ValMean:    55,
		RarityRate: 1.175,
	}
	if OpenChar(ctx, b) {
		config.Database.Model(ctx.Player.Boxes).Update("rare_char_boxes", gorm.Expr("rare_char_boxes - ?", 1))
	}
}
