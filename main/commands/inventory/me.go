package commands_inventory

import (
	"fmt"
	"math"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"time"

	"github.com/bwmarrin/discordgo"
)

var PlayerInfoCommand = &discord.Command{
	Name:        "my",
	Description: "View your RWBY Adventures profile.",
	Aliases:     discord.CmdAlias{"me"},
	Menu:        discord.GeneralMenu,
	Call:        Me,
}

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
	var resetStringBox = fmt.Sprintf("✅ - `Loots` (**%d/%d**)", p.LastBoxes.Amount, p.Maxlootbox)
	var resetStringGamble = fmt.Sprintf("✅ - `Gamble` (**%d/3**)", p.Gamble.Amount)
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
		resetStringBox = fmt.Sprintf("🕓 - `Loots` (%s)", TimeLeftString(p.LastBoxes.Time, 86400))
	}
	canGamble, _ := p.CanGamble()
	if !canGamble {
		resetStringGamble = fmt.Sprintf("🕓 - `Gamble` (%s)", TimeLeftString(p.Gamble.Time, 86400))
	}
	canDungeon := p.CanDungeon()
	if !canDungeon {
		resetStringDungeon = fmt.Sprintf("🕓 - `Dungeon` (%s)", TimeLeftString(p.Status.LastDungeon, 18000)) //5*60*60
	}
	if p.Missions.IsInMission {
		days := int(math.Ceil(float64(p.Missions.MissionMsgLeft) / 24))
		missionTime = fmt.Sprintf("`Time left : %d days.`\n", days)
	}
	if p.Missions.IsInHunt {
		days := int(math.Ceil(float64(p.Missions.HuntMsgLeft) / 24))
		huntTime = fmt.Sprintf("\n`Time left : %d days.`", days)
	}
	if p.Shop.LuckBoost {
		luckBoostString = fmt.Sprintf(" (%d/20)", p.Shop.LuckBoostTime)
	}

	// Embed Profile

	infoProfile := &discordgo.MessageEmbedField{
		Name: "**Player :**",
		Value: fmt.Sprintf(`Level **%d**.
			CP : **%d**/**%d**.
			Slots : **%d**.
			Boxes : **%d**/**%d**.
			Liens : **%dⱠ**.`,
			p.Level, p.CP, p.MaxCP, p.MaxChar(), p.Lootbox(), p.MaxChar(), p.Balance,
		),
		Inline: true,
	}

	// Embed cool downs

	infoCoolDown := &discordgo.MessageEmbedField{
		Name: "**Cooldowns :**",
		Value: fmt.Sprintf(`%s
		%s
		%s`,
			resetStringBox, resetStringGamble, resetStringDungeon,
		),
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
		Value: fmt.Sprintf(`Characters : **%d/%d**.
			Classic boxes : **%d**.
			Rare boxes : **%d**.
			Limited boxes : **%d**.
			Special boxes : **%d**.`,
			p.CharAmount(), p.MaxChar(), p.Boxes.Boxes, p.Boxes.RareBoxes, charLimited, len(p.SpecialBoxes),
		),
		Inline: true,
	}

	// Embed pack

	infoPack := &discordgo.MessageEmbedField{
		Name: "**Pack :**",
		Value: fmt.Sprintf(`Grimms : **%d/%d**.
			Classic **G**boxes : **%d**.
			Rare **G**boxes : **%d**.
			Limited **G**boxes : **%d**.`,
			p.GrimmAmount(), p.MaxChar(), p.Boxes.GrimmBoxes, p.Boxes.RareGrimmBoxes, grimmLimited,
		),
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
		Value: fmt.Sprintf(`More infos [here](https://me.rwbyadventures.com/%s).
			More settings [here](https://settings.rwbyadventures.com/).`,
			ctx.Author.ID,
		),
		Inline: true,
	}
	infoReport := &discordgo.MessageEmbedField{
		Name: "**Other :**",
		Value: "Found a bug or need help ?\n" +
			"`/report I lost my character :c`",
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
