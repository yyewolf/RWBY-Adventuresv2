package commands_inventory

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
	var char *models.Character
	var grimm *models.Grimm
	arg, _ := ctx.Arguments.GetArg("id", 0)
	isGrimm, index, err := arg.CharGrimmParse()
	if err != nil {
		char = &ctx.Player.SelectedChar
		grimm = &ctx.Player.SelectedGrimm
		isGrimm = ctx.Player.SelectedType == models.GrimmType
		if isGrimm {
			for i, g := range ctx.Player.Grimms {
				if g.GrimmID == grimm.GrimmID {
					index = i + 1
					break
				}
			}
		} else {
			for i, c := range ctx.Player.Characters {
				if c.CharID == char.CharID {
					index = i + 1
					break
				}
			}
		}
	} else {
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
	}
	latest, err := ctx.Arguments.GetArg("latest", 1)
	if err == nil {
		if v, ok := latest.Value.(bool); ok && v {
			isGrimm, char, grimm, index = ctx.Player.GetLatestChar()
			index += 1
		}
	}
	if isGrimm {
		grimmInfo(ctx, grimm, index)
		return
	}
	charInfo(ctx, char, index)
}

func charInfo(ctx *discord.CmdContext, char *models.Character, number int) {
	// Useful data
	original := char.ToRealChar()
	imgData, _ := static.CharBox.ReadFile(original.ImageFile)
	imgDecoded := bytes.NewBuffer(imgData)

	// DiscordGo Stuff

	Complex := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Reader: imgDecoded,
				Name:   "ch.png",
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title: "Level " + strconv.Itoa(char.Level) + " " + char.Name + ". (n°" + strconv.Itoa(number) + ")",
			Color: models.CharRarityToColor(char.Rarity),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: "**Statistics :**",
					Value: "Category : **" + original.Category + "**\n" +
						"XP : " + strconv.FormatInt(char.XP, 10) + "/" + strconv.FormatInt(char.XPCap, 10) + "\n" +
						"Value : " + fmt.Sprintf("%.2f", char.Stats.Value) + "%\n" +
						"Rarity : " + char.RarityString() + "\n" +
						"Health : " + strconv.Itoa(char.Stats.Health) + "\n" +
						"Armor : " + strconv.Itoa(char.Stats.Armor) + "\n" +
						"Damage : " + strconv.Itoa(char.Stats.Damage),
				},
				{
					Name:  "**Special thanks to :**",
					Value: original.ImageAuthors,
				},
			},
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

func grimmInfo(ctx *discord.CmdContext, grimm *models.Grimm, number int) {
	// Useful data
	original := grimm.ToRealGrimm()
	imgData, _ := static.CharBox.ReadFile(original.ImageFile)
	imgDecoded := bytes.NewBuffer(imgData)

	// DiscordGo Stuff

	Complex := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Reader: imgDecoded,
				Name:   "ch.png",
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title: "Level " + strconv.Itoa(grimm.Level) + " " + grimm.Name + ". (n°" + strconv.Itoa(number) + ")",
			Color: models.CharRarityToColor(grimm.Rarity),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: "**Statistics :**",
					Value: "Category : **" + original.Category + "**\n" +
						"XP : " + strconv.FormatInt(grimm.XP, 10) + "/" + strconv.FormatInt(grimm.XPCap, 10) + "\n" +
						"Value : " + fmt.Sprintf("%.2f", grimm.Stats.Value) + "%\n" +
						"Rarity : " + grimm.RarityString() + "\n" +
						"Health : " + strconv.Itoa(grimm.Stats.Health) + "\n" +
						"Armor : " + strconv.Itoa(grimm.Stats.Armor) + "\n" +
						"Damage : " + strconv.Itoa(grimm.Stats.Damage),
				},
				{
					Name:  "**Special thanks to :**",
					Value: original.ImageAuthors,
				},
			},
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
