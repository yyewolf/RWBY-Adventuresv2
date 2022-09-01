package commands_market

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ViewAuctions(ctx *discord.CmdContext) {
	if len(ctx.Player.Market.Listings) == 0 {
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageEmbed{
				Title:       ctx.Author.Username + "'s auctions & bid :",
				Description: "You don't have any auctions at the moment.",
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
	maxpage := int64(math.Ceil(float64(len(ctx.Player.Market.Auctions)) / 10))
	if page > maxpage {
		page = maxpage
	}

	fields := []*discordgo.MessageEmbedField{}
	for i, a := range ctx.Player.Market.Auctions {
		//Correct interval : [(p-1)*10, p*10]
		if int64(i) < (page-1)*10 || int64(i) >= page*10 {
			continue
		}

		a.Ended = a.EndsAt < time.Now().Unix()
		if a.Ended {
			a.End()
			if len(a.Bidders) > 0 {
				var personaString string
				if a.Type == models.CharType {
					personaString = a.Char.FullString()
				} else {
					personaString = a.Grimm.FullString()
				}

				go func() {
					var ch *discordgo.Channel
					var err error
					// Create User DM
					if ch, err = discord.Session.UserChannelCreate(a.SellerID); err != nil {
						return
					}
					discord.Session.ChannelMessageSendEmbed(ch.ID, &discordgo.MessageEmbed{
						Title:       "Auction Ended",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("Your auction on `%s` has ended, you earned : **%d** Liens.", personaString, a.Bidders[0].Bid),
					})
				}()

				go func() {
					var ch *discordgo.Channel
					var err error
					// Create User DM
					if ch, err = discord.Session.UserChannelCreate(a.Bidders[0].UserID); err != nil {
						return
					}
					discord.Session.ChannelMessageSendEmbed(ch.ID, &discordgo.MessageEmbed{
						Title:       "Auction Ended",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("You purchased `%s` for **%d** Liens on an auction.", personaString, a.Bidders[0].Bid),
					})
				}()
			} else {
				var personaString string
				if a.Type == models.CharType {
					personaString = a.Char.FullString()
				} else {
					personaString = a.Grimm.FullString()
				}

				go func() {
					var ch *discordgo.Channel
					var err error
					// Create User DM
					if ch, err = discord.Session.UserChannelCreate(a.SellerID); err != nil {
						return
					}
					discord.Session.ChannelMessageSendEmbed(ch.ID, &discordgo.MessageEmbed{
						Title:       "Auction Ended",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("Your auction on `%s` has ended, nobody placed a bid.", personaString),
					})
				}()
			}
			continue
		}

		var str string
		if a.SellerID == ctx.Author.ID {
			if a.Type == 0 {
				str = fmt.Sprintf("`%s at %d Liens`", a.Char.FullString(), a.Bidders[0].Bid)
			} else {
				str = fmt.Sprintf("`%s at %d Liens`", a.Grimm.FullString(), a.Bidders[0].Bid)
			}
		} else {
			if a.Type == 0 {
				str = fmt.Sprintf("`%s where YOU bidded %d Liens`", a.Char.FullString(), a.Bidders[0].Bid)
			} else {
				str = fmt.Sprintf("`%s where YOU bidded %d Liens`", a.Grimm.FullString(), a.Bidders[0].Bid)
			}
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("ID : `%s`", a.ID),
			Value: str,
		})
	}
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s's auctions (page %d/%d)", ctx.Author.Username, page, maxpage),
			Color:       config.Botcolor,
			Description: "These are your market auctions.\nTo view another page type `/market auctions view {page}`.",
			Fields:      fields,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL("512"),
			},
			Footer: discord.DefaultFooter,
		},
	})
}
