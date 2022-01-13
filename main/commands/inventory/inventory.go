package commands_inventory

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var InventoryCommand = &discord.Command{
	Name:        "inventory",
	Description: "View informations about a certain persona.",
	Menu:        discord.PersonasMenu,
	Call:        Inventory,
	Aliases: discord.CmdAlias{
		"inv",
	},
	Args: []discord.Arg{
		{
			Name:        "page",
			Description: "The page of your inventory you want to check.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "name",
			Description: "Filter your inventory by name.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "level",
			Description: "Filter your inventory by level.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "rarity",
			Description: "Filter your inventory by rarity.",
			Size:        1,
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
					Name:  "Very_Rare",
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
			Name:        "arms",
			Description: "Filter your inventory by amount of arm.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "value_above",
			Description: "Filter your inventory by value.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "value_below",
			Description: "Filter your inventory by value.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionString,
		},
	},
}

type inventoryMenuData struct {
	UserID string
	Page   int
}

var charPerPage = 10

func Inventory(ctx *discord.CmdContext) {
	var fields []*discordgo.MessageEmbedField
	var page int
	// Filters
	var charName string
	var level int
	var valueAbove float64
	var valueBelow float64
	var buffs int
	var rarity int

	// Arguments
	pageA := ctx.Arguments.GetArg("page", 0, 1)
	charNameA := ctx.Arguments.GetArg("name", 1, "")
	levelA := ctx.Arguments.GetArg("level", 2, 0)
	valueAboveA := ctx.Arguments.GetArg("value_above", 3, 0)
	valueBelowA := ctx.Arguments.GetArg("value_below", 4, 0)
	buffsA := ctx.Arguments.GetArg("buffs", 5, 0)
	rarityA := ctx.Arguments.GetArg("rarity", 6, 0)
	page, _ = strconv.Atoi(fmt.Sprint(pageA.Value))
	charName = fmt.Sprint(charNameA.Value)
	level, _ = strconv.Atoi(fmt.Sprint(levelA.Value))
	valueAbove, _ = strconv.ParseFloat(fmt.Sprint(valueAboveA.Value), 64)
	valueBelow, _ = strconv.ParseFloat(fmt.Sprint(valueBelowA.Value), 64)
	buffs, _ = strconv.Atoi(fmt.Sprint(buffsA.Value))
	rarity, _ = strconv.Atoi(fmt.Sprint(rarityA.Value))

	filtering := charNameA.Found || levelA.Found || valueAboveA.Found || valueBelowA.Found || buffsA.Found || rarityA.Found

	// Useful stuff
	pageMax := int(math.Ceil(float64(len(ctx.Player.Characters)) / 10))
	if page > pageMax {
		page = pageMax
	}

	// Character field
	if len(ctx.Player.Characters) != 0 {
		charsField := &discordgo.MessageEmbedField{
			Name: fmt.Sprintf("Characters (page %d/%d) :", page, pageMax),
		}
		n := 0
		if filtering {
			for _, char := range ctx.Player.Characters {
				if char.CheckConditions(charName, level, valueAbove, valueBelow, rarity, buffs) {
					n++
				}
			}
			pageMax = int(math.Ceil(float64(n) / 10))
			if page > pageMax {
				page = pageMax
			}
		}
		n = 0
		// Filtering
		for i, char := range ctx.Player.Characters {
			if filtering && char.CheckConditions(charName, level, valueAbove, valueBelow, rarity, buffs) {
				continue
			}
			if n%charPerPage != 0 {
				charsField.Value += "\n"
			}
			//Correct interval : [(p-1)*10, p*10]
			if n < (page-1)*10 || n >= page*10 {
				n++
				continue
			}
			n++
			charsField.Value += fmt.Sprintf("`C%d | %s`", i+1, char.FullString())
		}

		fields = append(fields, charsField)
	}

	// Selected character
	selectedField := &discordgo.MessageEmbedField{
		Name:  "**Selected character/grimm**",
		Value: "You did not select any character/grimm yet.",
	}
	if ctx.Player.SelectedType == 0 {
		if ctx.Player.SelectedChar.Name != "" {
			selectedField = &discordgo.MessageEmbedField{
				Name:  "**Selected character**",
				Value: fmt.Sprintf("`%s`", ctx.Player.SelectedChar.FullString()),
			}
		}
	} else {
		if ctx.Player.SelectedGrimm.Name != "" {
			selectedField = &discordgo.MessageEmbedField{
				Name:  "**Selected grimm**",
				Value: fmt.Sprintf("`%s`", ctx.Player.SelectedGrimm.FullString()),
			}
		}
	}
	fields = append(fields, selectedField)

	ctx.Player.FillPlayerMarket()
	if len(ctx.Player.Market.Listings) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  "Listings :",
			Value: fmt.Sprintf("You have %d listings on the market.\nType `%slistings` for details.", len(ctx.Player.Market.Listings), ctx.Guild.Prefix),
		})
	}

	if len(ctx.Player.Market.Auctions) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  "Auctions :",
			Value: fmt.Sprintf("You have %d auctions on the market.\nType `%sauctions` for details.", len(ctx.Player.Market.Auctions), ctx.Guild.Prefix),
		})
	}

	reply := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s's characters", ctx.Author.Username),
		Description: fmt.Sprintf("To select a character, please type `%sselect <PersonaID>`.", ctx.Guild.Prefix),
		Color:       config.Botcolor,
		Fields:      fields,
		Footer:      discord.DefaultFooter,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Author.AvatarURL("512"),
		},
	}

	ctx.Reply(discord.ReplyParams{
		Content: reply,
	})
}
