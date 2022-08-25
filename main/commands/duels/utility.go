package commands_duels

import (
	"bytes"
	"fmt"
	"rwby-adventures/main/static"
	"strconv"

	"image"
	_ "image/jpeg"
	"image/png"

	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
)

func duelEmbed(duel *BattleStruct, Player int, Opponent int) *discordgo.MessageEmbed {
	EmbedFields := []*discordgo.MessageEmbedField{}
	Character := duel.Chars[Player]
	for i := range Character.atks {
		CurrentAtk := &discordgo.MessageEmbedField{
			Name:   "**" + Character.atks[i].Name + "**",
			Value:  "```APACHE\n",
			Inline: true,
		}
		if Character.atks[i].Damages != 0 {
			CurrentAtk.Value += "Damages : " + strconv.Itoa(Character.atks[i].Damages) + "%.\n"
		}
		if Character.atks[i].Heal != 0 {
			CurrentAtk.Value += "Heal : " + strconv.Itoa(Character.atks[i].Heal) + "%.\n"
		}
		if Character.atks[i].StunChance != 0 {
			CurrentAtk.Value += "StunChances : " + strconv.Itoa(Character.atks[i].StunChance) + "%.\n"
		}
		if Character.atks[i].Speed != 0 {
			CurrentAtk.Value += "Speed : " + strconv.Itoa(Character.atks[i].Speed) + "%.\n"
		}
		CurrentAtk.Value += "Cooldown : " + strconv.Itoa(Character.atks[i].Every) + " turns.\n"
		CurrentAtk.Value += "Reaction : " + strconv.Itoa(i+1) + "."
		CurrentAtk.Value += "```"
		EmbedFields = append(EmbedFields, CurrentAtk)
	}
	if duel.Chars[Player].priority > 0 {
		CurrentAtk := &discordgo.MessageEmbedField{
			Name:   "**Semblance**",
			Value:  "```APACHE\n",
			Inline: true,
		}
		CurrentAtk.Value += "Description : " + duel.Chars[Player].semblance.Desc + "\n"
		CurrentAtk.Value += "Cooldown : " + strconv.Itoa(duel.Chars[Player].semblance.Every) + " turns.\n"
		CurrentAtk.Value += "Reaction : S."
		CurrentAtk.Value += "```"
		EmbedFields = append(EmbedFields, CurrentAtk)
	}
	Embed := &discordgo.MessageEmbed{
		Title:       "Duel against " + duel.Players[Opponent].User.Username,
		Description: "You have multiple attacks that you can use.",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: Character.iconURL,
		},
		Fields: EmbedFields,
		Color:  Character.RarityToColor(),
	}
	return Embed
}

func createDuelImage(Personas ...*BattlePersona) *bytes.Reader {
	d, err := static.DatabaseFS.ReadFile("database/images/VS.png")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	backgroundImage, _, err := image.Decode(bytes.NewBuffer(d))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	positions := [][]int{{40, 40}, {500, 40}}

	battleImage := gg.NewContext(426*2, 240*2)
	battleImage.DrawImage(backgroundImage, 0, 0)

	for i, p := range Personas {
		d, err = static.DatabaseFS.ReadFile(p.imagePath)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		image, _, err := image.Decode(bytes.NewBuffer(d))
		if err != nil {
			fmt.Println(err)
			return nil
		}

		battleImage.DrawImage(image, positions[i][0], positions[i][1])
	}

	battleImage.Clip()

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, battleImage.Image())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bytes.NewReader(buffer.Bytes())
}
