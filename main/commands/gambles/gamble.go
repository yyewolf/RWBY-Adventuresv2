package commands_gambles

import (
	"fmt"
	"math/rand"
	"rwby-adventures/config"
	commands_boxes "rwby-adventures/main/commands/boxes"
	"rwby-adventures/main/discord"
	"time"

	"github.com/bwmarrin/discordgo"
)

func gamble(ctx *discord.CmdContext) {
	if !ctx.Player.CanGamble() {
		t := ctx.Player.GambleCooldown()
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("Sorry but you still have to wait **%dh %dm and %ds** before you can gamble again", int(t.Hours()), int(t.Minutes())%60, int(t.Seconds())%60),
			Ephemeral: true,
		})
		return
	}
	if ctx.Player.TotalBalance() < 700 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have enough liens to gamble right now, you need at least 700 Liens.",
			Ephemeral: true,
		})
		return
	}

	canLootChar := ctx.Player.CharAmount() < ctx.Player.MaxChar()

	var lucky bool
	//Luck Boosterreturn
	if ctx.Player.Shop.LuckBoost {
		lucky = true
		ctx.Player.Shop.LuckBoostTime--
		ctx.Player.SendLuckNotice(ctx.Session)
		ctx.Player.Shop.Save()
	}

	// l is used to have more chances of getting the same thing twice
	//loot1, display1, l := gamblePickLoot(canLootChar, lucky, 0)
	//loot2, display2, l := gamblePickLoot(canLootChar, lucky, l)
	//loot3, display3, l := gamblePickLoot(canLootChar, lucky, l)

	// /*
	// 	STATS
	// */
	// if lucky {
	// 	iGAmt++
	// } else {
	// 	oGAmt++
	// }

	all := pickLoots(canLootChar, lucky)

	loots := []string{all[0].name, all[1].name, all[2].name}
	displays := []string{all[0].display, all[1].display, all[2].display}
	counts := make(map[string]int)
	var lootText string
	for _, val := range loots {
		counts[val]++
	}
	if counts["xp"] >= 2 {
		var XP int64
		if counts["xp"] == 3 {
			XP = ctx.Player.CalcSelectedXP(7, false)
		} else {
			XP = ctx.Player.CalcSelectedXP(4, false)
		}
		lootText += fmt.Sprintf("Good job, you earned **%d**XP !", XP)
		ctx.GiveSelectionXP(XP, true)
	} else if counts["money"] >= 2 {
		var liens int64
		if counts["money"] == 3 {
			liens = rand.Int63n(100) + 680
		} else {
			liens = rand.Int63n(70) + 340
		}
		lootText += fmt.Sprintf("Good job, you earned **%d**Ⱡ !", liens)
		ctx.Player.Balance += liens

		ctx.Player.Save()
	} else if counts["lootbox"] >= 2 {
		if counts["lootbox"] == 3 {
			if ctx.Player.Lootbox()+2 <= ctx.Player.MaxChar() {
				ctx.Player.Boxes.Boxes++
				ctx.Player.Boxes.GrimmBoxes++
				lootText += "Good job, you just won 1 loot box and 1 grimm box !"

				ctx.Player.Boxes.Save()
			} else {
				liens := rand.Int63n(100) + 680
				lootText += fmt.Sprintf("Good job, you earned %dⱠ !", liens)
				ctx.Player.Balance += liens

				ctx.Player.Save()
			}
		} else {
			if ctx.Player.Lootbox()+1 <= ctx.Player.MaxChar() {
				ctx.Player.Boxes.Boxes++
				lootText += "Good job, you just won 1 loot box !"

				ctx.Player.Boxes.Save()
			} else {
				liens := rand.Int63n(70) + 340
				lootText += fmt.Sprintf("Good job, you earned %dⱠ !", liens)
				ctx.Player.Balance += liens

				ctx.Player.Save()
			}
		}
	} else if counts["rarelootbox"] >= 2 {
		if counts["rarelootbox"] == 3 {
			ctx.Player.Boxes.RareBoxes++
			ctx.Player.Boxes.RareGrimmBoxes++
			lootText += "Good job, you just won 1 rare loot box and 1 rare grimm box !"

			ctx.Player.Boxes.Save()
		} else {
			ctx.Player.Boxes.RareBoxes++
			lootText += "Good job, you just won 1 rare loot box !"

			ctx.Player.Boxes.Save()
		}
	} else if counts["char"] >= 3 {
		//Adds cool badge to winner!
		ctx.Player.AddBadgeName("Lucky boi")
		rng := rand.Intn(2)

		b := &commands_boxes.BoxFilter{
			Box:        "Gamble",
			ValStd:     15,
			ValMean:    62.5,
			RarityRate: 1,
		}
		if rng == 0 {
			commands_boxes.OpenChar(ctx, b)
		} else {
			commands_boxes.OpenGrimm(ctx, b)
		}
	} else if counts["arm"] >= 3 {
		ctx.Player.Arms++
		lootText += "Good job, you just won an arm !"

		ctx.Player.Save()
	} else if counts["lose"] >= 2 {
		var liens int64
		if counts["lose"] == 3 {
			liens = rand.Int63n(100) + 380
			lootText += fmt.Sprintf("Too bad, you lost **%d**Ⱡ! That must be hard...", liens)
		} else {
			liens = rand.Int63n(70) + 130
			lootText += fmt.Sprintf("Too bad, you lost **%d**Ⱡ!", liens)
		}
		ctx.Player.Balance -= liens

		ctx.Player.Save()
	} else {
		liens := rand.Int63n(30) + 70
		lootText += fmt.Sprintf("Too bad, you lost **%d**Ⱡ!", liens)
		ctx.Player.Balance -= liens

		ctx.Player.Save()
	}
	ctx.Player.Gamble.Amount++
	ctx.Player.Gamble.Time = time.Now().Unix()
	ctx.Player.Gamble.Save()

	//Base Message
	images := createGambleImage(loots[0], loots[1], loots[2])
	if len(images) == 0 {
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.MessageEmbed{
				Title:       "RWBY Gamble !",
				Color:       config.Botcolor,
				Description: fmt.Sprintf("You're doing a Gamble ? Here is what you got :\n\n%s", lootText),
				Footer:      discord.DefaultFooter,
			},
		})
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "RWBY Gamble !",
		Color:       config.Botcolor,
		Description: "You're doing a Gamble ? Here is what you got :",
		Image: &discordgo.MessageEmbedImage{
			URL: config.GambleHost + images[0],
		},
		Footer: discord.DefaultFooter,
	}
	_, err := ctx.Reply(discord.ReplyParams{
		Content: embed,
	})
	fmt.Println(err)

	time.Sleep(5 * time.Second)
	//ANIMATION
	for i := 0; i < 3; i++ {
		embed.Description += "\n**" + displays[i] + "** !"
		embed.Image.URL = config.GambleHost + images[i+1]
		ctx.Reply(discord.ReplyParams{
			Content: embed,
			Edit:    true,
		})
		time.Sleep(6 * time.Second)
	}
	embed.Description += "\n\n" + lootText
	ctx.Reply(discord.ReplyParams{
		Content: embed,
		Edit:    true,
	})
}
