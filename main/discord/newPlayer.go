package discord

import (
	"rwby-adventures/config"

	"github.com/bwmarrin/discordgo"
)

func HandleNewPlayer(ctx *CmdContext) {
	// Useful stuff
	g := ctx.Guild
	p := ctx.Player

	WelcomeFieldBattles := &discordgo.MessageEmbedField{
		Name:  "**Battles :**",
		Value: "You are able to duel your friends, but do not worry, your characters will never really faint.",
	}
	WelcomeFieldLootboxes := &discordgo.MessageEmbedField{
		Name:  "**Lootboxes :**",
		Value: "To find new characters, you will need to open lootboxes, and in order to do so you will be able to find once in a while a lootbox.\nYou received a free one by the way, come on, try it !\nType `" + g.Prefix + "open` to open it.",
	}
	WelcomeFieldCharacters := &discordgo.MessageEmbedField{
		Name:  "**Characters :**",
		Value: "Not all of the characters from RWBY are implemented yet, but they will eventually reach out and be available in lootboxes !",
	}
	WelcomeFieldInfos := &discordgo.MessageEmbedField{
		Name:  "**More Infos :**",
		Value: "We suggest that you check out our FAQ [here](https://rwbyadventures.com/faq/)",
	}
	WelcomeFieldDisclaimer := &discordgo.MessageEmbedField{
		Name:  "**Disclaimer :**",
		Value: "*This is not endorsed by Rooster Teeth in any way. Views, opinions, and thoughts are all my own. Rooster Teeth and RWBY are trade names or registered trademarks of Rooster Teeth Productions, LLC. © Rooster Teeth Productions, LLC.*",
	}

	WelcomeEmbed := &discordgo.MessageEmbed{
		Title:       "Welcome to RWBY Adventures, " + ctx.Author.Username + " !",
		Color:       config.Botcolor,
		Description: "By typing one of my command, you just began your new adventure, you will need to search ; train ; battle with your characters in order to win more fights. \nType `r!help` to see the help menu.\nType `r!tutorial` to see the tutorial.",
		Footer:      DefaultFooter,
		Fields:      []*discordgo.MessageEmbedField{WelcomeFieldLootboxes, WelcomeFieldBattles, WelcomeFieldCharacters, WelcomeFieldInfos, WelcomeFieldDisclaimer},
	}

	ctx.Reply(ReplyParams{
		Content: WelcomeEmbed,
	})
	p.IsNew = false
	p.Badges = 0
	p.Settings = config.CanDMPlayer
	config.Database.Create(p)
}
