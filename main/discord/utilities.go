package discord

import (
	"fmt"
	"math"
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
	if ctx.Player.SelectedChar != nil {
		ctx.GiveCharXP(ctx.Player.SelectedChar, add, notif)
	}
	if ctx.Player.SelectedGrimm != nil {
		ctx.GiveGrimmXP(ctx.Player.SelectedGrimm, add, notif)
	}
}

func (ctx *CmdContext) GiveXP(char *models.Character, grimm *models.Grimm, add int64, notif bool) {
	ctx.Player.Status.LastXP = time.Now().Unix()
	config.Database.Save(ctx.Player.Status)
	if char != nil {
		ctx.GiveCharXP(char, add, notif)
	}
	if grimm != nil {
		ctx.GiveGrimmXP(grimm, add, notif)
	}
}

func (ctx *CmdContext) GiveCharXP(char *models.Character, add int64, notif bool) {
	levelUp := false
	if ctx.Player.Shop.XPBoost && ctx.Player.Shop.XPBoostTime > 0 {
		// ctx.Player.sendXPNotice(s)
		ctx.Player.Shop.XPBoostTime--
		config.Database.Save(ctx.Player.Shop)
	}
	for char.XP+add > char.XPCap {
		levelUp = true
		//if level up
		add -= char.XPCap - char.XP
		char.Level++
		char.XP = 0
		char.XPCap = char.CalcXPCap()
		char.CalcStats()
	}
	char.XP += add
	char.XPCap = char.CalcXPCap()
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
			Content:  content,
			FollowUp: true,
		})
	}

	//Adds XP to char
	config.Database.Save(char)
}

func (ctx *CmdContext) GiveGrimmXP(grimm *models.Grimm, add int64, notif bool) {
	levelUp := false
	if ctx.Player.Shop.XPBoost && ctx.Player.Shop.XPBoostTime > 0 {
		// ctx.Player.sendXPNotice(s)
		ctx.Player.Shop.XPBoostTime--
		config.Database.Save(ctx.Player.Shop)
	}
	for grimm.XP+add > grimm.XPCap {
		levelUp = true
		//if level up
		add -= grimm.XPCap - grimm.XP
		grimm.Level++
		grimm.XP = 0
		grimm.XPCap = grimm.CalcXPCap()
		grimm.CalcStats()
	}
	grimm.XP += add
	grimm.XPCap = grimm.CalcXPCap()
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
			Content:  content,
			FollowUp: true,
		})
	}

	//Adds XP to char
	config.Database.Save(grimm)
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
		DM:      true,
	})
}
