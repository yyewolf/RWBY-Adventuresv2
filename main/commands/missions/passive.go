package commands_missions

import (
	"fmt"
	"math"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func PassiveMissions(ctx *discord.CmdContext) {
	if ctx.Player.Missions.IsInMission {
		continueMission(ctx)
	}
	if ctx.Player.Missions.IsInHunt {
		continueHunt(ctx)
	}

	if !ctx.Player.Missions.IsInMission && !ctx.Player.Missions.IsInHunt {
		newMission(ctx)
	}
}

func newMission(ctx *discord.CmdContext) {
	//Rolls the dice now to save computing power
	rng := rand.Float64() * 100
	//0.20% chance
	if rng > 10 {
		return
	}

	if !ctx.Player.Missions.IsInHunt && !ctx.Player.Missions.IsInMission && !ctx.Player.Missions.CanGoHunt && !ctx.Player.Missions.CanGoToMission {
		choice := rand.Float64()
		if choice < 0.5 {
			spawnHunt(ctx)
			return
		}
		spawnMission(ctx)
	} else if !ctx.Player.Missions.IsInMission && !ctx.Player.Missions.CanGoToMission {
		spawnMission(ctx)
	} else if !ctx.Player.Missions.IsInHunt && !ctx.Player.Missions.CanGoHunt {
		spawnHunt(ctx)
	}
}

func spawnMission(ctx *discord.CmdContext) {
	ctx.Player.Missions.CanGoToMission = true
	ctx.Player.Missions.MissionType = rand.Intn(6)
	ctx.Player.Missions.MissionMsgLeft = 150 + rand.Intn(150)

	days := int(math.Ceil(float64(ctx.Player.Missions.MissionMsgLeft) / 24))
	//Prepare the embed
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("A mission is available for you, %s !", ctx.Author.Username),
		Description: fmt.Sprintf("Use </mission go:%s> to go on a mission, %s!\n", MissionCommand.ID, ctx.Author.Mention()) +
			fmt.Sprintf("**%s**.\n", missionToString(ctx.Player.Missions.MissionType)) +
			"Your character will leave your inventory and come back after a while.\n" +
			fmt.Sprintf("Your character will be gone for : **%d days**.\nTo end a mission by yourself use </mission end:%s>.", days, MissionCommand.ID),
		Color: config.Botcolor,
	}
	ctx.Reply(discord.ReplyParams{
		Content:   embed,
		Automated: true,
	})

	config.Database.Save(ctx.Player.Missions)
}

func spawnHunt(ctx *discord.CmdContext) {
	ctx.Player.Missions.CanGoHunt = true
	ctx.Player.Missions.HuntType = rand.Intn(6)
	ctx.Player.Missions.HuntMsgLeft = 150 + rand.Intn(150)

	days := int(math.Ceil(float64(ctx.Player.Missions.HuntMsgLeft) / 24))
	//Prepare the embed
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("A hunt is available for you, %s !", ctx.Author.Username),
		Description: fmt.Sprintf("Use </hunt go:%s> to go on a hunt, %s!\n", HuntCommand.ID, ctx.Author.Mention()) +
			fmt.Sprintf("**%s**.\n", huntToString(ctx.Player.Missions.HuntType)) +
			"Your grimm will leave your inventory and come back after a while.\n" +
			fmt.Sprintf("Your grimm will be gone for : **%d days**.\nTo end a hunt by yourself use </hunt end:%s>.", days, HuntCommand.ID),
		Color: config.Botcolor,
	}
	ctx.Reply(discord.ReplyParams{
		Content:   embed,
		Automated: true,
	})

	config.Database.Save(ctx.Player.Missions)
}

