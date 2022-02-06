package commands_boxes

import (
	"fmt"
	"math"
	"math/rand"
	chars "rwby-adventures/characters"
	"rwby-adventures/config"
	"rwby-adventures/grimms"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	uuid "github.com/satori/go.uuid"
)

var BoxPerPage = 5

type specialBoxesMenuData struct {
	UserID string
	Page   int
}

func GiveChar(ctx *discord.CmdContext, b *BoxFilter, Loot *chars.CharacterStruct) (success bool) {
	var isLucky bool
	//Luck Boosts
	if ctx.Player.Shop.LuckBoost {
		isLucky = true
		ctx.Player.Shop.LuckBoostTime--
		config.Database.Save(ctx.Player.Shop)
		ctx.Player.SendLuckNotice(ctx.Session)
	}

	// Value finder
	Value := rand.NormFloat64()*b.ValStd + b.ValMean
	for Value < 0 || Value > 100 {
		Value = rand.NormFloat64()*b.ValStd + b.ValMean
	}
	if isLucky {
		Value += 2.5
	}

	// Rarity finder
	var Rarity int
	rdm := rand.ExpFloat64() / b.RarityRate
	for i := 0; float64(i) < rdm; i++ {
		if float64(i+1) > rdm {
			Rarity = i
			return
		}
	}

	// Level finder
	Rate := 0.4
	if isLucky {
		Rate = 0.3
	}
	Level := int(rand.ExpFloat64() / Rate)
	for Level > 40 {
		Level = int(rand.ExpFloat64() / Rate)
	}

	ID := uuid.NewV4().String()
	Char := &models.Character{
		Name:   Loot.Name,
		CharID: ID,
		UserID: ctx.Author.ID,
		Level:  Level,
		XP:     0,
		Stats:  models.CharacterStats(Loot.Stats),
		Rarity: Rarity,
	}
	Char.XPCap = Char.CalcXPCap()
	config.Database.Create(Char)

	Cplx := Char.ToLootedEmbed(ctx.Author.Mention(), ctx.ID, b.Box, Loot)
	Cplx.Embed.Footer = discord.DefaultFooter

	ctx.Reply(discord.ReplyParams{
		Content: Cplx,
	})
	return true
}

func GiveGrimm(ctx *discord.CmdContext, b *BoxFilter, Loot *grimms.GrimmStruct) (success bool) {
	var isLucky bool
	//Luck Boosts
	if ctx.Player.Shop.LuckBoost {
		isLucky = true
		ctx.Player.Shop.LuckBoostTime--
		config.Database.Save(ctx.Player.Shop)
		ctx.Player.SendLuckNotice(ctx.Session)
	}

	// Value finder
	Value := rand.NormFloat64()*b.ValStd + b.ValMean
	for Value < 0 || Value > 100 {
		Value = rand.NormFloat64()*b.ValStd + b.ValMean
	}
	if isLucky {
		Value += 2.5
	}

	// Rarity finder
	var Rarity int
	rdm := rand.ExpFloat64() / b.RarityRate
	for i := 0; float64(i) < rdm; i++ {
		if float64(i+1) > rdm {
			Rarity = i
			return
		}
	}

	// Level finder
	Rate := 0.4
	if isLucky {
		Rate = 0.3
	}
	Level := int(rand.ExpFloat64() / Rate)
	for Level > 40 {
		Level = int(rand.ExpFloat64() / Rate)
	}

	ID := uuid.NewV4().String()
	Grimm := &models.Grimm{
		Name:    Loot.Name,
		GrimmID: ID,
		UserID:  ctx.Author.ID,
		Level:   Level,
		XP:      0,
		Stats:   models.GrimmStat(Loot.Stats),
		Rarity:  Rarity,
	}
	Grimm.XPCap = Grimm.CalcXPCap()
	config.Database.Create(Grimm)

	Cplx := Grimm.ToLootedEmbed(ctx.Author.Mention(), ctx.ID, b.Box, Loot)
	Cplx.Embed.Footer = discord.DefaultFooter

	ctx.Reply(discord.ReplyParams{
		Content: Cplx,
	})
	return true
}

func OpenSpecial(ctx *discord.CmdContext, index int) {
	if len(ctx.Player.SpecialBoxes) <= index {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have this **Special Box**.",
			Ephemeral: true,
		})
		return
	}
	b := &BoxFilter{
		IncludeLimited: true,
		Box:            "Special Box",
		ValStd:         15,
		ValMean:        58,
		RarityRate:     1.125,
	}

	Box := ctx.Player.SpecialBoxes[index]

	L1 := b.FilterChars()
	L2 := b.FilterGrimms()
	n := 0
