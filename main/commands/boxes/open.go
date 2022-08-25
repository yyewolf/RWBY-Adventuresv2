package commands_boxes

import (
	"bytes"
	"fmt"
	"math/rand"
	chars "rwby-adventures/characters"
	"rwby-adventures/config"
	"rwby-adventures/grimms"
	commands_inventory "rwby-adventures/main/commands/inventory"
	"rwby-adventures/main/discord"
	"rwby-adventures/main/static"
	"rwby-adventures/models"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	uuid "github.com/satori/go.uuid"
)

type BoxFilter struct {
	IncludeLimited bool
	OnlyLimited    bool
	Category       string
	Box            string

	// Value
	ValStd  float64
	ValMean float64

	// Rarity
	RarityRate float64
}

type openMenuData struct {
	UserID    string
	Character *models.Character
	Grimm     *models.Grimm
	Type      int
}

func (b *BoxFilter) FilterChars() (List []chars.CharacterStruct) {
	for _, char := range chars.BaseCharacters {
		if !strings.Contains(char.Category, b.Category) {
			continue
		}
		if b.OnlyLimited && !char.Limited {
			continue
		}
		if b.IncludeLimited && char.Limited {
			List = append(List, char)
		}
		if !char.Limited {
			List = append(List, char)
		}
	}
	return
}

func (b *BoxFilter) FilterGrimms() (List []grimms.GrimmStruct) {
	for _, grimm := range grimms.BaseGrimms {
		if !strings.Contains(grimm.Category, b.Category) {
			continue
		}
		if b.OnlyLimited && !grimm.Limited {
			continue
		}
		if b.IncludeLimited && grimm.Limited {
			List = append(List, grimm)
		}
		if !grimm.Limited {
			List = append(List, grimm)
		}
	}
	return
}

func OpenChar(ctx *discord.CmdContext, b *BoxFilter) (success bool) {
	var isLucky bool

	L := b.FilterChars()
	Loot := L[rand.Intn(len(L))]

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

	Loot.Value = Value

	// Rarity finder
	var Rarity int
	rdm := rand.ExpFloat64() / b.RarityRate
	for i := 0; float64(i) < rdm; i++ {
		if float64(i+1) > rdm {
			Rarity = i
			break
		}
	}

	// Level finder
	Rate := 0.4
	if isLucky {
		Rate = 0.3
	}
	Level := int(rand.ExpFloat64()/Rate) + 1
	for Level > 40 {
		Level = int(rand.ExpFloat64()/Rate) + 1
	}

	ID := uuid.NewV4().String()
	Char := &models.Character{
		Name:   Loot.Name,
		CharID: ID,
		UserID: ctx.Author.ID,
		Value:  Loot.Value,
		Level:  Level,
		XP:     0,
		Stats:  models.CharacterStats(Loot.Stats),
		Rarity: Rarity,
	}
	Char.XPCap = Char.CalcXPCap()
	Char.CalcStats()
	config.Database.Create(Char)

	Cplx := Char.ToLootedEmbed(ctx.Author.Mention(), ctx.ID, b.Box, &Loot)
	Cplx.Embed.Footer = discord.DefaultFooter

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          openMenu,
		Data: &openMenuData{
			UserID:    ctx.Author.ID,
			Character: Char,
			Type:      models.CharType,
		},
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content: Cplx,
	})
	return true
}

func OpenGrimm(ctx *discord.CmdContext, b *BoxFilter) (success bool) {
	var isLucky bool

	L := b.FilterGrimms()
	Loot := L[rand.Intn(len(L))]

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

	Loot.Value = Value

	// Rarity finder
	var Rarity int
	rdm := rand.ExpFloat64() / b.RarityRate
	for i := 0; float64(i) < rdm; i++ {
		if float64(i+1) > rdm {
			Rarity = i
			break
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
		Value:   Loot.Value,
		Level:   Level,
		XP:      0,
		Stats:   models.GrimmStat(Loot.Stats),
		Rarity:  Rarity,
	}
	Grimm.XPCap = Grimm.CalcXPCap()
	Grimm.CalcStats()
	config.Database.Create(Grimm)

	Cplx := Grimm.ToLootedEmbed(ctx.Author.Mention(), ctx.ID, b.Box, &Loot)
	Cplx.Embed.Footer = discord.DefaultFooter

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          openMenu,
		Data: &openMenuData{
			UserID: ctx.Author.ID,
			Grimm:  Grimm,
			Type:   models.GrimmType,
		},
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content: Cplx,
	})
	return true
}

func openMenu(ctx *discord.CmdContext) {
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if ctx.Author.ID != ctx.Menu.SourceContext.Author.ID {
		return
	}
	d := ctx.Menu.Data.(*openMenuData)

	if d.Type == models.CharType {
		// Useful data
		char := d.Character
		original := char.ToRealChar()
		imgData, _ := static.DatabaseFS.ReadFile(original.ImageFile)
		imgDecoded := bytes.NewBuffer(imgData)

		// DiscordGo Stuff

		discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
			MenuID:        ctx.ID,
			SourceContext: ctx,
			Call:          commands_inventory.MenuInfo,
			Data: &commands_inventory.InfoMenuData{
				UserID:  ctx.Author.ID,
				Char:    char,
				IsGrimm: false,
			},
		}, 0)

		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageSend{
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
				Components: commands_inventory.InfoComponent(ctx.ID),
			},
			Edit: true,
		})
	} else {
		// Useful data
		grimm := d.Grimm
		original := grimm.ToRealGrimm()
		imgData, _ := static.DatabaseFS.ReadFile(original.ImageFile)
		imgDecoded := bytes.NewBuffer(imgData)

		// DiscordGo Stuff

		discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
			MenuID:        ctx.ID,
			SourceContext: ctx,
			Call:          commands_inventory.MenuInfo,
			Data: &commands_inventory.InfoMenuData{
				UserID:  ctx.Author.ID,
				Grimm:   grimm,
				IsGrimm: true,
			},
		}, 0)

		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageSend{
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
				Components: commands_inventory.InfoComponent(ctx.ID),
			},
			Edit: true,
		})
	}
}
