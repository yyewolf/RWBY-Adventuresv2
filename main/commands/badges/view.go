package commands_badges

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

func BadgesView(ctx *discord.CmdContext) {
	//Initialize embed and sends the message
	embed := &discordgo.MessageEmbed{
		Title:       ctx.Author.Username + "'s badges :",
		Description: "Sorry, you haven't got any badge yet :c",
		Color:       config.Botcolor,
		Footer:      discord.DefaultFooter,
	}

	howMuch := 0
	//Prepares the badges fields
	for _, b := range ctx.Player.Badges {
		if b.Badge == nil {
			continue
		}
		field := &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s **%s**", b.Badge.Emoji, b.Badge.Name),
			Value:  b.Badge.Description,
			Inline: true,
		}
		embed.Fields = append(embed.Fields, field)
		embed.Description = ""
		howMuch++
	}
	embed.Title = fmt.Sprintf("%s's badges (**%d**/%d):", ctx.Author.Username, howMuch, len(models.DefaultBadges))

	ctx.Reply(discord.ReplyParams{
		Content: embed,
	})
}
