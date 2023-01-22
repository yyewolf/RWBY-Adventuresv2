package commands_settings

import (
	"errors"
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var GuildCommand = &discord.Command{
	Name:        "config",
	Description: "To configure your server.",
	SubCommands: []*discord.Command{
		{
			Name:        "channel",
			Description: "To configure the guild channel.",
			Menu:        discord.ConfigurationMenu,
			Call:        guildChannel,
		},
		{
			Name:        "toggle",
			Description: "To toggle the guild channel.",
			Menu:        discord.ConfigurationMenu,
			Call:        guildEnable,
		},
	},
}

func HasPermission(User *discordgo.Member, s *discordgo.Session, Perm int) (isHe bool, err error) {
	if s == nil || User == nil {
		return false, errors.New("nil pointer")
	}
	g, err := s.Guild(User.GuildID)
	if err != nil {
		return false, err
	}
	for _, roleID := range User.Roles {
		for i := 0; i < len(g.Roles); i++ {
			if g.Roles[i].ID == roleID {
				if g.Roles[i].Permissions&int64(Perm) == int64(Perm) {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func guildChannel(ctx *discord.CmdContext) {
	member, _ := ctx.Session.GuildMember(ctx.GuildID, ctx.Author.ID)
	perm, err := HasPermission(member, ctx.Session, discordgo.PermissionAdministrator)
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content:   "An error occurred trying to get your permissions, try again a bit later.",
			Ephemeral: true,
		})
	}
	if !perm {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have permission to use this command.",
			Ephemeral: true,
		})
		return
	}
	ctx.Guild.AutomatedMessagesChannelID = ctx.ChannelID
	ctx.Guild.Save()
	ctx.Reply(discord.ReplyParams{
		Content:   "Guild channel set.",
		Ephemeral: true,
	})
}

func guildEnable(ctx *discord.CmdContext) {
	member, _ := ctx.Session.GuildMember(ctx.GuildID, ctx.Author.ID)
	perm, err := HasPermission(member, ctx.Session, discordgo.PermissionAdministrator)
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content:   "An error occurred trying to get your permissions, try again a bit later.",
			Ephemeral: true,
		})
	}
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content:   "An error occurred trying to get your permissions, try again a bit later.",
			Ephemeral: true,
		})
	}
	if !perm {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have permission to use this command.",
			Ephemeral: true,
		})
		return
	}
	ctx.Guild.AutomatedMessagesEnabled = !ctx.Guild.AutomatedMessagesEnabled
	ctx.Guild.Save()
	if ctx.Guild.AutomatedMessagesEnabled {
		ctx.Reply(discord.ReplyParams{
			Content:   "The guild channel has been enabled.",
			Ephemeral: true,
		})
		return
	}
	ctx.Reply(discord.ReplyParams{
		Content:   "The guild channel has been disabled.",
		Ephemeral: true,
	})
}
