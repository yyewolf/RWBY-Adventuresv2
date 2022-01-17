package commands_inventory

import (
	"rwby-adventures/main/discord"
)

func Favorites(ctx *discord.CmdContext) {
	ctx.Arguments = append(ctx.Arguments, &discord.CommandArg{
		Name:  "favorites",
		Value: true,
		Found: true,
	})

	Inventory(ctx)
}
