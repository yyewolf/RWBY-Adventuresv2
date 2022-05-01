package commmands_roleplay

import (
	"fmt"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

var RPCommand = &discord.Command{
	Name:        "rp",
	Description: "Roleplay as a persona.",
	Menu:        discord.RoleplayMenu,
	Args: []discord.Arg{
		{
			Name:        "message",
			Description: "Message you want to send.",
			Required:    true,
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "image_url",
			Description: "If you want to send an attachment.",
			Type:        discordgo.ApplicationCommandOptionString,
		},
		{
			Name:        "persona_id",
			Description: "If you want to use another persona (defaults to your selected).",
			Type:        discordgo.ApplicationCommandOptionString,
		},
	},
	Call: rp,
}

func rp(ctx *discord.CmdContext) {
	persona := ctx.Arguments.GetArg("persona_id", 2, "")
	var char *models.Character
	var grimm *models.Grimm
	var t = -1
	for _, c := range ctx.Player.Characters {
		if c.CharID == fmt.Sprint(persona.Value) {
			char = c
			t = models.CharType
			break
		}
	}
	for _, g := range ctx.Player.Grimms {
		if g.GrimmID == fmt.Sprint(persona.Value) {
			grimm = g
			t = models.GrimmType
			break
		}
	}

	if t == -1 {
		char = ctx.Player.SelectedChar
		grimm = ctx.Player.SelectedGrimm
		t = ctx.Player.SelectedType
	}

	if t == models.CharType {
		rpChar(ctx, char)
	} else {
		rpGrimm(ctx, grimm)
	}
}

func rpChar(ctx *discord.CmdContext, char *models.Character) {
	if char == nil {
		return
	}
	msg := ctx.Arguments.GetArg("message", 0, "")
	image_url := ctx.Arguments.GetArg("image_url", 1, "")
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s says :", char.Name),
			Description: fmt.Sprint(msg.Value),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: char.ToRealChar().IconURL,
			},
			Image: &discordgo.MessageEmbedImage{
				URL: fmt.Sprint(image_url.Value),
			},
			Color: char.RarityToColor(),
			Footer: &discordgo.MessageEmbedFooter{
				IconURL: ctx.Author.AvatarURL("512"),
				Text:    ctx.Author.Username,
			},
		},
	})
}

func rpGrimm(ctx *discord.CmdContext, grimm *models.Grimm) {
	if grimm == nil {
		return
	}
	msg := ctx.Arguments.GetArg("message", 0, "")
	image_url := ctx.Arguments.GetArg("image_url", 1, "")
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s says :", grimm.Name),
			Description: fmt.Sprint(msg.Value),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: grimm.ToRealGrimm().IconURL,
			},
			Image: &discordgo.MessageEmbedImage{
				URL: fmt.Sprint(image_url.Value),
			},
			Color: grimm.RarityToColor(),
			Footer: &discordgo.MessageEmbedFooter{
				IconURL: ctx.Author.AvatarURL("512"),
				Text:    ctx.Author.Username,
			},
		},
	})
}
