package commands_inventory

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var InventoryCommand = &discord.Command{
	Name:        "inventory",
	Description: "Check out your inventory.",
	Menu:        discord.InventoryMenu,
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
		{
			Name:        "favorites",
			Description: "Filter your inventory by favorites.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionBoolean,
		},
	},
}

type inventoryMenuData struct {
	UserID  string
	Page    int
	Filters *models.InvFilters
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
	var favorites bool

	// Arguments
	pageA := ctx.Arguments.GetArg("page", 0, 1)
	charNameA := ctx.Arguments.GetArg("name", 1, "")
	levelA := ctx.Arguments.GetArg("level", 2, 0)
	valueAboveA := ctx.Arguments.GetArg("value_above", 3, 0)
	valueBelowA := ctx.Arguments.GetArg("value_below", 4, 0)
	buffsA := ctx.Arguments.GetArg("buffs", 5, 0)
	rarityA := ctx.Arguments.GetArg("rarity", 6, 0)
	favoritesA := ctx.Arguments.GetArg("favorites", 7, false)
	page, _ = strconv.Atoi(fmt.Sprint(pageA.Value))
	charName = fmt.Sprint(charNameA.Value)
	level, _ = strconv.Atoi(fmt.Sprint(levelA.Value))
	valueAbove, _ = strconv.ParseFloat(fmt.Sprint(valueAboveA.Value), 64)
	valueBelow, _ = strconv.ParseFloat(fmt.Sprint(valueBelowA.Value), 64)
	buffs, _ = strconv.Atoi(fmt.Sprint(buffsA.Value))
	rarity, _ = strconv.Atoi(fmt.Sprint(rarityA.Value))
	favorites, _ = favoritesA.Value.(bool)

	filtering := charNameA.Found || levelA.Found || valueAboveA.Found || valueBelowA.Found || buffsA.Found || rarityA.Found

	filters := &models.InvFilters{
		Name:      charName,
		Level:     level,
		ValAbove:  valueAbove,
		ValBelow:  valueBelow,
		Buffs:     buffs,
		Rarity:    rarity,
		Favorites: favorites,
		Filtering: filtering,
	}

	pageMax := int(math.Ceil(float64(len(ctx.Player.Characters)) / 10))

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
	if len(ctx.Player.Characters) != 0 {
		n := 0
		for _, char := range ctx.Player.Characters {
			if char.CheckConditions(filters) {
				n++
			}
		}
		pageMax = int(math.Ceil(float64(n) / 10))
		if page > pageMax {
			page = pageMax
		}
		n = 0
		charsField := &discordgo.MessageEmbedField{
			Name: fmt.Sprintf("Characters (page %d/%d) :", page, pageMax),
		}
		// Filtering
		for i, char := range ctx.Player.Characters {
			if !char.CheckConditions(filters) {
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

		if n > 0 {
			fields = append(fields, charsField)
		} else {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  "Characters :",
				Value: "You have no characters to be shown (maybe change your filters?)",
			})
		}
	} else {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  "Characters :",
			Value: "You have no characters to be shown.",
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
			Value: fmt.Sprintf("You have %d listings on the market.\nType `%slistings` for details.", len(ctx.Player.Market.Listings), ctx.Guild.Prefix),
		})
	}

	if len(ctx.Player.Market.Auctions) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  "Auctions :",
			Value: fmt.Sprintf("You have %d auctions on the market.\nType `%sauctions` for details.", len(ctx.Player.Market.Auctions), ctx.Guild.Prefix),
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
		Call:          InventoryPages,
		Data: &inventoryMenuData{
			UserID:  ctx.Author.ID,
			Page:    page,
			Filters: filters,
		},
	}, 0)

	reply := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s's characters", ctx.Author.Username),
			Description: fmt.Sprintf("To select a character, please type `%sselect <PersonaID>`.", ctx.Guild.Prefix),
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

func InventoryPages(ctx *discord.CmdContext) {
	d := ctx.Menu.Data.(*inventoryMenuData)

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

	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	Inventory(ctx)
}

func inventoryComponent(menuID string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label: "Prev",
					Emoji: discordgo.ComponentEmoji{
						Name: "⬅️",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-prev",
				},
				&discordgo.Button{
					Label: "Refresh",
					Emoji: discordgo.ComponentEmoji{
						Name: "🔄",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-refresh",
				},
				&discordgo.Button{
					Label: "Next",
					Emoji: discordgo.ComponentEmoji{
						Name: "➡️",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-next",
				},
			},
		},
	}
}
