package commands

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func TimeLeftString(startTime, endTime int64) string {
	t := (startTime + endTime) - time.Now().Unix()
	H := int(float64(t) / float64(3600))
	M := int(math.Mod(float64(t)/float64(60), 60))
	S := int(math.Mod(float64(t), 60))

	return fmt.Sprintf("**%02d hours %02d minutes %02d seconds**", H, M, S)
}

func BoolStringFormat(b bool) string {
	if b {
		return "✔️"
	}
	return "❌"
}

func Me(ctx *discord.CmdContext) {
	// Useful variable
	p := ctx.Player

	// Strings
	var resetStringBox = "✅ - `Loots` (**" + strconv.Itoa(p.LastBoxes.Amount) + "/" + strconv.Itoa(p.Maxlootbox) + "**)"
	var resetStringGamble = "✅ - `Gamble` (**" + strconv.Itoa(p.Gamble.Amount) + "/3**)"
	var resetStringDungeon = "✅ - `Dungeon`"
	var missionTime = ""
	var huntTime = ""
	var luckBoostString string
	var charLimited int
	var grimmLimited int

	for _, b := range p.LimitedBoxes {
		if b.Type == models.CharType {
			charLimited++
		} else {
			grimmLimited++
		}
	}

	canLootBox, _ := p.CanDropLootBox()
	if !canLootBox {
		resetStringBox = "🕓 - `Loots` (" + TimeLeftString(p.LastBoxes.Time, 86400) + ")"
	}
	canGamble, _ := p.CanGamble()
	if !canGamble {
		resetStringGamble = "🕓 - `Gamble` (" + TimeLeftString(p.Gamble.Time, 86400) + ")"
	}
	canDungeon := p.CanDungeon()
	if !canDungeon {
		resetStringDungeon = "🕓 - `Dungeon` (" + TimeLeftString(p.Status.LastDungeon, 18000) + ")" //5*60*60
	}
	if p.Missions.IsInMission {
		days := strconv.Itoa(int(math.Ceil(float64(p.Missions.MissionMsgLeft) / 24)))
		missionTime = "`Time left : " + days + " days.`\n"
	}
	if p.Missions.IsInHunt {
		days := strconv.Itoa(int(math.Ceil(float64(p.Missions.HuntMsgLeft) / 24)))
		huntTime = "\n`Time left : " + days + " days.`"
	}
	if p.Shop.LuckBoost {
		luckBoostString = " (" + strconv.Itoa(p.Shop.LuckBoostTime) + "/20)"
	}

	// Embed Profile

	infoProfile := &discordgo.MessageEmbedField{
		Name: "**Player :**",
		Value: "Level **" + strconv.FormatInt(p.Level, 10) + "**.\n" +
			"CP : **" + strconv.FormatInt(p.CP, 10) + "**/**" + strconv.FormatInt(p.MaxCP, 10) + "**.\n" +
			"Slots : **" + strconv.Itoa(p.MaxChar()) + "**.\n" +
			"Boxes : **" + strconv.Itoa(p.Lootbox()) + "/" + strconv.Itoa(p.MaxChar()) + "**." + "\n" +
			"Liens : **" + strconv.FormatInt(p.Balance, 10) + "Ⱡ**.",
		Inline: true,
	}

	// Embed cool downs

	infoCoolDown := &discordgo.MessageEmbedField{
		Name: "**Cooldowns :**",
		Value: resetStringBox + "\n" +
			resetStringGamble + "\n" +
			resetStringDungeon,
		Inline: true,
	}

	// Embed missions

	infoMission := &discordgo.MessageEmbedField{
		Name: "**Mission / Hunt :**",
		Value: "`Can mission : " + BoolStringFormat(p.Missions.CanGoToMission) + "`\n" +
			"`Is in mission : " + BoolStringFormat(p.Missions.IsInMission) + "`\n" +
			missionTime +
			"`================`\n" +
			"`Can hunt : " + BoolStringFormat(p.Missions.CanGoHunt) + "`\n" +
			"`Is in hunt : " + BoolStringFormat(p.Missions.IsInHunt) + "`" +
			huntTime,
		Inline: true,
	}

	// Embed inventory

	infoInventory := &discordgo.MessageEmbedField{
		Name: "**Inventory :**",
		Value: "Characters : **" + strconv.Itoa(p.CharAmount()) + "/" + strconv.Itoa(p.MaxChar()) + "**.\n" +
			"Classic boxes : **" + strconv.Itoa(p.Boxes.Boxes) + "**.\n" +
			"Rare boxes : **" + strconv.Itoa(p.Boxes.RareBoxes) + "**.\n" +
			"Limited boxes : **" + strconv.Itoa(charLimited) + "**.\n" +
			"Special boxes : **" + strconv.Itoa(len(p.SpecialBoxes)) + "**.",
		Inline: true,
	}

	// Embed pack

	infoPack := &discordgo.MessageEmbedField{
		Name: "**Pack :**",
		Value: "Grimms : **" + strconv.Itoa(p.GrimmAmount()) + "/" + strconv.Itoa(p.MaxChar()) + "**.\n" +
			"Classic **G**boxes : **" + strconv.Itoa(p.Boxes.GrimmBoxes) + "**.\n" +
			"Rare **G**boxes : **" + strconv.Itoa(p.Boxes.GrimmBoxes) + "**.\n" +
			"Limited **G**boxes : **" + strconv.Itoa(grimmLimited) + "**.",
		Inline: true,
	}
	infoBoosts := &discordgo.MessageEmbedField{
		Name: "**Boosts :**",
		Value: "`XP Boost : " + BoolStringFormat(p.Shop.XPBoost) + "`\n" +
			"`Luck Boost : " + BoolStringFormat(p.Shop.LuckBoost) + luckBoostString + "`",
		Inline: true,
	}
	infoMore := &discordgo.MessageEmbedField{
		Name: "**More :**",
		Value: "More infos [here](https://me.rwbyadventures.com/" + ctx.Author.ID + ").\n" +
			"More settings [here](https://settings.rwbyadventures.com/).",
		Inline: true,
	}
	infoReport := &discordgo.MessageEmbedField{
		Name: "**Other :**",
		Value: "Found a bug or need help ?\n" +
			"`r!report I lost my character :c`",
		Inline: true,
	}

	InfoEmbed := &discordgo.MessageEmbed{
		Title: "These are your infos " + ctx.Author.Username + ".",
		Color: config.Botcolor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Author.AvatarURL("512"),
		},
		Footer: discord.DefaultFooter,
		Fields: []*discordgo.MessageEmbedField{infoProfile, infoInventory, infoPack, infoMission, infoBoosts, infoMore, infoCoolDown, infoReport},
	}

	ctx.Reply(discord.ReplyParams{
		Content: InfoEmbed,
	})
}
