package main

import (
	"rwby-adventures/commands"
	"rwby-adventures/discord"

	"github.com/bwmarrin/discordgo"
)

func StartDiscord() {
	discord.Start()

	help := &discord.Command{
		Name:        "help",
		Description: "View the help menu.",
		Aliases:     discord.CmdAlias{"h"},
		Menu:        discord.GeneralMenu,
		Call:        commands.Help,
	}
	me := &discord.Command{
		Name:        "info",
		Description: "View your infos.",
		Aliases:     discord.CmdAlias{"me"},
		Menu:        discord.GeneralMenu,
		Call:        commands.Me,
	}
	del := &discord.Command{
		Name:        "delete",
		Description: "Delete all your informations.",
		Aliases:     discord.CmdAlias{"del"},
		Menu:        discord.GeneralMenu,
		Call:        commands.Delete,
	}
	discord.AddCmd(help)
	discord.AddCmd(me)
	discord.AddCmd(del)

	discord.MakeEmbed()

	discord.CommandRouter.LoadSlashCommands([]*discordgo.Session{discord.Session})
}
