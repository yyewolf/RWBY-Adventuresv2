package commands

import (
	"fmt"
	"rwby-adventures/discord"
)

func Me(ctx *discord.CmdContext) {
	ctx.Reply(discord.ReplyParams{
		Content: fmt.Sprintf("Hello, are you new : %v\nThis is your time : %d", ctx.Player.IsNew, ctx.Player.Shop.LuckBoostTime),
		ID:      ctx.ID,
		Edit:    ctx.IsComponent,
	})
}
