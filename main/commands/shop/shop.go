package commands_shop

import (
	"errors"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Shop(ctx *discord.CmdContext) {
	g := ctx.Guild
	BPrice := int64(12000*(ctx.Player.Shop.Extensions+1) + 5000)
	embed := &discordgo.MessageEmbed{
		Title:       "Shop :",
		Color:       config.Botcolor,
		Description: "To buy an item use : `" + g.Prefix + "buy {item}`.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "XP Boost (N°1)",
				Value:  "Augment XP for a certain amount of time.\nPrice : 325Ⱡ",
				Inline: true,
			},
			{
				Name:   "Luck Boost (N°2)",
				Value:  "Augment your luck for 15 lucky actions.\nPrice : 650Ⱡ",
				Inline: true,
			},
			{
				Name:   "Classic Loot box (N°3)",
				Value:  "Basically 2 loot boxes .\nPrice : 1750Ⱡ",
				Inline: true,
			},
			{
				Name:   "Classic Grimm box (N°4)",
				Value:  "Basically 2 grimm boxes .\nPrice : 2500Ⱡ",
				Inline: true,
			},
			{
				Name:   "Rare Loot box (N°5)",
				Value:  "3 Rare loot boxes.\nPrice : 5000Ⱡ",
				Inline: true,
			},
			{
				Name:   "Rare Grimm box (N°6)",
				Value:  "3 Rare grimm boxes.\nPrice : 7500Ⱡ",
				Inline: true,
			},
			{
				Name:   "Rare loot box (N°7)",
				Value:  "1 Rare loot box.\nPrice : 10 Loot Boxes",
				Inline: true,
			},
			{
				Name:   "Rare Grimm box (N°8)",
				Value:  "1 Rare grimm box.\nPrice : 10 Grimm Boxes",
				Inline: true,
			},
			{
				Name:   "Backpack Expension (N°9)",
				Value:  "20 more characters in your inventory.\nYour price : " + strconv.FormatInt(BPrice, 10) + "Ⱡ\n",
				Inline: true,
			},
		},
		Footer: discord.DefaultFooter,
	}

	ctx.Reply(discord.ReplyParams{
		Content: embed,
	})
}

func buyXPBoost(p *models.Player) (string, error) {
	if p.TotalBalance() < 325 {
		return "You don't have enough money.", errors.New("err")
	}
	if p.Shop.XPBoostTime > 1 {
		return "You already have an XP boost.", errors.New("err")
	}
	p.Shop.XPBoost = true
	p.Shop.XPBoostTime = 100
	p.Balance -= 325

	p.Save()
	p.Shop.Save()
	return "You bought a XP Booster.", nil
}

func buyLuckBoost(p *models.Player) (string, error) {
	if p.TotalBalance() < 650 {
		return "You don't have enough money.", errors.New("err")
	}
	if p.Shop.LuckBoostTime > 1 {
		return "You already have a luck boost.", errors.New("err")
	}
	p.Shop.LuckBoost = true
	p.Shop.LuckBoostTime = 20
	p.Balance -= 650

	p.Save()
	p.Shop.Save()
	return "You bought a Luck Booster.", nil
}
