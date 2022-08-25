package commands_badges

import (
	"fmt"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

func BadgesGive(ctx *discord.CmdContext) {
	if ctx.Author.ID != "144472011924570113" {
		ctx.Reply(discord.ReplyParams{
			Content: "You can't give badges to other people.",
		})
		return
	}

	arg := ctx.Arguments.GetArg("badge", 0, "")
	if !arg.Found {
		ctx.Reply(discord.ReplyParams{
			Content: "You need to specify an opponent.",
		})
		return
	}
	badgeName := arg.Value.(string)

	arg = ctx.Arguments.GetArg("player", 1, "")
	if !arg.Found {
		ctx.Reply(discord.ReplyParams{
			Content: "You need to specify an opponent.",
		})
		return
	}
	playerID := arg.Value.(string)

	playerBadge := &models.PlayerBadges{
		DiscordID: playerID,
	}

	for _, b := range models.DefaultBadges {
		if b.Name == badgeName {
			playerBadge.BadgeID = b.BadgeID
			break
		}
	}

	playerBadge.Save()

	ctx.Reply(discord.ReplyParams{
		Content:   fmt.Sprintf("You gave %s the **%s** badge.", playerID, badgeName),
		Ephemeral: true,
	})
}
