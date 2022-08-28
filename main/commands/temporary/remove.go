package commands_temporary

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

func Delete(ctx *discord.CmdContext) {
	config.Database.Delete(ctx.Player.Status, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Missions, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Shop, ctx.Author.ID)
	config.Database.Delete(ctx.Player.LastBoxes, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Gamble, ctx.Author.ID)
	config.Database.Delete(ctx.Player.LimitedBoxes, ctx.Author.ID)
	config.Database.Delete(ctx.Player.SpecialBoxes, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Boxes, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Stats, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Badges, ctx.Author.ID)

	for _, a := range ctx.Player.Market.Auctions {
		config.Database.Select("Stats").Where("user_id=?", a.ID).Delete(ctx.Player.Characters)
		config.Database.Select("Stats").Where("user_id=?", a.ID).Delete(ctx.Player.Grimms)
		config.Database.Delete(a.Bidders, a.ID)
	}
	for _, l := range ctx.Player.Market.Listings {
		config.Database.Select("Stats").Where("user_id=?", l.ID).Delete(ctx.Player.Characters)
		config.Database.Select("Stats").Where("user_id=?", l.ID).Delete(ctx.Player.Grimms)
	}

	trades := []models.Trade{}
	config.Database.Joins("UserSends", `"UserSends"."trade_id"=trades.trade_id AND "UserSends"."user_id"=?`, ctx.Author.ID).Joins("TargetSends", `"TargetSends"."trade_id"=trades.trade_id`).Find(&trades)
	for _, trade := range trades {
		config.Database.Delete(trade)
		config.Database.Delete(trade.TargetSends)
		config.Database.Delete(trade.UserSends)
	}

	trades = []models.Trade{}
	config.Database.Joins("TargetSends", `"TargetSends"."trade_id"=trades.trade_id AND "TargetSends"."user_id"=?`, ctx.Author.ID).Joins("UserSends", `"UserSends"."trade_id"=trades.trade_id`).Find(&trades)
	for _, trade := range trades {
		config.Database.Delete(trade)
		config.Database.Delete(trade.TargetSends)
		config.Database.Delete(trade.UserSends)
	}

	config.Database.Delete(ctx.Player.Market.Auctions, ctx.Author.ID)
	config.Database.Delete(ctx.Player.Market.Listings, ctx.Author.ID)

	config.Database.Select("Stats").Where("user_id=?", ctx.Author.ID).Delete(ctx.Player.Characters)
	config.Database.Select("Stats").Where("user_id=?", ctx.Author.ID).Delete(ctx.Player.Grimms)
	config.Database.Where("discord_id=?", ctx.Author.ID).Delete(ctx.Player)
	ctx.Reply(discord.ReplyParams{
		Content: "Deleted all your infos. F",
	})
}