func continueMission(ctx *discord.CmdContext) {
	ctx.Player.Missions.MissionMsgLeft--
	if ctx.Player.Missions.MissionMsgLeft > 0 {
		config.Database.Save(ctx.Player.Missions)
		return
	}
	ctx.Player.Missions.IsInMission = false
	config.Database.Save(ctx.Player.Missions)
	ctx.Player.CharInMission.InMission = false
	config.Database.Save(ctx.Player.CharInMission)

	//MONEY LOOT
	Money := rand.Intn(300) + 75
	//Saves player
	ctx.Player.Status.LastXP = time.Now().Unix()
	ctx.Player.Balance += int64(Money)
	config.Database.Save(ctx.Player.Status)

	//MISSION FINISHED
	earningText := strconv.Itoa(Money) + "Ⱡ"
	//LOOTBOX LOOTs
	rand.Seed(time.Now().UTC().UnixNano())
	randfloat := rand.Float64()
	//1% chance
	hasLootbox := randfloat < 0.01
	if hasLootbox && ctx.Player.Lootbox() < ctx.Player.MaxChar() {
		ctx.Player.Boxes.Boxes++
		config.Database.Save(ctx.Player.Boxes)
		earningText += " and 1 lootbox"
	}
	add := ctx.Player.CharInMission.CalcXP(10, ctx.Player.Shop.XPBoost)
	earningText += fmt.Sprintf("\nYour character also earned : %dXP.", add)
	ctx.GiveXP(ctx.Player.CharInMission, nil, add, false)
	//END XP
	//SEND MESSAGE
	winText := fmt.Sprintf("Hey, %s !\n", ctx.Author.Username)
	winText += missionWinMessages(ctx.Player.Missions.MissionType)
	winText = strings.ReplaceAll(winText, "{mission.Earned}", earningText)

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("Your character is back, %s.", ctx.Author.Username),
			Description: winText,
			Color:       config.Botcolor,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL("512"),
			},
			Footer: discord.DefaultFooter,
		},
		ID: ctx.Author.ID,
		DM: true,
	})
	ctx.Player.SelectedChar = ctx.Player.CharInMission
	ctx.Player.SelectedID = ctx.Player.CharInMission.CharID
	ctx.Player.SelectedType = models.CharType
	ctx.Player.CharInMission = nil
	config.Database.Save(ctx.Player)

	for _, v := range ctx.Player.Characters {
		v.InMission = false
		config.Database.Save(v)
	}
}

func continueHunt(ctx *discord.CmdContext) {
	ctx.Player.Missions.HuntMsgLeft--
	if ctx.Player.Missions.HuntMsgLeft > 0 {
		config.Database.Save(ctx.Player.Missions)
		return
	}
	ctx.Player.Missions.IsInHunt = false
	config.Database.Save(ctx.Player.Missions)
	ctx.Player.GrimmInHunt.InHunt = false
	config.Database.Save(ctx.Player.GrimmInHunt)

	//MONEY LOOT
	Money := rand.Intn(300) + 75
	//Saves player
	ctx.Player.Status.LastXP = time.Now().Unix()
	ctx.Player.Balance += int64(Money)
	config.Database.Save(ctx.Player.Status)

	//MISSION FINISHED
	earningText := strconv.Itoa(Money) + "Ⱡ"
	//LOOTBOX LOOTs
	rand.Seed(time.Now().UTC().UnixNano())
	randfloat := rand.Float64()
	//1% chance
	hasLootbox := randfloat < 0.01
	if hasLootbox && ctx.Player.Lootbox() < ctx.Player.MaxChar() {
		ctx.Player.Boxes.Boxes++
		config.Database.Save(ctx.Player.Boxes)
		earningText += " and 1 lootbox"
	}
	add := ctx.Player.GrimmInHunt.CalcXP(10, ctx.Player.Shop.XPBoost)
	earningText += fmt.Sprintf("\nYour grimm also earned : %dXP.", add)
	ctx.GiveXP(nil, ctx.Player.GrimmInHunt, add, false)
	//END XP
	//SEND MESSAGE
	winText := fmt.Sprintf("Hey, %s !\n", ctx.Author.Username)
	winText += huntWinMessages(ctx.Player.Missions.HuntType)
	winText = strings.ReplaceAll(winText, "{hunt.Earned}", earningText)

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("Your grimm is back, %s.", ctx.Author.Username),
			Description: winText,
			Color:       config.Botcolor,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL("512"),
			},
			Footer: discord.DefaultFooter,
		},
		ID: ctx.Author.ID,
		DM: true,
	})
	ctx.Player.SelectedGrimm = ctx.Player.GrimmInHunt
	ctx.Player.SelectedID = ctx.Player.GrimmInHunt.GrimmID
	ctx.Player.SelectedType = models.GrimmType
	ctx.Player.GrimmInHunt = nil
	config.Database.Save(ctx.Player)

	for _, v := range ctx.Player.Grimms {
		v.InHunt = false
		config.Database.Save(v)
	}
}
