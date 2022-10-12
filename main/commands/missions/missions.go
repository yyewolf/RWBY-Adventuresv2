package commands_missions

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

var MissionCommand = &discord.Command{
	Name:        "mission",
	Description: "Commands regarding missions.",
	Menu:        discord.InventoryMenu,
	SubCommands: []*discord.Command{
		{
			Name:        "go",
			Description: "Starts a mission (with your selected character).",
			Menu:        discord.InventoryMenu,
			Call:        joinMission,
		},
		{
			Name:        "check",
			Description: "Checks your current mission status.",
			Menu:        discord.InventoryMenu,
			Call:        checkMission,
		},
		{
			Name:        "end",
			Description: "Ends a mission (you will not get any reward).",
			Menu:        discord.InventoryMenu,
			Call:        endMission,
		},
	},
}

func joinMission(ctx *discord.CmdContext) {
	//Check if the player is able to do a mission first
	if !ctx.Player.Missions.CanGoToMission {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have any missions.",
			Ephemeral: true,
		})
		return
	}
	//Check if the player has enough characters
	if len(ctx.Player.Characters) <= 1 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have enough characters to send one in a mission.",
			Ephemeral: true,
		})
		return
	}
	//This is to avoid multiple characters in mission
	if ctx.Player.Missions.IsInMission {
		ctx.Reply(discord.ReplyParams{
			Content:   "You already have a character in mission.",
			Ephemeral: true,
		})
		return
	}
	//This is to avoid multiple characters in mission
	if ctx.Player.SelectedChar == nil {
		ctx.Reply(discord.ReplyParams{
			Content:   "You must select a character in order to do that.",
			Ephemeral: true,
		})
		return
	}

	//Puts every variable to avoid any glitch (needed if caching)

	ctx.Player.CharInMission = ctx.Player.SelectedChar
	ctx.Player.CharInMission.InMission = true
	for _, char := range ctx.Player.Characters {
		if char.CharID != ctx.Player.SelectedID {
			ctx.Player.SelectedID = char.CharID
			ctx.Player.SelectedChar = char
			ctx.Player.SelectedType = models.CharType
			break
		}
	}

	ctx.Reply(discord.ReplyParams{
		Content:   "Your character went on a mission (it's now **out** of your inventory for a while !).",
		Ephemeral: true,
	})
	config.Database.Save(ctx.Player.CharInMission)
	ctx.Player.Missions.CanGoToMission = false
	ctx.Player.Missions.IsInMission = true
	config.Database.Save(ctx.Player.Missions)
	config.Database.Save(ctx.Player)
}

func endMission(ctx *discord.CmdContext) {
	if !ctx.Player.Missions.IsInMission {
		ctx.Reply(discord.ReplyParams{
			Content:   "You are not currently on a mission.",
			Ephemeral: true,
		})
		return
	}

	ctx.Reply(discord.ReplyParams{
		Content:   "Your character has been returned to you.",
		Ephemeral: true,
	})
	ctx.Player.Missions.IsInMission = false
	config.Database.Save(ctx.Player.Missions)
	ctx.Player.CharInMission.InMission = false
	config.Database.Save(ctx.Player.CharInMission)
	ctx.Player.SelectedChar = ctx.Player.CharInMission
	ctx.Player.SelectedID = ctx.Player.CharInMission.CharID
	ctx.Player.SelectedType = models.CharType
	config.Database.Save(ctx.Player)
}

func checkMission(ctx *discord.CmdContext) {
	if !ctx.Player.Missions.CanGoToMission {
		ctx.Reply(discord.ReplyParams{
			Content:   "You are not able to go on a mission.",
			Ephemeral: true,
		})
		return
	}
	days := int(math.Ceil(float64(ctx.Player.Missions.MissionMsgLeft) / 24))

	ctx.Reply(discord.ReplyParams{
		Content: "You do have a mission awaiting !\n" +
			fmt.Sprintf("Use </mission go:%s> to go on a mission, %s!\n", ctx.Command.ID, ctx.Author.Mention()) +
			fmt.Sprintf("**%s**.\n", missionToString(ctx.Player.Missions.MissionType)) +
			"Your character will leave your inventory and come back after a while.\n" +
			fmt.Sprintf("Your character will be gone for : **%d days**.\nTo end a mission by yourself use </mission end:%s>.", days, ctx.Command.ID),
		Ephemeral: true,
	})
}
