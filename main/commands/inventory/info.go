package commands_inventory

import (
	"bytes"
	"fmt"
	"rwby-adventures/main/discord"
	"rwby-adventures/main/static"
	"rwby-adventures/models"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var InfoCommand = &discord.Command{
	Name:        "info",
	Description: "View informations about a certain persona.",
	Menu:        discord.InventoryMenu,
	Call:        Info,
	Args: []discord.Arg{
		{
			Name:        "id",
			Description: "Identification number of your persona.",
			Size:        1,
			Required:    false,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "latest",
			Description: "Whether or not you want to view the infos of your latest persona.",
			Size:        1,
			Required:    false,
			Type:        discordgo.ApplicationCommandOptionBoolean,
		},
	},
}

type InfoMenuData struct {
	UserID  string
	Char    *models.Character
	Grimm   *models.Grimm
	IsGrimm bool
}

func Info(ctx *discord.CmdContext) {
	var char *models.Character
	var grimm *models.Grimm
	var pickLatest bool
	latest := ctx.Arguments.GetArg("latest", 1, false)
	if v, ok := latest.Value.(bool); ok && v {
		pickLatest = true
	}
	arg := ctx.Arguments.GetArg("id", 0, "")
	isGrimm, index, err := arg.CharGrimmParse()
	if !pickLatest {
		if err != nil {
			if ctx.Player.SelectedChar == nil && ctx.Player.SelectedGrimm == nil {
				ctx.Reply(discord.ReplyParams{
					Content:   "You have not selected any persona.",
					Ephemeral: true,
				})
				return
			}
			char = ctx.Player.SelectedChar
			grimm = ctx.Player.SelectedGrimm
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
				grimm = ctx.Player.Grimms[index-1]
			} else {
				if index > len(ctx.Player.Characters) {
					ctx.Reply(discord.ReplyParams{
						Content:   "You don't have any character with this number.",
						Ephemeral: true,
					})
					return
				}
				char = ctx.Player.Characters[index-1]
			}
		}
	}

	if pickLatest {
		isGrimm, char, grimm, index, err = ctx.Player.GetLatestPersona()
		index += 1
		if err != nil {
			ctx.Reply(discord.ReplyParams{
				Content:   "You don't have any persona.",
				Ephemeral: true,
			})
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
	imgData, _ := static.DatabaseFS.ReadFile(original.ImageFile)
	imgDecoded := bytes.NewBuffer(imgData)

	// DiscordGo Stuff

	complex := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Reader: imgDecoded,
				Name:   "ch.png",
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title: "Level " + strconv.Itoa(char.Level) + " " + char.Name + ". (ID: C" + strconv.Itoa(number) + ")",
			Color: char.RarityToColor(),
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
		Components: InfoComponent(ctx.ID),
	}

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          MenuInfo,
		Data: &InfoMenuData{
			UserID:  ctx.Author.ID,
			Char:    char,
			IsGrimm: false,
		},
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content: complex,
	})
}

func grimmInfo(ctx *discord.CmdContext, grimm *models.Grimm, number int) {
	// Useful data
	original := grimm.ToRealGrimm()
	imgData, _ := static.DatabaseFS.ReadFile(original.ImageFile)
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
			Title: "Level " + strconv.Itoa(grimm.Level) + " " + grimm.Name + ". (ID: G" + strconv.Itoa(number) + ")",
			Color: grimm.RarityToColor(),
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
		Components: InfoComponent(ctx.ID),
	}

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          MenuInfo,
		Data: &InfoMenuData{
			UserID:  ctx.Author.ID,
			Grimm:   grimm,
			IsGrimm: true,
		},
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content: Complex,
	})
}

func MenuInfo(ctx *discord.CmdContext) {
	// Reply to the interaction so it is seamless for the player
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if ctx.Author.ID != ctx.Menu.SourceContext.Author.ID {
		return
	}
	d := ctx.Menu.Data.(*InfoMenuData)
	split := strings.Split(ctx.ComponentData.CustomID, "-")
	switch split[1] {
	case "remove":
		remove(ctx, &removeData{
			Char:     d.Char,
			Grimm:    d.Grimm,
			isGrimm:  d.IsGrimm,
			FollowUp: true,
		})
	case "pick":
		Select(ctx, &SelectData{
			Char:     d.Char,
			Grimm:    d.Grimm,
			isGrimm:  d.IsGrimm,
			FollowUp: true,
		})
	case "addfav":
		AddFavorite(ctx, &AddFavoriteData{
			Char:     d.Char,
			Grimm:    d.Grimm,
			isGrimm:  d.IsGrimm,
			FollowUp: true,
		})
	case "remfav":
		RemoveFavorite(ctx, &RemFavoriteData{
			Char:     d.Char,
			Grimm:    d.Grimm,
			isGrimm:  d.IsGrimm,
			FollowUp: true,
		})
	default:
		return
	}
}

func InfoComponent(menuID string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label: "Remove",
					Emoji: discordgo.ComponentEmoji{
						Name: "❎",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-remove",
				},
				&discordgo.Button{
					Label: "Add Favorite",
					Emoji: discordgo.ComponentEmoji{
						Name: "⭐",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-addfav",
				},
				&discordgo.Button{
					Label: "Remove Favorite",
					Emoji: discordgo.ComponentEmoji{
						Name: "💥",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-remfav",
				},
				&discordgo.Button{
					Label: "Pick",
					Emoji: discordgo.ComponentEmoji{
						Name: "⛏️",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-pick",
				},
			},
		},
	}
}
