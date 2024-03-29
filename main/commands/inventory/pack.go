package commands_inventory

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	market_commands "rwby-adventures/main/commands/market"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var PackCommand = &discord.Command{
	Name:        "pack",
	Description: "Check out your pack (grimm inventory).",
	Menu:        discord.InventoryMenu,
	Call:        Pack,
	Args: []discord.Arg{
		{
			Name:        "page",
			Description: "The page of your pack you want to check.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "name",
			Description: "Filter your pack by name.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "level",
			Description: "Filter your pack by level.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "rarity",
			Description: "Filter your pack by rarity.",
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
			Description: "Filter your pack by amount of arm.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "value_above",
			Description: "Filter your pack by value.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "value_below",
			Description: "Filter your pack by value.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "favorites",
			Description: "Filter your pack by favorites.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionBoolean,
		},
	},
}

func Pack(ctx *discord.CmdContext) {
	var fields []*discordgo.MessageEmbedField
	var page int
	// Filters
	var grimmName string
	var level int
	var valueAbove float64
	var valueBelow float64
	var buffs int
	var rarity int
	var favorites bool

	// Arguments
	pageA := ctx.Arguments.GetArg("page", 0, 1)
	charNameA := ctx.Arguments.GetArg("name", 1, "")
	levelA := ctx.Arguments.GetArg("level", 2, -1)
	valueAboveA := ctx.Arguments.GetArg("value_above", 3, -1)
	valueBelowA := ctx.Arguments.GetArg("value_below", 4, -1)
	buffsA := ctx.Arguments.GetArg("buffs", 5, -1)
	rarityA := ctx.Arguments.GetArg("rarity", 6, -1)
	favoritesA := ctx.Arguments.GetArg("favorites", 7, false)
	page, _ = strconv.Atoi(fmt.Sprint(pageA.Value))
	grimmName = fmt.Sprint(charNameA.Value)
	level, _ = strconv.Atoi(fmt.Sprint(levelA.Value))
	valueAbove, _ = strconv.ParseFloat(fmt.Sprint(valueAboveA.Value), 64)
	valueBelow, _ = strconv.ParseFloat(fmt.Sprint(valueBelowA.Value), 64)
	buffs, _ = strconv.Atoi(fmt.Sprint(buffsA.Value))
	rarity, _ = strconv.Atoi(fmt.Sprint(rarityA.Value))
	favorites, _ = favoritesA.Value.(bool)

	filtering := charNameA.Found || levelA.Found || valueAboveA.Found || valueBelowA.Found || buffsA.Found || rarityA.Found

	filters := &models.InvFilters{
		Name:      grimmName,
		Level:     level,
		ValAbove:  valueAbove,
		ValBelow:  valueBelow,
		Buffs:     buffs,
		Rarity:    rarity,
		Favorites: favorites,
		Filtering: filtering,
	}

	pageMax := int(math.Ceil(float64(len(ctx.Player.Grimms)) / 10))

	if ctx.IsComponent {
		d := ctx.Menu.Data.(*inventoryMenuData)

		if d.Page <= 0 {
			d.Page = pageMax
		}
		if d.Page > pageMax {
			d.Page = 1
		}

		filters = d.Filters
		page = d.Page
	}

	// Useful stuff
	if page <= 0 {
		page = pageMax
	}
	if page > pageMax {
		page = 1
	}

	// Character field
	if len(ctx.Player.Grimms) != 0 {
		n := 0
		for _, grimm := range ctx.Player.Grimms {
			if grimm.CheckConditions(filters) {
				n++
			}
		}
		pageMax = int(math.Ceil(float64(n) / 10))
		if page > pageMax {
			page = pageMax
		}
		n = 0
		grimmsField := &discordgo.MessageEmbedField{
			Name: fmt.Sprintf("Grimms (page %d/%d) :", page, pageMax),
		}
		// Filtering
		for i, grimm := range ctx.Player.Grimms {
			if !grimm.CheckConditions(filters) {
				continue
			}
			if n%charPerPage != 0 {
				grimmsField.Value += "\n"
			}
			//Correct interval : [(p-1)*10, p*10]
			if n < (page-1)*10 || n >= page*10 {
				n++
				continue
			}
			n++
			grimmsField.Value += fmt.Sprintf("`G%d | %s`", i+1, grimm.FullString())
		}

		if n > 0 {
			fields = append(fields, grimmsField)
		} else {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  "Grimms :",
				Value: "You have no grimms to be shown (maybe change your filters?)",
			})
		}
	} else {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  "Grimms :",
			Value: "You have no grimms to be shown.",
		})
	}

	// Selected character
	selectedField := &discordgo.MessageEmbedField{
		Name:  "**Selected character/grimm**",
		Value: "You did not select any character/grimm yet.",
	}
	if ctx.Player.SelectedType == 0 {
		if ctx.Player.SelectedChar != nil {
			selectedField = &discordgo.MessageEmbedField{
				Name:  "**Selected character**",
				Value: fmt.Sprintf("`%s`", ctx.Player.SelectedChar.FullString()),
			}
		}
	} else {
		if ctx.Player.SelectedGrimm != nil {
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
			Value: fmt.Sprintf("You have %d listings on the market.\nUse </market listings view:%s> for details.", len(ctx.Player.Market.Listings), market_commands.MarketCommand.ID),
		})
	}

	if len(ctx.Player.Market.Auctions) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  "Auctions :",
			Value: fmt.Sprintf("You have %d auctions on the market.\nType </market auctions view:%s> for details.", len(ctx.Player.Market.Auctions), market_commands.MarketCommand.ID),
		})
	}

	menuID := ctx.ID
	editID := ""
	if ctx.IsComponent {
		// Keep old context if a button is pressed
		menuID = ctx.Menu.MenuID
		editID = ctx.Message.ID
		//ctx.Menu.SourceContext.ID = ctx.Message.ID
		//ctx = ctx.Menu.SourceContext
	}

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        menuID,
		SourceContext: ctx,
		Call:          PackPages,
		Data: &inventoryMenuData{
			UserID:  ctx.Author.ID,
			Page:    page,
			Filters: filters,
		},
	}, 0)

	reply := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s's Grimms", ctx.Author.Username),
			Description: fmt.Sprintf("To select a character, please use %s.", SelectCommand.Mention()),
			Color:       config.Botcolor,
			Fields:      fields,
			Footer:      discord.DefaultFooter,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL("512"),
			},
		},
		Components: inventoryComponent(menuID),
	}

	ctx.Reply(discord.ReplyParams{
		Content: reply,
		ID:      editID,
		Edit:    ctx.IsComponent,
	})
}

func PackPages(ctx *discord.CmdContext) {
	d := ctx.Menu.Data.(*inventoryMenuData)

	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	if ctx.Author.ID != d.UserID {
		return
	}

	switch strings.Split(ctx.ComponentData.CustomID, "-")[1] {
	case "prev":
		d.Page++
	case "refresh":
		break
	case "next":
		d.Page--
	default:
		return
	}

	Pack(ctx)
}
