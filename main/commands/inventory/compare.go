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

var CompCommand = &discord.Command{
	Name:        "compare",
	Description: "Use this command to generate a fake character",
	Menu:        discord.InventoryMenu,
	Call:        Compare,
	Args: []discord.Arg{
		{
			Name:        "type",
			Description: "Type of your persona.",
			Size:        1,
			Required:    true,
			Choices: []*discord.Choice{
				{
					Name:  "Character",
					Value: models.CharType,
				},
				{
					Name:  "Grimm",
					Value: models.GrimmType,
				},
			},
			Type: discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "name",
			Description: "Name of your persona.",
			Size:        1,
			Required:    true,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "level",
			Description: "Level of your persona.",
			Size:        1,
			Required:    true,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "value",
			Description: "Value of your persona.",
			Size:        1,
			Required:    true,
			Type:        discordgo.ApplicationCommandOptionNumber,
		},
		{
			Name:        "rarity",
			Description: "Rarity of your persona.",
			Size:        1,
			Required:    true,
			Choices: []*discord.Choice{
				{
					Name:  "Common",
					Value: 0,
				},
				{
					Name:  "Uncommon",
					Value: 1,
				},
				{
					Name:  "Rare",
					Value: 2,
				},
				{
					Name:  "Very Rare",
					Value: 3,
				},
				{
					Name:  "Legendary",
					Value: 4,
				},
				{
					Name:  "Collector",
					Value: 5,
				},
			},
			Type: discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "buffs",
			Description: "Buffs of your persona.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
	},
}

func Compare(ctx *discord.CmdContext) {
	personaType := ctx.Arguments.GetArg("type", 0, 0).Value.(int64)
	personaName := ctx.Arguments.GetArg("name", 1, "").Value.(string)
	personaLevel := ctx.Arguments.GetArg("level", 2, 0).Value.(int64)
	personaValue := ctx.Arguments.GetArg("value", 3, 0).Value.(float64)
	personaRarity := ctx.Arguments.GetArg("rarity", 4, 0).Value.(int64)
	buff := ctx.Arguments.GetArg("buffs", 5, 0)
	var personaBuffs int64
	if buff.Found {
		personaBuffs = buff.Raw.IntValue()
	}

	var message *discordgo.MessageSend
	if personaType == models.CharType {
		char := &models.Character{
			Name:          personaName,
			Level:         int(personaLevel),
			XP:            0,
			XPCap:         0,
			Value:         personaValue,
			Rarity:        int(personaRarity),
			Buffs:         int(personaBuffs),
			IsInFavorites: false,
		}
		char.CalcStats()

		original := char.ToRealChar()
		imgData, _ := static.DatabaseFS.ReadFile(original.ImageFile)
		imgDecoded := bytes.NewBuffer(imgData)

		message = &discordgo.MessageSend{
			Files: []*discordgo.File{
				{
					Reader: imgDecoded,
					Name:   "ch.png",
				},
			},
			Embed: &discordgo.MessageEmbed{
				Title: "Level " + strconv.Itoa(char.Level) + " " + char.Name + ".",
				Color: char.RarityToColor(),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name: "**Statistics :**",
						Value: "Category : **" + original.Category + "**\n" +
							"XP : " + strconv.FormatInt(char.XP, 10) + "/" + strconv.FormatInt(char.XPCap, 10) + "\n" +
							"Value : " + fmt.Sprintf("%.2f", char.Value) + "%\n" +
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
	} else {
		grimm := &models.Grimm{
			Name:          personaName,
			Level:         int(personaLevel),
			XP:            0,
			XPCap:         0,
			Value:         personaValue,
			Rarity:        int(personaRarity),
			Buffs:         int(personaBuffs),
			IsInFavorites: false,
		}
		grimm.CalcStats()

		// Useful data
		original := grimm.ToRealGrimm()
		imgData, _ := static.DatabaseFS.ReadFile(original.ImageFile)
		imgDecoded := bytes.NewBuffer(imgData)

		// DiscordGo Stuff

		message = &discordgo.MessageSend{
			Files: []*discordgo.File{
				{
					Reader: imgDecoded,
					Name:   "ch.png",
				},
			},
			Embed: &discordgo.MessageEmbed{
				Title: "Level " + strconv.Itoa(grimm.Level) + " " + grimm.Name + ".",
				Color: grimm.RarityToColor(),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name: "**Statistics :**",
						Value: "Category : **" + original.Category + "**\n" +
							"XP : " + strconv.FormatInt(grimm.XP, 10) + "/" + strconv.FormatInt(grimm.XPCap, 10) + "\n" +
							"Value : " + fmt.Sprintf("%.2f", grimm.Value) + "%\n" +
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
	}

	ctx.Reply(discord.ReplyParams{
		Content: message,
	})
}
