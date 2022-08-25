package main

import (
	"rwby-adventures/main/commands"
	commands_badges "rwby-adventures/main/commands/badges"
	commands_boxes "rwby-adventures/main/commands/boxes"
	commands_duel "rwby-adventures/main/commands/duels"
	commands_dungeon "rwby-adventures/main/commands/dungeons"
	commands_inventory "rwby-adventures/main/commands/inventory"
	commands_market "rwby-adventures/main/commands/market"
	commands_misc "rwby-adventures/main/commands/misc"
	commands_missions "rwby-adventures/main/commands/missions"
	commands_roleplay "rwby-adventures/main/commands/roleplay"
	commands_settings "rwby-adventures/main/commands/settings"
	commands_temporary "rwby-adventures/main/commands/temporary"
	commands_trade "rwby-adventures/main/commands/trade"
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
	discord.AddCmd(commands_temporary.ArenaCommand)
	discord.AddCmd(commands.HelpCommand)
	discord.AddCmd(commands_inventory.PlayerInfoCommand)
	discord.AddCmd(commands_inventory.InfoCommand)
	discord.AddCmd(commands_inventory.RemoveCommand)
	discord.AddCmd(commands_inventory.SelectCommand)
	discord.AddCmd(commands_inventory.FavoritesCommand)
	discord.AddCmd(commands_inventory.InventoryCommand)
	discord.AddCmd(commands_inventory.PackCommand)
	discord.AddCmd(commands_inventory.BalanceCommand)
	discord.AddCmd(commands_boxes.BoxesCommand)
	discord.AddCmd(commands_trade.TradesCommand)
	discord.AddCmd(commands_roleplay.RPCommand)
	discord.AddCmd(commands_roleplay.RollCommand)
	discord.AddCmd(commands_missions.MissionCommand)
	discord.AddCmd(commands_missions.HuntCommand)
	discord.AddCmd(commands_dungeon.DungeonCommand)
	discord.AddCmd(commands_settings.ReportCommand)
	discord.AddCmd(commands_misc.InviteCommand)
	discord.AddCmd(commands_misc.Eventcommand)
	discord.AddCmd(commands_market.MarketCommand)
	discord.AddCmd(commands_duel.DuelCommand)
	discord.AddCmd(commands_badges.BadgesCommand)

	discord.MakeEmbed()

	discord.CommandRouter.LoadSlashCommands([]*discordgo.Session{discord.Session})
}
