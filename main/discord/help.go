package discord

import (
	"rwby-adventures/config"
	"sort"

	"github.com/bwmarrin/discordgo"
)

type menuName string

const (
	GeneralMenu   menuName = "General"
	InventoryMenu menuName = "Inventory"
)

func menuEmoji(name string) string {
	switch name {
	case string(GeneralMenu):
		return "🖥️"
	case string(InventoryMenu):
		return "🛍️"
	}
	return ""
}

var HelpMenus map[string][]*Command
var MenuEmbed map[string]*discordgo.MessageEmbed

func MakeEmbed() {
	MenuEmbed = make(map[string]*discordgo.MessageEmbed)
	for menuName, cmds := range HelpMenus {
		embed := &discordgo.MessageEmbed{
			Title: menuName + " commands :",
			Color: config.Botcolor,
		}
		for i, cmd := range cmds {
			if i != 0 {
				embed.Description += "\n"
			}
			embed.Description += "`" + CommandRouter.Prefix + cmd.HelpName + " : " + cmd.Description + "`"
		}
		MenuEmbed[menuName] = embed
	}
}

func HelpComponent(menuID string, defaultMenu string) []discordgo.MessageComponent {
	var opts []discordgo.SelectMenuOption

	var keys []string
	for k := range MenuEmbed {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		if name == "" {
			continue
		}
		opt := discordgo.SelectMenuOption{
			Label: name,
			Value: name,
			Emoji: discordgo.ComponentEmoji{
				Name: menuEmoji(name),
			},
		}
		if name == defaultMenu {
			opt.Default = true
		}
		opts = append(opts, opt)
	}

	cmp := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.SelectMenu{
					CustomID: menuID,
					Options:  opts,
				},
			},
		},
	}

	return cmp
}
