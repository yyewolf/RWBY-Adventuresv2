package commands_market

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

func ViewListings(ctx *discord.CmdContext) {
	if len(ctx.Player.Market.Listings) == 0 {
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageEmbed{
				Title:       ctx.Author.Username + "'s listings :",
				Description: "You don't have any listings at the moment.",
				Color:       config.Botcolor,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: ctx.Author.AvatarURL("512"),
				},
			},
		})
		return
	}

	//Takes out the page from the message
	arg := ctx.Arguments.GetArg("page", 0, 1)
	page := arg.Raw.IntValue()
	if page <= 0 {
		page = 1
	}
	maxpage := int64(math.Ceil(float64(len(ctx.Player.Market.Listings)) / 10))
	if page > maxpage {
		page = maxpage
	}

	fields := []*discordgo.MessageEmbedField{}
	for i, l := range ctx.Player.Market.Listings {
		//Correct interval : [(p-1)*10, p*10]
		if int64(i) < (page-1)*10 || int64(i) >= page*10 {
			continue
		}

		var str string
		if l.Type == 0 {
			str = fmt.Sprintf("`%s at %d Liens`", l.Char.FullString(), l.Price)
		} else {
			str = fmt.Sprintf("`%s at %d Liens`", l.Grimm.FullString(), l.Price)
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("ID : `%s`", l.ID),
			Value: str,
		})
	}
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s's listings (page %d/%d)", ctx.Author.Username, page, maxpage),
			Color:       config.Botcolor,
			Description: fmt.Sprintf("These are your market listings.\nTo view another page use </market listings view:%s>.", ctx.Command.ID),
			Fields:      fields,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL("512"),
			},
			Footer: discord.DefaultFooter,
		},
	})
}
