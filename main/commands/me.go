package commands

import (
	"fmt"
	"rwby-adventures/discord"
)

func Me(ctx *discord.CmdContext) {
	ctx.Reply(discord.ReplyParams{
		Content: fmt.Sprintf("Hello, are you new : %v", ctx.Player.IsNew),
		ID:      ctx.ID,
		Edit:    ctx.IsComponent,
	})
}
