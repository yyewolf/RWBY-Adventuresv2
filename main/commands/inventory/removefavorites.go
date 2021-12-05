package commands_inventory

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

type RemFavoriteData struct {
	Char     *models.Character
	Grimm    *models.Grimm
	isGrimm  bool
	FollowUp bool
}

func RemoveFavoritePersona(ctx *discord.CmdContext) {
	// We parse user input
	var err error
	var pickLatest bool
	var isGrimm bool
	var index int
	var grimm *models.Grimm
	var char *models.Character
	var id *discord.CommandArg

	latest, err := ctx.Arguments.GetArg("latest", 1)
	if err == nil {
		if v, ok := latest.Value.(bool); ok && v {
			pickLatest = v
			goto skip
		}
	}

	id, err = ctx.Arguments.GetArg("id", 0)
	if err != nil && !pickLatest {
		ctx.Reply(discord.ReplyParams{
			Content:   "You need to input at least the ID of the persona you wish to add to your favorites.",
			Ephemeral: true,
		})
		return
	}
	isGrimm, index, err = id.CharGrimmParse()
	if err != nil && !pickLatest {
		ctx.Reply(discord.ReplyParams{
			Content:   "I did not understand the ID that you sent.",
			Ephemeral: true,
		})
		return
	}

	// We search for the character they sent

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

skip:
	if pickLatest {
		isGrimm, char, grimm, index, err = ctx.Player.GetLatestPersona()
		index += 1
		if err != nil {
			ctx.Reply(discord.ReplyParams{
				Content:   "You don't have any persona.",
				Ephemeral: true,
			})
		}
	}
	// We proceed to select the character

	RemoveFavorite(ctx, &RemFavoriteData{
		Grimm:   grimm,
		Char:    char,
		isGrimm: isGrimm,
	})
}

func RemoveFavorite(ctx *discord.CmdContext, data *RemFavoriteData) {
	var reply string

	if data.isGrimm {
		data.Grimm.IsInFavorites = false
		config.Database.Save(data.Grimm)
		reply = fmt.Sprintf("You have removed the following Grimm from your favorites :\n%s", data.Grimm.FullString())
	} else {
		data.Char.IsInFavorites = false
		config.Database.Save(data.Char)
		reply = fmt.Sprintf("You have removed the following Character from your favorites :\n%s", data.Char.FullString())
	}

	ctx.Reply(discord.ReplyParams{
		Content:   reply,
		FollowUp:  data.FollowUp,
		Ephemeral: true,
	})
}
