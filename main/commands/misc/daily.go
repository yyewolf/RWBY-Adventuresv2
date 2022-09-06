package commands_misc

import (
	"fmt"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var DailyCommand = &discord.Command{
	Name:        "daily",
	Description: "Claim your daily reward.",
	Menu:        discord.MiscMenu,
	Call:        daily,
}

func daily(ctx *discord.CmdContext) {
	if !ctx.Player.Status.Voted {
		//Sends the donation link
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageEmbed{
				Title:       "Daily Reward :",
				Description: "You didn't vote for the bot in the past **12 hours**.\nYou can vote [here](https://top.gg/bot/602105650968920094/vote).",
				Color:       config.Botcolor,
			},
		})
		return
	}
	var earnings []string

	ctx.Player.Status.DailyStreak++
	ctx.Player.Status.Voted = false
	ctx.Player.Status.Save()

	money := (rand.Intn(225) + 68) * (ctx.Player.Status.DailyStreak%10 + 1)
	earnings = append(earnings, fmt.Sprintf("**%d**Ⱡ", money))

	var normalLootBoxes int
	var normalGrimmBoxes int
	// random lootbox loot
	for i := 0; i < 3; i++ {
		rng := rand.Float64() * 100
		if rng < 7.5 {
			t := rand.Intn(2)
			if t == 0 {
				ctx.Player.Boxes.Boxes++
				normalLootBoxes++
			} else {
				ctx.Player.Boxes.GrimmBoxes++
				normalGrimmBoxes++
			}
		}
	}

	if normalLootBoxes == 3 {
		normalLootBoxes = 0
		ctx.Player.Boxes.Boxes -= 3
		ctx.Player.Boxes.RareBoxes++
		earnings = append(earnings, "**1** Rare Box")
	}
	if normalGrimmBoxes == 3 {
		normalGrimmBoxes = 0
		ctx.Player.Boxes.GrimmBoxes -= 3
		ctx.Player.Boxes.RareGrimmBoxes++
		earnings = append(earnings, "**1** Rare Grimm Box")
	}

	if normalLootBoxes > 0 {
		earnings = append(earnings, fmt.Sprintf("**%d** Box(es)", normalLootBoxes))
	}
	if normalGrimmBoxes > 0 {
		earnings = append(earnings, fmt.Sprintf("**%d** Grimm Box(es)", normalGrimmBoxes))
	}

	cp := ctx.Player.CalcCP(0.6)
	earnings = append(earnings, fmt.Sprintf("**%d** CP", cp))

	content := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Daily Reward : (%d 🔥)", ctx.Player.Status.DailyStreak),
		Description: "Thank you for your vote!\n\nYou earned :\n",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Author.AvatarURL("512"),
		},
		Color: config.Botcolor,
	}

	for _, earning := range earnings {
		content.Description += fmt.Sprintf("=> %s\n", earning)
	}

	ctx.Reply(discord.ReplyParams{
		Content: content,
	})

	ctx.GiveCP(cp, true)
	ctx.Player.Boxes.Save()
}
