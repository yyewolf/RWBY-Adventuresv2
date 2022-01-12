package commands_inventory

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var InventoryCommand = &discord.Command{
	Name:        "inventory",
	Description: "View informations about a certain persona.",
	Menu:        discord.PersonasMenu,
	Call:        Info,
	Aliases: discord.CmdAlias{
		"inv",
	},
	Args: []discord.Arg{},
}

type inventoryMenuData struct {
	UserID string
	Page   int
}

func Inventory(ctx *discord.CmdContext) {
	reply := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s's inventory", ctx.Author.Username),
		Description: fmt.Sprintf("To select a persona type `%sselect <PersonaID>`.", ctx.Guild.Prefix),
		Color:       config.Botcolor,
		Footer:      discord.DefaultFooter,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Author.AvatarURL("512"),
		},
	}

}
