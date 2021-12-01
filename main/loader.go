package main

import (
	"rwby-adventures/main/commands"
	"rwby-adventures/main/discord"

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
		Name:        "profile",
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
	addchar := &discord.Command{
		Name:        "addchar",
		Description: "Test.",
		Aliases:     discord.CmdAlias{"ac"},
		Menu:        discord.GeneralMenu,
		Call:        commands.Addchar,
	}
	info := &discord.Command{
		Name:        "info",
		Description: "Test.",
		Menu:        discord.GeneralMenu,
		Call:        commands.Info,
	}
	discord.AddCmd(help)
	discord.AddCmd(me)
	discord.AddCmd(del)
	discord.AddCmd(addchar)
	discord.AddCmd(info)

	discord.MakeEmbed()
	discord.LoadAllCharacters()

	discord.CommandRouter.LoadSlashCommands([]*discordgo.Session{discord.Session})
}
