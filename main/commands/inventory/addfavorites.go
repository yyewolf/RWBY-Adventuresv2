package commands_inventory

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

var FavoritesCommand = &discord.Command{
	Name:        "favorites",
	Description: "All commands regarding favorites",
	SubCommands: []*discord.Command{
		{
			Name:        "add",
			Description: "Adds a persona to your favorites.",
			Menu:        discord.PersonasMenu,
			Call:        AddFavoritePersona,
			Args: []discord.Arg{
				{
					Name:        "id",
					Description: "Identification number of your persona.",
					Size:        1,
					Required:    true,
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
		},
		{
			Name:        "remove",
			Description: "Removes a persona from your favorites.",
			Menu:        discord.PersonasMenu,
			Call:        RemoveFavoritePersona,
			Args: []discord.Arg{
				{
					Name:        "id",
					Description: "Identification number of your persona.",
					Size:        1,
					Required:    true,
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
		},
	},
}

type AddFavoriteData struct {
	Char     *models.Character
	Grimm    *models.Grimm
	isGrimm  bool
	FollowUp bool
}

func AddFavoritePersona(ctx *discord.CmdContext) {
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
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content:   "You need to input at least the ID of the persona you wish to add to your favorites.",
			Ephemeral: true,
		})
		return
	}
	isGrimm, index, err = id.CharGrimmParse()
	if err != nil {
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

	AddFavorite(ctx, &AddFavoriteData{
		Grimm:   grimm,
		Char:    char,
		isGrimm: isGrimm,
	})
}

func AddFavorite(ctx *discord.CmdContext, data *AddFavoriteData) {
	var reply string

	if data.isGrimm {
		data.Grimm.IsInFavorites = true
		config.Database.Save(data.Grimm)
		reply = fmt.Sprintf("You have added the following Grimm to your favorites :\n%s", data.Grimm.FullString())
	} else {
		data.Char.IsInFavorites = true
		config.Database.Save(data.Char)
		reply = fmt.Sprintf("You have added the following Character to your favorites :\n%s", data.Char.FullString())
	}

	ctx.Reply(discord.ReplyParams{
		Content:   reply,
		FollowUp:  data.FollowUp,
		Ephemeral: true,
	})
}
