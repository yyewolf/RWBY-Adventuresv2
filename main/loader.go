package main

import (
	"rwby-adventures/main/commands"
	commands_admin "rwby-adventures/main/commands/admin"
	commands_badges "rwby-adventures/main/commands/badges"
	commands_boxes "rwby-adventures/main/commands/boxes"
	commands_duel "rwby-adventures/main/commands/duels"
	commands_dungeon "rwby-adventures/main/commands/dungeons"
	commands_gamble "rwby-adventures/main/commands/gambles"
	commands_inventory "rwby-adventures/main/commands/inventory"
	commands_market "rwby-adventures/main/commands/market"
	commands_misc "rwby-adventures/main/commands/misc"
	commands_missions "rwby-adventures/main/commands/missions"
	commands_roleplay "rwby-adventures/main/commands/roleplay"
	commands_settings "rwby-adventures/main/commands/settings"
	commands_shop "rwby-adventures/main/commands/shop"
	commands_trade "rwby-adventures/main/commands/trade"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

func StartDiscord() {
	discord.Start()

	//discord.AddCmd(commands_temporary.TemporaryCommand)
	discord.AddCmd(commands.HelpCommand)
	discord.AddCmd(commands_inventory.PlayerInfoCommand)
	discord.AddCmd(commands_inventory.InfoCommand)
	discord.AddCmd(commands_inventory.RemoveCommand)
	discord.AddCmd(commands_inventory.SelectCommand)
	discord.AddCmd(commands_inventory.FavoritesCommand)
	discord.AddCmd(commands_inventory.InventoryCommand)
	discord.AddCmd(commands_inventory.PackCommand)
	discord.AddCmd(commands_inventory.BalanceCommand)
	discord.AddCmd(commands_inventory.JarCommand)
	discord.AddCmd(commands_inventory.CompCommand)
	discord.AddCmd(commands_boxes.BoxesCommand)
	discord.AddCmd(commands_trade.TradesCommand)
	discord.AddCmd(commands_roleplay.RPCommand)
	discord.AddCmd(commands_roleplay.RollCommand)
	discord.AddCmd(commands_missions.MissionCommand)
	discord.AddCmd(commands_missions.HuntCommand)
	discord.AddCmd(commands_dungeon.DungeonCommand)
	discord.AddCmd(commands_settings.ReportCommand)
	discord.AddCmd(commands_settings.GuildCommand)
	discord.AddCmd(commands_settings.OrderCommand)
	discord.AddCmd(commands_misc.InviteCommand)
	discord.AddCmd(commands_misc.Eventcommand)
	discord.AddCmd(commands_misc.DailyCommand)
	discord.AddCmd(commands_misc.StatCommand)
	discord.AddCmd(commands_market.MarketCommand)
	discord.AddCmd(commands_duel.DuelCommand)
	discord.AddCmd(commands_badges.BadgesCommand)
	discord.AddCmd(commands_shop.ShopCommand)
	discord.AddCmd(commands_gamble.GambleCommand)

	discord.AddCmd(commands_admin.AdminCommands)

	discord.CommandRouter.LoadSlashCommands([]*discordgo.Session{discord.Session})
}
