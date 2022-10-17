package admin

import (
	"fmt"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

func AddBox(ctx *discord.CmdContext) {
	t := ctx.Arguments.GetArg("type", 0, 0)
	a := ctx.Arguments.GetArg("amount", 1, 0)
	target := ctx.Arguments.GetArg("user", 2, nil)

	user := target.Raw.UserValue(ctx.Session)
	player := models.GetPlayer(user.ID)

	amount := int(a.Raw.IntValue())

	switch t.Raw.IntValue() {
	case 0:
		player.Boxes.Boxes += amount
	case 1:
		player.Boxes.RareBoxes += amount
	case 2:
		player.Boxes.GrimmBoxes += amount
	case 3:
		player.Boxes.RareGrimmBoxes += amount
	}

	player.Boxes.Save()

	ctx.Reply(discord.ReplyParams{
		Content: fmt.Sprintf("Added %d boxes to %s", amount, user.Username),
	})
}