reroll:
	Char := L1[rand.Intn(len(L1))]
	Grimm := L2[rand.Intn(len(L2))]
	if n < 3 {
		if Char.Name != Box.For || Grimm.Name != Box.For {
			n++
			goto reroll
		}
	}

	if Char.Name == Box.For {
		if GiveChar(ctx, b, &Char) {
			config.Database.Delete(ctx.Player.SpecialBoxes, "for=?", Box.For)
		}
	} else if Grimm.Name == Box.For {
		if GiveGrimm(ctx, b, &Grimm) {
			config.Database.Delete(ctx.Player.SpecialBoxes, "for=?", Box.For)
		}
	} else {
		rdm := rand.Intn(2)
		if rdm >= 1 {
			if GiveChar(ctx, b, &Char) {
				config.Database.Delete(ctx.Player.SpecialBoxes, "for=?", Box.For)
			}
		} else {
			if GiveGrimm(ctx, b, &Grimm) {
				config.Database.Delete(ctx.Player.SpecialBoxes, "for=?", Box.For)
			}
		}
	}
}

func OpenSpecialCmd(ctx *discord.CmdContext) {
	if len(ctx.Player.SpecialBoxes) <= 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You do not have enough **Special Boxes** to do that.",
			Ephemeral: true,
		})
		return
	}

	page := 1
	pageMax := int(math.Ceil(float64(len(ctx.Player.SpecialBoxes)) / float64(BoxPerPage)))

	if ctx.IsComponent {
		d := ctx.Menu.Data.(*specialBoxesMenuData)

		if d.Page <= 0 {
			d.Page = pageMax
		}
		if d.Page > pageMax {
			d.Page = 1
		}

		page = d.Page
	}

	var n int
	Lootboxes := []*discordgo.MessageEmbedField{}
	for i, spe := range ctx.Player.SpecialBoxes {
		//Correct interval : [(p-1)*max, p*max]
		if n < (page-1)*BoxPerPage || n >= page*BoxPerPage {
			n++
			continue
		}
		n++
		Lootboxes = append(Lootboxes, &discordgo.MessageEmbedField{
			Name:  "N°" + strconv.Itoa(i+1),
			Value: "**" + spe.For + "**",
		})
	}

	Cplx := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s's special loot boxes (page %d/%d) :", ctx.Author.Username, page, pageMax),
			Color:       config.Botcolor,
			Description: "To open a box or view another page use the buttons below.",
			Fields:      Lootboxes,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL("512"),
			},
			Footer: discord.DefaultFooter,
		},
		Components: SpecialBoxesComponent(ctx.ID),
	}

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          HandleSpecialMenu,
		Data: &specialBoxesMenuData{
			UserID: ctx.Author.ID,
			Page:   page,
		},
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content: Cplx,
		Edit:    ctx.IsComponent,
	})
}

func HandleSpecialMenu(ctx *discord.CmdContext) {
	data := ctx.Menu.Data.(*specialBoxesMenuData)

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredMessageUpdate,
		},
	})

	if ctx.Author.ID != data.UserID {
		return
	}

	// Display inventory and market listings of a player
	if len(ctx.Player.SpecialBoxes) == 0 {
		ctx.Reply(discord.ReplyParams{
			Content: "You have no special loot boxes !",
			Edit:    true,
		})
		return
	}

	switch strings.Split(ctx.ComponentData.CustomID, "-")[1] {
	case "next":
		data.Page++
		OpenSpecialCmd(ctx)
		break
	case "1":
		OpenSpecial(ctx, 1*data.Page-1)
		break
	case "2":
		OpenSpecial(ctx, 2*data.Page-1)
		break
	case "3":
		OpenSpecial(ctx, 3*data.Page-1)
		break
	case "4":
		OpenSpecial(ctx, 4*data.Page-1)
		break
	case "5":
		OpenSpecial(ctx, 5*data.Page-1)
		break
	case "prev":
		OpenSpecialCmd(ctx)
		break
	default:
		return
	}

}

func SpecialBoxesComponent(menuID string) []discordgo.MessageComponent {
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
					Label: "Next",
					Emoji: discordgo.ComponentEmoji{
						Name: "➡️",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-next",
				},
			},
		},
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label: "1",
					Emoji: discordgo.ComponentEmoji{
						Name: "1️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-1",
				},
				&discordgo.Button{
					Label: "2",
					Emoji: discordgo.ComponentEmoji{
						Name: "2️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-2",
				},
				&discordgo.Button{
					Label: "3",
					Emoji: discordgo.ComponentEmoji{
						Name: "3️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-3",
				},
				&discordgo.Button{
					Label: "4",
					Emoji: discordgo.ComponentEmoji{
						Name: "4️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-4",
				},
				&discordgo.Button{
					Label: "5",
					Emoji: discordgo.ComponentEmoji{
						Name: "5️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-5",
				},
			},
		},
	}
}
