package commands_inventory

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var BalanceCommand = &discord.Command{
	Name:        "balance",
	Description: "Check out your balance.",
	Aliases:     discord.CmdAlias{"bal"},
	Menu:        discord.GeneralMenu,
	Call:        Balance,
}

func Balance(ctx *discord.CmdContext) {
	biddedMoney := ctx.Player.Balance - ctx.Player.TotalBalance()
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s's balance", ctx.Author.Username),
		Description: fmt.Sprintf("You currently have : %dⱠ (Lien).\nYou also have **%d** Arms and **%d** Minions.", ctx.Player.Balance, ctx.Player.Arms, ctx.Player.Minions),
		Color:       config.Botcolor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Author.AvatarURL("512"),
		},
		Footer: discord.DefaultFooter,
	}
	if biddedMoney > 0 {
		embed.Description += fmt.Sprintf("\nYou bidded : %dⱠ.", biddedMoney)
	}
	ctx.Reply(discord.ReplyParams{
		Content: embed,
	})
}
