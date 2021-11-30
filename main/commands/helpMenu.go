package commands

import (
	"rwby-adventures/discord"

	"github.com/bwmarrin/discordgo"
)

func Help(ctx *discord.CmdContext) {
	menu := string(discord.GeneralMenu)
	menuID := ctx.ID
	if ctx.IsComponent {
		if ctx.Author.ID != ctx.Menu.SourceContext.Author.ID {
			return
		}
		ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredMessageUpdate,
		})
		menu = ctx.ComponentData.Values[0]
		menuID = ctx.Menu.MenuID
	}
	complex := &discordgo.MessageSend{
		Embed:      discord.MenuEmbed[menu],
		Components: discord.HelpComponent(menuID, menu),
	}
	complex.Embed.Footer = discord.DefaultFooter

	if ctx.IsComponent {
		// Keep old context if a button is pressed
		ctx.Menu.SourceContext.ID = ctx.Message.ID
		ctx = ctx.Menu.SourceContext
		ctx.IsComponent = true
	} else {
		discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
			MenuID:        ctx.ID,
			SourceContext: ctx,
			Call:          Help,
		}, 0)
	}

	ctx.Reply(discord.ReplyParams{
		Content: complex,
		ID:      ctx.ID,
		Edit:    ctx.IsComponent,
	})
}
