package admin

import (
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"time"
)

var UpdateCommand = &discord.Command{
	Name:        "update",
	Description: "Triggers the update message, don't forget to modify the embed in the code.",
	Restricted:  true,
	Call:        Update,
}

func Update(ctx *discord.CmdContext) {
	if ctx.Player.DiscordID != "144472011924570113" {
		ctx.Reply(discord.ReplyParams{
			Content:   "You are not allowed to use this command.",
			Ephemeral: true,
		})
		return
	}

	u := &models.Update{
		ID:   1,
		Time: time.Now(),
	}

	u.Save()

	c := u.GetMessage()

	ctx.Reply(discord.ReplyParams{
		Content:   "Enabled update, here's a sneak peak:",
		Ephemeral: true,
	})
	ctx.Reply(discord.ReplyParams{
		Content:   c,
		FollowUp:  true,
		Ephemeral: true,
	})
}

func PassiveUpdate(ctx *discord.CmdContext) {
	u := models.GetLastUpdate()
	if u == nil {
		return
	}
	// We don't want to send to dm channels
	if ctx.GuildID == "" {
		return
	}
	if ctx.Guild.LastUpdateMessage.Sub(u.Time) < 0 {
		c := u.GetMessage()
		ctx.Guild.LastUpdateMessage = u.Time
		ctx.Guild.Save()
		ctx.Reply(discord.ReplyParams{
			Content:   c,
			Automated: true,
		})
	}
}
