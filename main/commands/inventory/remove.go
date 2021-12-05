package commands_inventory

import (
	"fmt"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var RemoveCommand = &discord.Command{
	Name:        "remove",
	Description: "Remove a persona from your bags.",
	Menu:        discord.PersonasMenu,
	Call:        RemovePersona,
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

type removeData struct {
	Char     *models.Character
	Grimm    *models.Grimm
	isGrimm  bool
	FollowUp bool
}

func RemovePersona(ctx *discord.CmdContext) {
	// We parse user input
	var err error
	var pickLatest bool
	var isGrimm bool
	var index int
	var grimm *models.Grimm
	var char *models.Character
	var id *discord.CommandArg

	latest, err := ctx.Arguments.GetArg("latest", 1)
	if err == nil {
		if v, ok := latest.Value.(bool); ok && v {
			pickLatest = v
			goto skip
		}
	}

	id, err = ctx.Arguments.GetArg("id", 0)
	if err != nil && !pickLatest {
		ctx.Reply(discord.ReplyParams{
			Content:   "You need to input at least the ID of the persona you wish to remove.",
			Ephemeral: true,
		})
		return
	}
	isGrimm, index, err = id.CharGrimmParse()
	if err != nil && !pickLatest {
		ctx.Reply(discord.ReplyParams{
			Content:   "I did not understand the ID that you sent.",
			Ephemeral: true,
		})
		return
	}

	// We search for the character they sent

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

skip:
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
	// We proceed to remove the character
	// We don't preload the confirmation because the fonction can also be called from other places

	remove(ctx, &removeData{
		Grimm:   grimm,
		Char:    char,
		isGrimm: isGrimm,
	})
}

func remove(ctx *discord.CmdContext, rem *removeData) {
	var reply string
	var additionalData []string

	if rem.isGrimm {
		if rem.Char.IsInFavorites {
			additionalData = append(additionalData, "\nThis grimm is in your **favorites**!")
		}
		if rem.Char.ToRealChar().Limited {
			additionalData = append(additionalData, "\nThis grimm is a **limited edition**, make sure that you really want to delete it!")
		}
		reply = fmt.Sprintf("You are trying to remove :\n%s", rem.Grimm.FullString())
		for _, str := range additionalData {
			reply += str
		}
	} else {
		if rem.Char.IsInFavorites {
			additionalData = append(additionalData, "\nThis character is in your **favorites**!")
		}
		if rem.Char.ToRealChar().Limited {
			additionalData = append(additionalData, "\nThis character is a **limited edition**, make sure that you really want to delete it!")
		}
		reply = fmt.Sprintf("You are trying to remove :\n%s", rem.Char.FullString())
		for _, str := range additionalData {
			reply += str
		}
	}

	var e = []*discordgo.ComponentEmoji{
		{
			Name: "🐶",
		},
		{
			Name: "🐱",
		},
		{
			Name: "🐭",
		},
		{
			Name: "🐼",
		},
	}

	rand.Shuffle(len(e), func(i, j int) {
		e[i], e[j] = e[j], e[i]
	})

	r := rand.Intn(len(e))

	reply += fmt.Sprintf("\n\nTo confirm this operation, please click on the following emoji : %s", e[r].Name)

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          removeMenu,
		Data:          rem,
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content:    reply,
		Components: removeComponent(ctx.ID, e, e[r]),
		FollowUp:   rem.FollowUp,
	})
}

func removeMenu(ctx *discord.CmdContext) {
	// Reply to the interaction so it is seamless for the player
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if ctx.Author.ID != ctx.Menu.SourceContext.Author.ID {
		return
	}
	d := ctx.Menu.Data.(*removeData)
	split := strings.Split(ctx.ComponentData.CustomID, "-")
	switch split[1] {
	case "correct":
		var reply string
		var xpjar int64
		if d.isGrimm {
			if ctx.Player.GrimmAmount() <= 1 {
				ctx.Reply(discord.ReplyParams{
					Content:  "You do not have enough grimms to remove one right now.",
					FollowUp: true,
				})
				return
			}
			reply = fmt.Sprintf("You removed :\n%s", d.Grimm.FullString())
			xpjar = int64(float64(d.Grimm.Level) * float64(d.Grimm.XP) * float64(rand.Intn(4)+3) / 100.0)

			config.Database.Select("Stats").Where("user_id=?", ctx.Author.ID).Delete(d.Grimm)

			if ctx.Player.SelectedID == d.Grimm.GrimmID {
				for _, g := range ctx.Player.Grimms {
					if g.GrimmID != d.Grimm.GrimmID {
						ctx.Player.SelectedID = g.GrimmID
						config.Database.Save(ctx.Player)
						break
					}
				}
			}
		} else {
			if ctx.Player.CharAmount() <= 1 {
				ctx.Reply(discord.ReplyParams{
					Content:  "You do not have enough characters to remove one right now.",
					FollowUp: true,
				})
				return
			}
			reply = fmt.Sprintf("You removed :\n%s", d.Char.FullString())
			xpjar = int64(float64(d.Char.Level) * float64(d.Char.XP) * float64(rand.Intn(4)+3) / 100.0)

			config.Database.Select("Stats").Where("user_id=?", ctx.Author.ID).Delete(d.Char)

			if ctx.Player.SelectedID == d.Char.CharID {
				for _, c := range ctx.Player.Characters {
					if c.CharID != d.Char.CharID {
						ctx.Player.SelectedID = c.CharID
						config.Database.Save(ctx.Player)
						break
					}
				}
			}
		}
		reply += fmt.Sprintf("\nYou also earned **%dXP** that have been added to your **XP Jar** (`/jar`) !", xpjar)
		ctx.Reply(discord.ReplyParams{
			Content:   reply,
			FollowUp:  true,
			Ephemeral: true,
		})
	case "notcorrect":
		ctx.Reply(discord.ReplyParams{
			Content:   "You did not click the correct emoji.",
			FollowUp:  true,
			Ephemeral: true,
		})
	default:
		return
	}
}

func removeComponent(menuID string, emojis []*discordgo.ComponentEmoji, correct *discordgo.ComponentEmoji) []discordgo.MessageComponent {
	var c []discordgo.MessageComponent
	for i, e := range emojis {
		action := fmt.Sprintf("notcorrect-%d", i)
		if e.Name == correct.Name {
			action = fmt.Sprintf("correct-%d", i)
		}
		c = append(c, discordgo.Button{
			Emoji:    *e,
			CustomID: fmt.Sprintf("%s-%s", menuID, action),
		})
	}
	return []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: c,
		},
	}

}
