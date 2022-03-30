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
	} else {
		newMission(ctx)
	}
	if ctx.Player.Missions.IsInHunt {
		continueHunt(ctx)
	} else {
		newMission(ctx)
	}
}

func newMission(ctx *discord.CmdContext) {
	//Rolls the dice now to save computing power
	randfloat := rand.Float64()
	//0.20% chance
	willSpawn := randfloat < 1
	if !willSpawn {
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

	days := strconv.Itoa(int(math.Ceil(float64(ctx.Player.Missions.MissionMsgLeft) / 24)))
	//Prepare the embed
	g := ctx.Guild
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("A mission is available for you, %s !", ctx.Author.Username),
		Description: "Type `" + g.Prefix + "mission go` to go on a mission, " + ctx.Author.Mention() + " !\n" +
			"**" + missionToString(ctx.Player.Missions.MissionType) + "**.\n" +
			"Your character will leave your inventory and come back after a while.\n" +
			"Your character will be gone for : **" + days + " days**.\n" +
			"To end a mission by yourself type `" + g.Prefix + "mission end`.",
		Color: config.Botcolor,
	}
	ctx.Reply(discord.ReplyParams{
		Content: embed,
	})

	config.Database.Save(ctx.Player.Missions)
}

func spawnHunt(ctx *discord.CmdContext) {
	ctx.Player.Missions.CanGoHunt = true
	ctx.Player.Missions.HuntType = rand.Intn(6)
	ctx.Player.Missions.HuntMsgLeft = 150 + rand.Intn(150)

	days := strconv.Itoa(int(math.Ceil(float64(ctx.Player.Missions.MissionMsgLeft) / 24)))
	//Prepare the embed
	g := ctx.Guild
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("A hunt is available for you, %s !", ctx.Author.Username),
		Description: "Type `" + g.Prefix + "hunt go` to go on a hunt, " + ctx.Author.Mention() + " !\n" +
			"**" + huntToString(ctx.Player.Missions.HuntType) + "**.\n" +
			"Your grimm will leave your inventory and come back after a while.\n" +
			"Your grimm will be gone for : **" + days + " days**.\n" +
			"To end a mission by yourself type `" + g.Prefix + "mission end`.",
		Color: config.Botcolor,
	}
	ctx.Reply(discord.ReplyParams{
		Content: embed,
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
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Your character is back, %s.", ctx.Author.Username),
		Description: winText,
		Color:       config.Botcolor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Author.AvatarURL("512"),
		},
		Footer: discord.DefaultFooter,
	}

	ctx.Reply(discord.ReplyParams{
		Content: embed,
		DM:      true,
	})
	ctx.Player.SelectedChar = ctx.Player.CharInMission
	ctx.Player.SelectedID = ctx.Player.CharInMission.CharID
	ctx.Player.SelectedType = models.CharType
	ctx.Player.CharInMission = nil
	config.Database.Save(ctx.Player)
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
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Your grimm is back, %s.", ctx.Author.Username),
		Description: winText,
		Color:       config.Botcolor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Author.AvatarURL("512"),
		},
		Footer: discord.DefaultFooter,
	}

	ctx.Reply(discord.ReplyParams{
		Content: embed,
		DM:      true,
	})
	ctx.Player.SelectedGrimm = ctx.Player.GrimmInHunt
	ctx.Player.SelectedID = ctx.Player.GrimmInHunt.GrimmID
	ctx.Player.SelectedType = models.GrimmType
	ctx.Player.GrimmInHunt = nil
	config.Database.Save(ctx.Player)
}
