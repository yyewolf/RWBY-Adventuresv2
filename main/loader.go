package main

import (
	"os"
	"os/signal"
	"rwby-adventures/commands"
	"rwby-adventures/discord"
	"syscall"

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
	discord.AddCmd(help)
	discord.AddCmd(me)

	discord.MakeEmbed()

	discord.CommandRouter.LoadSlashCommands([]*discordgo.Session{discord.Session})

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Session.Close()
}
