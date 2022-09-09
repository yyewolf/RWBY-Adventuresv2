package discord

import (
	"fmt"
	"math"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/models"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func UnixFormatter(timestamp int64) string {
	t := timestamp - time.Now().Unix()
	H := int(float64(t) / float64(3600))
	M := int(math.Mod(float64(t)/float64(60), 60))
	S := int(math.Mod(float64(t), 60))

	return fmt.Sprintf("**%02d hours %02d minutes %02d seconds**", H, M, S)
}

func UnixNanoFormatter(timestamp int64) string {
	t := timestamp - time.Now().UnixNano()
	H := int(float64(t) / float64(3600))
	M := int(math.Mod(float64(t)/float64(60), 60))
	S := int(math.Mod(float64(t), 60))

	return fmt.Sprintf("**%02d hours %02d minutes %02d seconds**", H, M, S)
}

func (ctx *CmdContext) GiveSelectionXP(add int64, notif bool) {
	ctx.Player.Status.LastXP = time.Now().Unix()
	config.Database.Save(ctx.Player.Status)
	if ctx.Player.SelectedChar.Valid() {
		ctx.GiveCharXP(ctx.Player.SelectedChar, add, notif)
	}
	if ctx.Player.SelectedGrimm.Valid() {
		ctx.GiveGrimmXP(ctx.Player.SelectedGrimm, add, notif)
	}
}

func (ctx *CmdContext) GiveXP(char *models.Character, grimm *models.Grimm, add int64, notif bool) {
	ctx.Player.Status.LastXP = time.Now().Unix()
	config.Database.Save(ctx.Player.Status)
	if char.Valid() {
		ctx.GiveCharXP(char, add, notif)
	}
	if grimm.Valid() {
		ctx.GiveGrimmXP(grimm, add, notif)
	}
}

func (ctx *CmdContext) GiveCharXP(char *models.Character, add int64, notif bool) {
	if ctx.Player.Shop.XPBoost && ctx.Player.Shop.XPBoostTime > 0 {
		// ctx.Player.sendXPNotice(s)
		ctx.Player.Shop.XPBoostTime--
		ctx.Player.Shop.Save()
	}
	levelUp := char.GiveXP(add)
	if levelUp && notif {
		content := &discordgo.MessageEmbed{
			Title: "Congratulations !",
			Description: fmt.Sprintf("%s, your **%s** has leveled up!\n", ctx.Author.Mention(), char.Name) +
				fmt.Sprintf("Level : **%d**.\n", char.Level) +
				fmt.Sprintf("XP : **%d/%d**", char.XP, char.XPCap),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: char.ToRealChar().IconURL,
			},
			Color: config.Botcolor,
		}
		ctx.Reply(ReplyParams{
			Content: content,
			// ID:       ctx.Author.ID,
			// DM:       true,
			Automated: true,
			FollowUp:  true,
		})
		char.Stats.Save()
	}

	//Adds XP to char
	char.Save()
}

func (ctx *CmdContext) GiveGrimmXP(grimm *models.Grimm, add int64, notif bool) {
	if ctx.Player.Shop.XPBoost && ctx.Player.Shop.XPBoostTime > 0 {
		// ctx.Player.sendXPNotice(s)
		ctx.Player.Shop.XPBoostTime--
		config.Database.Save(ctx.Player.Shop)
	}
	levelUp := grimm.GiveXP(add)
	if levelUp && notif {
		content := &discordgo.MessageEmbed{
			Title: "Congratulations !",
			Description: fmt.Sprintf("%s, your **%s** has leveled up!\n", ctx.Author.Mention(), grimm.Name) +
				fmt.Sprintf("Level : **%d**.\n", grimm.Level) +
				fmt.Sprintf("XP : **%d/%d**", grimm.XP, grimm.XPCap),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: grimm.ToRealGrimm().IconURL,
			},
			Color: config.Botcolor,
		}
		ctx.Reply(ReplyParams{
			Content:   content,
			Automated: true,
			FollowUp:  true,
		})
	}

	//Adds XP to char
	config.Database.Save(grimm)
}

