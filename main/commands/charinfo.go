package commands

import (
	"bytes"
	"fmt"
	"rwby-adventures/main/discord"
	"rwby-adventures/main/static"
	"rwby-adventures/models"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Info(ctx *discord.CmdContext) {
	CharInfo(ctx)
}

func CharInfo(ctx *discord.CmdContext) {
	PlayerChar := ctx.Player.Characters[0]
	RealChar := PlayerChar.ToRealChar()
	CharacterImageData, _ := static.CharBox.ReadFile("database/Images/" + RealChar.ImageFile)
	CharacterImageDecoded := bytes.NewBuffer(CharacterImageData)
	Image := &discordgo.File{
		Reader: CharacterImageDecoded,
		Name:   "ch.png",
	}
	MessageFields := []*discordgo.MessageEmbedField{}
	StatField := &discordgo.MessageEmbedField{
		Name: "**Statistics :**",
		Value: "Category : **" + RealChar.Category + "**\n" +
			"XP : " + strconv.FormatInt(PlayerChar.XP, 10) + "/" + strconv.FormatInt(PlayerChar.XPCap, 10) + "\n" +
			"Value : " + fmt.Sprintf("%.2f", PlayerChar.Stats.Value) + "%\n" +
			"Rarity : " + PlayerChar.RarityString() + "\n" +
			"Health : " + strconv.Itoa(PlayerChar.Stats.Health) + "\n" +
			"Armor : " + strconv.Itoa(PlayerChar.Stats.Armor) + "\n" +
			"Damage : " + strconv.Itoa(PlayerChar.Stats.Damage),
	}
	SpecialThanks := &discordgo.MessageEmbedField{
		Name:  "**Special thanks to :**",
		Value: RealChar.ImageAuthors,
	}
	MessageFields = append(MessageFields, StatField, SpecialThanks)
	Complex := &discordgo.MessageSend{
		Files: []*discordgo.File{Image},
		Embed: &discordgo.MessageEmbed{
			Title:       "Level " + strconv.Itoa(PlayerChar.Level) + " " + PlayerChar.Name + ". (n°" + strconv.Itoa(1) + ")",
			Description: "❎: Remove\n⭐: Add to favorites\n💥: Remove from favorites\n⛏️: Select (pick)",
			Color:       models.CharRarityToColor(PlayerChar.Rarity),
			Fields:      MessageFields,
			Image: &discordgo.MessageEmbedImage{
				URL: "attachment://ch.png",
			},
			Footer: discord.DefaultFooter,
		},
	}

	ctx.Reply(discord.ReplyParams{
		Content: Complex,
	})
}
