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
	del := &discord.Command{
		Name:        "delete",
		Description: "Delete all your informations.",
		Aliases:     discord.CmdAlias{"del"},
		Menu:        discord.GeneralMenu,
		Call:        commands_temporary.Delete,
	}
	discord.AddCmd(addchar)
	discord.AddCmd(del)
	discord.AddCmd(commands.HelpCommand)
	discord.AddCmd(commands_inventory.PlayerInfoCommand)
	discord.AddCmd(commands_inventory.InfoCommand)
	discord.AddCmd(commands_inventory.RemoveCommand)
	discord.AddCmd(commands_inventory.SelectCommand)
	discord.AddCmd(commands_inventory.FavoritesCommand)
	discord.AddCmd(commands_inventory.InventoryCommand)
	discord.AddCmd(commands_inventory.PackCommand)
	discord.AddCmd(commands_inventory.BalanceCommand)

	discord.MakeEmbed()

	discord.CommandRouter.LoadSlashCommands([]*discordgo.Session{discord.Session})
}
