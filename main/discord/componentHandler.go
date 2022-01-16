package discord

import (
	"rwby-adventures/models"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func routeComponents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}
	split := strings.Split(i.MessageComponentData().CustomID, "-")
	val, found := ActiveMenus.Get(split[0])
	if !found {
		return
	}
	m := val.(*Menus)
	ctx := &CmdContext{
		Session:   s,
		ID:        i.ID,
		ChannelID: i.ChannelID,
		GuildID:   i.GuildID,

		IsComponent:   true,
		IsInteraction: true,
		Menu:          m,
		ComponentData: i.MessageComponentData(),
		Interaction:   i.Interaction,
		Message:       i.Interaction.Message,
	}
	if i.Member != nil {
		ctx.Author = i.Member.User
	} else {
		ctx.Author = i.User
	}

	ctx.Player = models.GetPlayer(ctx.Author.ID)
	ctx.Guild = models.GetGuild(ctx.GuildID)

	m.MenuID = split[0]

	s.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	m.Call(ctx)
}
