package commands_temporary

import (
	"rwby-adventures/main/discord"
)

func AddBox(ctx *discord.CmdContext) {
	arg := ctx.Arguments.GetArg("type", 0, 0)

	switch arg.Raw.IntValue() {
	case 0:
		ctx.Player.Boxes.Boxes += 5
	case 1:
		ctx.Player.Boxes.RareBoxes += 5
	case 2:
		ctx.Player.Boxes.GrimmBoxes += 5
	case 3:
		ctx.Player.Boxes.RareGrimmBoxes += 5
	}

	ctx.Player.Boxes.Save()

	ctx.Reply(discord.ReplyParams{
		Content: "Gave you 5 boxes of your choice.",
	})
}
