package commands_market

import (
	"fmt"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strings"

	"github.com/bwmarrin/discordgo"
	uuid "github.com/satori/go.uuid"
)

type ListingMenuData struct {
	UserID  string
	Listing *models.Listing
	isGrimm bool
}

func CreateListing(ctx *discord.CmdContext) {
	var char *models.Character
	var grimm *models.Grimm
	arg := ctx.Arguments.GetArg("id", 0, "")
	price := ctx.Arguments.GetArg("price", 1, 0)
	note := ctx.Arguments.GetArg("note", 2, "")
	isGrimm, index, err := arg.CharGrimmParse()
	// No ID => selected character
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
	} else { // We search for ID
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

	data := &ListingMenuData{
		UserID: ctx.Player.DiscordID,
		Listing: &models.Listing{
			ID:         uuid.NewV4().String(),
			SellerID:   ctx.Player.DiscordID,
			SellerName: ctx.Author.Username,
			Price:      int(price.Value.(float64)),
			Note:       fmt.Sprint(note.Value),

			Grimm: grimm,
			Char:  char,
		},
		isGrimm: isGrimm,
	}

	var personaStr string

	if isGrimm {
		data.Listing.Type = models.GrimmType
		personaStr = grimm.FullString()
	} else {
		data.Listing.Type = models.CharType
		personaStr = char.FullString()
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

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          createListingMenu,
		Data:          data,
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       "Create listing confirmation",
			Description: fmt.Sprintf("You are creating a listing for `%s` at **%d** Liens?\n\nTo confirm this operation, please click on the following emoji : %s", personaStr, data.Listing.Price, e[r].Name),
		},
		Components: createListingComponent(ctx.ID, e, e[r]),
	})
}

func createListingMenu(ctx *discord.CmdContext) {
	// Reply to the interaction so it is seamless for the player
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if ctx.Author.ID != ctx.Menu.SourceContext.Author.ID {
		return
	}
	d := ctx.Menu.Data.(*ListingMenuData)
	split := strings.Split(ctx.ComponentData.CustomID, "-")
	switch split[1] {
	case "correct":
		var reply string
		var err error
		if d.isGrimm {
			if ctx.Player.GrimmAmount() <= 1 {
				ctx.Reply(discord.ReplyParams{
					Content:  "You do not have enough grimms to market one right now.",
					FollowUp: true,
				})
				return
			}
			d.Listing.Grimm, err = models.GetGrimm(d.Listing.Grimm.GrimmID)
			if err != nil {
				ctx.Reply(discord.ReplyParams{
					Content:  "There has been an error retrieving the grimm.",
					FollowUp: true,
				})
				return
			}
			if d.Listing.Grimm.InHunt || d.Listing.Grimm.UserID != ctx.Player.DiscordID {
				ctx.Reply(discord.ReplyParams{
					Content:  "You cannot market that grimm.",
					FollowUp: true,
				})
				return
			}
			reply = fmt.Sprintf("You listed : %s for %d Liens.", d.Listing.Grimm.FullString(), d.Listing.Price)

			models.CreateListing(d.Listing)

			if ctx.Player.SelectedID == d.Listing.Grimm.GrimmID {
				for _, g := range ctx.Player.Grimms {
					if g.GrimmID != d.Listing.Grimm.GrimmID {
						ctx.Player.SelectedID = g.GrimmID
						config.Database.Save(ctx.Player)
						break
					}
				}
			}
		} else {
			if ctx.Player.CharAmount() <= 1 {
				ctx.Reply(discord.ReplyParams{
					Content:  "You do not have enough characters to market one right now.",
					FollowUp: true,
				})
				return
			}
			d.Listing.Char, err = models.GetChar(d.Listing.Char.CharID)
			if err != nil {
				ctx.Reply(discord.ReplyParams{
					Content:  "There has been an error retrieving the character.",
					FollowUp: true,
				})
				return
			}
			if d.Listing.Char.InMission || d.Listing.Char.UserID != ctx.Player.DiscordID {
				ctx.Reply(discord.ReplyParams{
					Content:  "You cannot market that character.",
					FollowUp: true,
				})
				return
			}
			reply = fmt.Sprintf("You listed : `%s` for **%d** Liens.", d.Listing.Char.FullString(), d.Listing.Price)

			models.CreateListing(d.Listing)

			if ctx.Player.SelectedID == d.Listing.Char.CharID {
				for _, c := range ctx.Player.Characters {
					if c.CharID != d.Listing.Char.CharID {
						ctx.Player.SelectedID = c.CharID
						config.Database.Save(ctx.Player)
						break
					}
				}
			}
		}
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

func createListingComponent(menuID string, emojis []*discordgo.ComponentEmoji, correct *discordgo.ComponentEmoji) []discordgo.MessageComponent {
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
