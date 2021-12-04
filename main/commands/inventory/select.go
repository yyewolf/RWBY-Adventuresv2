package commands_inventory

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

var SelectCommand = &discord.Command{
	Name:        "select",
	Description: "Select a persona.",
	Menu:        discord.GeneralMenu,
	Call:        SelectPersona,
	Args: []discord.Arg{
		{
			Name:        "id",
			Description: "Identification number of your persona.",
			Size:        1,
			Required:    false,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "latest",
			Description: "Whether or not you want to view the infos of your latest persona.",
			Size:        1,
			Required:    false,
			Type:        discordgo.ApplicationCommandOptionBoolean,
		},
	},
}

type SelectData struct {
	Char     *models.Character
	Grimm    *models.Grimm
	isGrimm  bool
	FollowUp bool
}

func SelectPersona(ctx *discord.CmdContext) {
	// We parse user input
	var pickLatest bool
	latest, err := ctx.Arguments.GetArg("latest", 1)
	if err == nil {
		if v, ok := latest.Value.(bool); ok {
			pickLatest = v
		}
	}

	id, err := ctx.Arguments.GetArg("id", 0)
	if err != nil && !pickLatest {
		ctx.Reply(discord.ReplyParams{
			Content:   "You need to input at least the ID of the persona you wish to Select.",
			Ephemeral: true,
		})
		return
	}
	isGrimm, index, err := id.CharGrimmParse()
	if err != nil && !pickLatest {
		ctx.Reply(discord.ReplyParams{
			Content:   "I did not understand the ID that you sent.",
			Ephemeral: true,
		})
		return
	}

	// We search for the character they sent

	var grimm *models.Grimm
	var char *models.Character

	if isGrimm {
		if index > len(ctx.Player.Grimms) {
			ctx.Reply(discord.ReplyParams{
				Content:   "You don't have any grimm with this number.",
				Ephemeral: true,
			})
			return
		}
		grimm = &ctx.Player.Grimms[index-1]
	} else {
		if index > len(ctx.Player.Characters) {
			ctx.Reply(discord.ReplyParams{
				Content:   "You don't have any character with this number.",
				Ephemeral: true,
			})
			return
		}
		char = &ctx.Player.Characters[index-1]
	}

	if pickLatest {
		isGrimm, char, grimm, index = ctx.Player.GetLatestPersona()
		index += 1
	}
	// We proceed to select the character

	Select(ctx, &SelectData{
		Grimm:   grimm,
		Char:    char,
		isGrimm: isGrimm,
	})
}

func Select(ctx *discord.CmdContext, data *SelectData) {
	var reply string

	if data.isGrimm {
		ctx.Player.SelectedType = models.GrimmType
		ctx.Player.SelectedID = data.Grimm.GrimmID
		reply = fmt.Sprintf("You have selected :\n%s", data.Grimm.FullString())
	} else {
		ctx.Player.SelectedType = models.CharType
		ctx.Player.SelectedID = data.Char.CharID
		reply = fmt.Sprintf("You have selected :\n%s", data.Char.FullString())
	}
	config.Database.Save(ctx.Player)

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageSend{
			Content: reply,
		},
		FollowUp: data.FollowUp,
	})
}