func (ctx *CmdContext) GiveCP(CP int64, notif bool) *discordgo.MessageEmbed {
	if ctx.Player.Shop.XPBoost && ctx.Player.Shop.XPBoostTime > 0 {
		// ctx.Player.sendXPNotice(s)
		ctx.Player.Shop.XPBoostTime--
		config.Database.Save(ctx.Player.Shop)
	}
	before := ctx.Player.Level
	levelUp := ctx.Player.GiveCP(CP)
	levelEarned := ctx.Player.Level - before
	if levelUp {
		var lootboxes int
		var grimmboxes int
		var money int64
		var arms int
		var minions int
		var backpacks int

		for i := int64(0); i < levelEarned; i++ {
			// Lootboxes
			rng := rand.Float64() * 100
			if rng < 12.5 {
				amount := rand.Intn(int(math.Sqrt(float64(ctx.Player.Level)))) + 1
				lootboxes += amount
			}

			// Grimmboxes
			rng = rand.Float64() * 100
			if rng < 12.5 {
				amount := rand.Intn(int(math.Sqrt(float64(ctx.Player.Level)))) + 1
				grimmboxes += amount
			}

			// Liens
			money += rand.Int63n(153+(ctx.Player.Level-i+1)*6) + 54

			// Every 10 levels
			if (ctx.Player.Level-i+1)%10 == 0 {
				// 17.5% chance
				rng = rand.Float64() * 100
				if rng < 17.5 {
					if rng < 5 {
						arms++
					} else {
						minions++
					}
				}

				// 10% chance
				rng = rand.Float64() * 100
				if rng < 10 {
					backpacks++
				}
			}
		}

		// Save everything
		ctx.Player.Boxes.Boxes += lootboxes
		ctx.Player.Boxes.GrimmBoxes += grimmboxes
		ctx.Player.Balance += money
		ctx.Player.Arms += arms
		ctx.Player.Minions += minions
		ctx.Player.Shop.Extensions += backpacks

		var earnings []string
		earnings = append(earnings, fmt.Sprintf("%d Box(es)", lootboxes))
		earnings = append(earnings, fmt.Sprintf("%d Grimm Box(es)", grimmboxes))
		earnings = append(earnings, fmt.Sprintf("%d Ⱡ (Liens)", money))
		earnings = append(earnings, fmt.Sprintf("%d Arms", arms))
		earnings = append(earnings, fmt.Sprintf("%d Minions", minions))
		earnings = append(earnings, fmt.Sprintf("%d Backpacks", backpacks))

		content := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s, you just leveled up !", ctx.Author.Username),
			Description: fmt.Sprintf("You are now level **%d**!\n\nYou earned :\n", ctx.Player.Level),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL("512"),
			},
			Color: config.Botcolor,
		}

		for _, earning := range earnings {
			content.Description += fmt.Sprintf("=> %s\n", earning)
		}

		// Save everything to database
		config.Database.Save(ctx.Player)
		config.Database.Save(ctx.Player.Shop)
		config.Database.Save(ctx.Player.Boxes)

		if notif {
			ctx.Reply(ReplyParams{
				Content:  content,
				FollowUp: true,
			})
		} else {
			return content
		}
	}

	config.Database.Save(ctx.Player)
	return nil
}

func (ctx *CmdContext) GetUnix() (i int64, err error) {
	i, err = strconv.ParseInt(ctx.ID, 10, 64)
	if err != nil {
		return
	}
	i = int64(float64((i/4194304)+1420070400000) * 0.001)
	return
}

func (ctx *CmdContext) SendLuckBoostNotice() {
	if ctx.Player.Shop.LuckBoostTime != 1 {
		return
	}
	embed := &discordgo.MessageEmbed{
		Title: "Your Luck Boost is running out !",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://pm1.narvii.com/6719/01c7db349c4c882b866c06aeb1e3784c6e0c30fc_hq.jpg",
		},
		Color: config.Botcolor,
	}
	ctx.Reply(ReplyParams{
		Content: embed,
		ID:      ctx.Author.ID,
		DM:      true,
	})
}
