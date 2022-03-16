package commands_missions

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

var HuntCommand = &discord.Command{
	Name:        "hunt",
	Description: "Commands regarding hunts.",
	Menu:        discord.InventoryMenu,
	SubCommands: []*discord.Command{
		{
			Name:        "go",
			Description: "Starts a hunt (with your selected grimm).",
			Menu:        discord.InventoryMenu,
			Call:        joinHunt,
		},
		{
			Name:        "check",
			Description: "Checks your current hunt status.",
			Menu:        discord.InventoryMenu,
			Call:        checkHunt,
		},
		{
			Name:        "end",
			Description: "Ends a hunt (you will not get any reward).",
			Menu:        discord.InventoryMenu,
			Call:        endHunt,
		},
	},
}

func joinHunt(ctx *discord.CmdContext) {
	//Check if the player is able to do a hunt first
	if !ctx.Player.Missions.CanGoHunt {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have any hunts.",
			Ephemeral: true,
		})
		return
	}
	//Check if the player has enough grimms
	if len(ctx.Player.Grimms) <= 1 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have enough grimms to send one in a hunt.",
			Ephemeral: true,
		})
		return
	}
	//This is to avoid multiple grimms in hunt
	if ctx.Player.Missions.IsInHunt {
		ctx.Reply(discord.ReplyParams{
			Content:   "You already have a grimm in hunt.",
			Ephemeral: true,
		})
		return
	}
	//This is to avoid multiple grimms in hunt
	if ctx.Player.SelectedChar == nil {
		ctx.Reply(discord.ReplyParams{
			Content:   "You must select a grimm in order to do that.",
			Ephemeral: true,
		})
		return
	}

	//Puts every variable to avoid any glitch (needed if caching)
	ctx.Player.SelectedGrimm.InHunt = true

	ctx.Player.GrimmInHunt = ctx.Player.SelectedGrimm
	for _, grimm := range ctx.Player.Grimms {
		if grimm.GrimmID != ctx.Player.SelectedGrimm.GrimmID {
			ctx.Player.SelectedID = grimm.GrimmID
			ctx.Player.SelectedGrimm = grimm
			ctx.Player.SelectedType = models.GrimmType
			break
		}
	}

	ctx.Reply(discord.ReplyParams{
		Content:   "Your grimm went on a hunt (it's now **out** of your inventory for a while !).",
		Ephemeral: true,
	})
	config.Database.Save(ctx.Player.SelectedGrimm)
	ctx.Player.Missions.CanGoHunt = false
	ctx.Player.Missions.IsInHunt = true
	config.Database.Save(ctx.Player.Missions)
	config.Database.Save(ctx.Player)
}

func endHunt(ctx *discord.CmdContext) {
	if !ctx.Player.Missions.IsInHunt {
		ctx.Reply(discord.ReplyParams{
			Content:   "Your are not currently on a hunt.",
			Ephemeral: true,
		})
		return
	}

	ctx.Reply(discord.ReplyParams{
		Content:   "Your grimm has been returned to you.",
		Ephemeral: true,
	})
	ctx.Player.Missions.IsInHunt = false
	config.Database.Save(ctx.Player.Missions)
	ctx.Player.GrimmInHunt.InHunt = false
	config.Database.Save(ctx.Player.GrimmInHunt)
	ctx.Player.SelectedGrimm = ctx.Player.GrimmInHunt
	ctx.Player.SelectedID = ctx.Player.GrimmInHunt.GrimmID
	ctx.Player.SelectedType = models.GrimmType
	config.Database.Save(ctx.Player)
}

func checkHunt(ctx *discord.CmdContext) {
	if !ctx.Player.Missions.CanGoHunt {
		ctx.Reply(discord.ReplyParams{
			Content:   "You are not able to go on a hunt.",
			Ephemeral: true,
		})
		return
	}
	days := int(math.Ceil(float64(ctx.Player.Missions.HuntMsgLeft) / 24))

	ctx.Reply(discord.ReplyParams{
		Content: "You do have a hunt awaiting !\n" +
			fmt.Sprintf("Type `%shunt go` to go on a hunt, %s!\n", ctx.Guild.Prefix, ctx.Author.Mention()) +
			fmt.Sprintf("**%s**.\n", huntToString(ctx.Player.Missions.HuntType)) +
			"Your grimm will leave your inventory and come back after a while.\n" +
			fmt.Sprintf("Your grimm will be gone for : **%d days**.\nTo end a hunt by yourself type `%shunt end`.", days, ctx.Guild.Prefix),
		Ephemeral: true,
	})
}
