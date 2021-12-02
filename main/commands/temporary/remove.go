package commands_temporary

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
)

func Delete(ctx *discord.CmdContext) {
	config.Database.Delete(ctx.Player, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Status, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Missions, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Shop, ctx.Author.ID)
	config.Database.Delete(ctx.Player.LastBoxes, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Gamble, ctx.Author.ID)
	config.Database.Delete(ctx.Player.LimitedBoxes, ctx.Author.ID)
	config.Database.Delete(ctx.Player.SpecialBoxes, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Boxes, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Characters, "user_id = ?", ctx.Author.ID)
	config.Database.Delete(ctx.Player.Grimms, "user_id = ?", ctx.Author.ID)
	ctx.Reply(discord.ReplyParams{
		Content: "Deleted all your infos.",
	})
}
