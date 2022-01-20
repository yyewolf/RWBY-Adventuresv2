package commands_trade

import (
	"fmt"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

func RefuseTrade(ctx *discord.CmdContext) {
	arg := ctx.Arguments.GetArg("id", 0, "")
	if !arg.Found {
		ctx.Reply(discord.ReplyParams{
			Content: "You need to put a Trade ID.",
		})
		return
	}
	id := arg.Value.(string)
	trade, err := models.GetTrade(id)
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content: "Trade ID not found.",
		})
		return
	}
	if trade.ReceiverID != ctx.Author.ID {
		ctx.Reply(discord.ReplyParams{
			Content: "Trade ID not found.",
		})
		return
	}
	SenderDM, _ := discord.Session.UserChannelCreate(trade.SenderID)

	trade.Delete()
	ctx.Reply(discord.ReplyParams{
		Content:   fmt.Sprintf("You have refused trade `%s`.", trade.ID),
		Ephemeral: true,
	})
	discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("Trade `%s` has been refused by the user.", trade.ID))
}
