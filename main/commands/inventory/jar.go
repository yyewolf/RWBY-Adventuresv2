package commands_inventory

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var JarCommand = &discord.Command{
	Name:        "jar",
	Description: "Use and view your XP Jar.",
	SubCommands: []*discord.Command{
		{
			Name:        "view",
			Description: "View your XP Jar.",
			Menu:        discord.InventoryMenu,
			Call:        JarView,
		},
		{
			Name:        "use",
			Description: "Use your XP Jar.",
			Menu:        discord.InventoryMenu,
			Args: []discord.Arg{
				{
					Name:        "amount",
					Description: "Amount of XP to use.",
					Size:        1,
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
			},
			Call: JarUse,
		},
	},
}

func JarView(ctx *discord.CmdContext) {
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s's XP Jar", ctx.Author.Username),
			Description: fmt.Sprintf("You have **%d** XP left in your jar!\nUse </jar use:%s> to use it on your selected character.", ctx.Player.Jar, ctx.Command.ID),
			Color:       config.Botcolor,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://vignette.wikia.nocookie.net/rwby/images/6/62/Jinn_infobox.png/revision/latest/top-crop/width/320/height/320?cb=20181118074021",
			},
			Footer: discord.DefaultFooter,
		},
	})
}

func JarUse(ctx *discord.CmdContext) {
	arg := ctx.Arguments.GetArg("amount", 0, 0)
	amount := arg.Raw.IntValue()
	if amount > ctx.Player.Jar {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have that much XP in your jar.",
			Ephemeral: true,
		})
		return
	}
	if amount < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You can't use a negative amount of XP.",
			Ephemeral: true,
		})
		return
	}
	ctx.Player.Jar -= amount
	ctx.GiveSelectionXP(amount, false)

	ctx.Player.Save()
	ctx.Reply(discord.ReplyParams{
		Content:   fmt.Sprintf("You have used **%d** XP from your jar.", amount),
		Ephemeral: true,
	})
}
