package commands_inventory

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

type AddBuffData struct {
	Char     *models.Character
	Grimm    *models.Grimm
	isGrimm  bool
	FollowUp bool
}

func AddBuff(ctx *discord.CmdContext, data *AddBuffData) {
	var reply string

	if data.isGrimm {
		if ctx.Player.Minions <= 0 {
			ctx.Reply(discord.ReplyParams{
				Content:   "You don't have any minions to add.",
				FollowUp:  data.FollowUp,
				Ephemeral: true,
			})
			return
		}
		if data.Grimm.Buffs >= config.MaxBuffs {
			reply = fmt.Sprintf("This grimm already has **%d** minions.", config.MaxBuffs)
		} else {
			data.Grimm.Buffs++
			ctx.Player.Minions--
			reply = fmt.Sprintf("You have added a minion to : `%s`.", data.Grimm.FullString())
			config.Database.Save(data.Grimm)
			config.Database.Save(ctx.Player)
		}
	} else {
		if ctx.Player.Arms <= 0 {
			ctx.Reply(discord.ReplyParams{
				Content:   "You don't have any arms to add.",
				FollowUp:  data.FollowUp,
				Ephemeral: true,
			})
			return
		}
		if data.Char.Buffs >= config.MaxBuffs {
			reply = fmt.Sprintf("This char already has **%d** arms.", config.MaxBuffs)
		} else {
			data.Char.Buffs++
			ctx.Player.Arms--
			reply = fmt.Sprintf("You have added an arm to : `%s`.", data.Char.FullString())
			config.Database.Save(data.Char)
			config.Database.Save(ctx.Player)
		}
	}

	ctx.Reply(discord.ReplyParams{
		Content:   reply,
		FollowUp:  data.FollowUp,
		Ephemeral: true,
	})
}
