package main

import (
	"rwby-adventures/main/commands"
	commands_inventory "rwby-adventures/main/commands/inventory"
	commands_temporary "rwby-adventures/main/commands/temporary"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

func StartDiscord() {
	discord.Start()

	addchar := &discord.Command{
		Name:        "addchar",
		Description: "Test.",
		Aliases:     discord.CmdAlias{"ac"},
		Menu:        discord.GeneralMenu,
		Call:        commands_temporary.Addchar,
	}

	help := &discord.Command{
		Name:        "help",
		Description: "View the help menu.",
		Aliases:     discord.CmdAlias{"h"},
		Menu:        discord.GeneralMenu,
		Call:        commands.Help,
	}
	me := &discord.Command{
		Name:        "profile",
		Description: "View your infos.",
		Aliases:     discord.CmdAlias{"me"},
		Menu:        discord.GeneralMenu,
		Call:        commands_inventory.Me,
	}
	del := &discord.Command{
		Name:        "delete",
		Description: "Delete all your informations.",
		Aliases:     discord.CmdAlias{"del"},
		Menu:        discord.GeneralMenu,
		Call:        commands_temporary.Delete,
	}
	info := &discord.Command{
		Name:        "info",
		Description: "Test.",
		Menu:        discord.GeneralMenu,
		Call:        commands_inventory.Info,
		Args: []discord.Arg{
			{
				Name:        "id",
				Description: "Identification number of your persona.",
				Size:        1,
				Required:    false,
				Type:        discordgo.ApplicationCommandOptionString,
			},
			{
				Name:        "latest",
				Description: "Whether or not you want to view the infos of your latest persona.",
				Size:        1,
				Required:    false,
				Type:        discordgo.ApplicationCommandOptionBoolean,
			},
		},
	}
	discord.AddCmd(help)
	discord.AddCmd(me)
	discord.AddCmd(del)
	discord.AddCmd(addchar)
	discord.AddCmd(info)

	discord.MakeEmbed()

	discord.CommandRouter.LoadSlashCommands([]*discordgo.Session{discord.Session})
}
