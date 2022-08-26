package commands_shop

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

func Buy(ctx *discord.CmdContext) {
	arg := ctx.Arguments.GetArg("item", 0, 0)
	number := arg.Raw.IntValue()

	embed := &discordgo.MessageEmbed{
		Title: "",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.imgur.com/vGBfys2.png",
		},
		Description: "You pleased the shopkeeper !",
		Color:       config.Botcolor,
	}

	var err error

	switch number {
	case 1:
		embed.Title, err = buyXPBoost(ctx.Player)
		break
	case 2:
		embed.Title, err = buyLuckBoost(ctx.Player)
		break
	case 3:
		embed.Title, err = buyLootbox(ctx.Player)
		break
	case 4:
		embed.Title, err = buyGrimmbox(ctx.Player)
		break
	case 5:
		embed.Title, err = buyRarelootbox(ctx.Player)
		break
	case 6:
		embed.Title, err = buyRareGrimmbox(ctx.Player)
		break
	case 7:
		embed.Title, err = exchangeRareGrimmbox(ctx.Player)
		break
	case 8:
		embed.Title, err = exchangeRareLootbox(ctx.Player)
		break
	case 9:
		embed.Title, err = buyBackpack(ctx.Player)
		break
	default:
		ctx.Reply(discord.ReplyParams{
			Content: "You didn't choose a valid item.",
		})
		return
	}
	if err != nil {
		embed = &discordgo.MessageEmbed{
			Title: embed.Title,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://static.wikia.nocookie.net/rwby/images/8/80/Shopkeep_infobox_2.png/revision/latest?cb=20190926133636",
			},
			Description: "You displeased the shopkeeper !",
			Color:       config.Botcolor,
		}
	}

	ctx.Reply(discord.ReplyParams{
		Content: embed,
	})
}
