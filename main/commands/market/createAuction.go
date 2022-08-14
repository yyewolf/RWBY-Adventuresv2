package commands_market

import (
	"fmt"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/main/market"
	"rwby-adventures/microservices"
	"rwby-adventures/models"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	uuid "github.com/satori/go.uuid"
)

type AuctionMenuData struct {
	UserID  string
	Auction *models.Auction
	isGrimm bool
}

func CreateAuction(ctx *discord.CmdContext) {
	var char *models.Character
	var grimm *models.Grimm
	arg := ctx.Arguments.GetArg("id", 0, "")
	duration := ctx.Arguments.GetArg("duration", 1, "")
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

	d := duration.Value.(float64)

	if d < 4 {
		ctx.Reply(discord.ReplyParams{
			Content:   "Auction duration must be at least 4 hours.",
			Ephemeral: true,
		})
		return
	}

	data := &AuctionMenuData{
		UserID: ctx.Player.DiscordID,
		Auction: &models.Auction{
			ID:         uuid.NewV4().String(),
			SellerID:   ctx.Player.DiscordID,
			SellerName: ctx.Author.Username,
			StartedAt:  time.Now().Unix(),
			EndsAt:     time.Now().Add(time.Hour * time.Duration(d)).Unix(),

			Grimm: grimm,
			Char:  char,
		},
		isGrimm: isGrimm,
	}

	var personaStr string

	if isGrimm {
		data.Auction.Type = models.GrimmType
		personaStr = grimm.FullString()
	} else {
		data.Auction.Type = models.CharType
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
		Call:          createAuctionMenu,
		Data:          data,
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       "Create auction confirmation",
			Color:       config.Botcolor,
			Description: fmt.Sprintf("You are creating an auction for `%s` for **%.0f** hours?\n\nTo confirm this operation, please click on the following emoji : %s", personaStr, d, e[r].Name),
		},
		Components: createAuctionComponent(ctx.ID, e, e[r]),
	})
}

func createAuctionMenu(ctx *discord.CmdContext) {
	// Reply to the interaction so it is seamless for the player
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if ctx.Author.ID != ctx.Menu.SourceContext.Author.ID {
		return
	}
	d := ctx.Menu.Data.(*AuctionMenuData)
	split := strings.Split(ctx.ComponentData.CustomID, "-")
	switch split[1] {
	case "correct":
		var reply string
		var err error
		if d.isGrimm {
			if ctx.Player.GrimmAmount() <= 1 {
				ctx.Reply(discord.ReplyParams{
					Content:  "You do not have enough grimms to auction one right now.",
					FollowUp: true,
				})
				return
			}
			d.Auction.Grimm, err = models.GetGrimm(d.Auction.Grimm.GrimmID)
			if err != nil {
				ctx.Reply(discord.ReplyParams{
					Content:  "There has been an error retrieving the grimm.",
					FollowUp: true,
				})
				return
			}
			if d.Auction.Grimm.InHunt || d.Auction.Grimm.UserID != ctx.Player.DiscordID {
				ctx.Reply(discord.ReplyParams{
					Content:  "You cannot auction that grimm.",
					FollowUp: true,
				})
				return
			}
			reply = fmt.Sprintf("You auctioned : `%s`.", d.Auction.Grimm.FullString())

			models.CreateAuction(d.Auction)

			if ctx.Player.SelectedID == d.Auction.Grimm.GrimmID {
				for _, g := range ctx.Player.Grimms {
					if g.GrimmID != d.Auction.Grimm.GrimmID {
						ctx.Player.SelectedID = g.GrimmID
						config.Database.Save(ctx.Player)
						break
					}
				}
			}
		} else {
			if ctx.Player.CharAmount() <= 1 {
				ctx.Reply(discord.ReplyParams{
					Content:  "You do not have enough characters to auction one right now.",
					FollowUp: true,
				})
				return
			}
			d.Auction.Char, err = models.GetChar(d.Auction.Char.CharID)
			if err != nil {
				ctx.Reply(discord.ReplyParams{
					Content:  "There has been an error retrieving the character.",
					FollowUp: true,
				})
				return
			}
			if d.Auction.Char.InMission || d.Auction.Char.UserID != ctx.Player.DiscordID {
				ctx.Reply(discord.ReplyParams{
					Content:  "You cannot auction that character.",
					FollowUp: true,
				})
				return
			}
			reply = fmt.Sprintf("You auctioned : `%s`.", d.Auction.Char.FullString())

			models.CreateAuction(d.Auction)

			if ctx.Player.SelectedID == d.Auction.Char.CharID {
				for _, c := range ctx.Player.Characters {
					if c.CharID != d.Auction.Char.CharID {
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

		market.UpdateAuctions(&microservices.MarketCreate{
			ID: d.Auction.ID,
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

func createAuctionComponent(menuID string, emojis []*discordgo.ComponentEmoji, correct *discordgo.ComponentEmoji) []discordgo.MessageComponent {
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
