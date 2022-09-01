package commands_market

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

func RemoveListing(ctx *discord.CmdContext) {
	arg := ctx.Arguments.GetArg("id", 0, "")
	id := arg.Raw.StringValue()
	l, err := models.GetListing(id)
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageEmbed{
				Title:       "Remove Listing",
				Description: "Couldn't find this listing.",
				Color:       config.Botcolor,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: ctx.Author.AvatarURL("512"),
				},
				Footer: discord.DefaultFooter,
			},
		})
		return
	}

	if l.SellerID != ctx.Author.ID {
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageEmbed{
				Title:       "Remove Listing",
				Description: "Couldn't find this listing.",
				Color:       config.Botcolor,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: ctx.Author.AvatarURL("512"),
				},
				Footer: discord.DefaultFooter,
			},
		})
		return
	}

	var personaString string
	if l.Type == models.CharType {
		l.Char.UserID = ctx.Author.ID
		personaString = l.Char.FullString()
		l.Char.Save()
	} else {
		l.Grimm.UserID = ctx.Author.ID
		personaString = l.Grimm.FullString()
		l.Grimm.Save()
	}

	l.Delete()

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       "Remove Listing",
			Description: fmt.Sprintf("You retrieved your `%s`.", personaString),
			Color:       config.Botcolor,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL("512"),
			},
			Footer: discord.DefaultFooter,
		},
	})
}
