package discord

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/bwmarrin/discordgo"
)

type ReplyParams struct {
	Content     interface{}
	Components  []discordgo.MessageComponent
	Interaction *discordgo.Interaction
	Files       []*discordgo.File
	ID          string
	ChannelID   string
	GuildID     string

	DM        bool
	Edit      bool
	Delete    bool
	FollowUp  bool
	Ephemeral bool

	embeds []*discordgo.MessageEmbed
}

func (c *CmdContext) Reply(p ReplyParams) (st *discordgo.Message, err error) {
	if p.DM {
		channel, err := c.Session.UserChannelCreate(p.ID)
		if err != nil {
			p.ChannelID = c.ChannelID
			p.Content = "Sorry <@" + p.ID + ">, but I cannot contact you through DMs, check your privacy settings!"
		} else {
			p.ChannelID = channel.ID
		}
	}

	if p.ChannelID == "" {
		p.ChannelID = c.ChannelID
	}
	if p.GuildID == "" {
		p.GuildID = c.GuildID
	}

	// Better compatibility

	switch p.Content.(type) {
	case *discordgo.MessageEdit:
		val := p.Content.(*discordgo.MessageEdit)
		p.embeds = []*discordgo.MessageEmbed{}
		if val.Embed != nil {
			p.embeds = []*discordgo.MessageEmbed{
				val.Embed,
			}
		}
	case *discordgo.MessageSend:
		val := p.Content.(*discordgo.MessageSend)
		files := val.Files
		if val.File != nil {
			if files == nil {
				files = []*discordgo.File{val.File}
			}
		}
		val.Files = files
		val.File = nil
		p.embeds = []*discordgo.MessageEmbed{}
		if val.Embed != nil {
			p.embeds = []*discordgo.MessageEmbed{
				val.Embed,
			}
		}
	}

	if c.IsInteraction {
		return c.ReplyInteraction(p)
	}
	return c.ReplyClassic(p)
}

func (c *CmdContext) ReplyClassic(p ReplyParams) (st *discordgo.Message, err error) {
	if p.Delete {
		err = c.Session.ChannelMessageDelete(p.ChannelID, p.ID)
		return
	}
	switch p.Content.(type) {
	case string:
		if p.Edit {
			return c.Session.ChannelMessageEdit(p.ChannelID, p.ID, fmt.Sprint(p.Content))
		}
		if len(p.Components) > 0 {
			v := &discordgo.MessageSend{
				Content:    fmt.Sprint(p.Content),
				Components: p.Components,
			}
			return c.Session.ChannelMessageSendComplex(p.ChannelID, v)
		}
		return c.Session.ChannelMessageSend(p.ChannelID, fmt.Sprint(p.Content))
	case *discordgo.MessageEmbed:
		if p.Edit {
			v := &discordgo.MessageEdit{
				Embed:      p.Content.(*discordgo.MessageEmbed),
				Components: p.Components,

				ID:      p.ID,
				Channel: p.ChannelID,
			}
			return c.Session.ChannelMessageEditComplex(v)
		}
		v := &discordgo.MessageSend{
			Embed:      p.Content.(*discordgo.MessageEmbed),
			Components: p.Components,
		}
		return c.Session.ChannelMessageSendComplex(p.ChannelID, v)
	case *discordgo.MessageSend:
		if p.Edit {
			complex := p.Content.(*discordgo.MessageSend)
			v := &discordgo.MessageEdit{
				Content:    &complex.Content,
				Embed:      complex.Embed,
				Components: complex.Components,

				ID:      p.ID,
				Channel: p.ChannelID,
			}
			return c.Session.ChannelMessageEditComplex(v)
		}
		return c.Session.ChannelMessageSendComplex(p.ChannelID, p.Content.(*discordgo.MessageSend))
	case *discordgo.MessageEdit:
		return c.Session.ChannelMessageEditComplex(p.Content.(*discordgo.MessageEdit))
	default:
		fmt.Println("unknown")
	}
	return
}

func (c *CmdContext) ReplyInteraction(p ReplyParams) (st *discordgo.Message, err error) {
	var flags uint64
	if p.Ephemeral {
		flags = 1 << 6
	}
	if p.Delete {
		err = c.Session.InteractionResponseDelete(p.ChannelID, c.Interaction)
		return
	}
	switch p.Content.(type) {
	case string:
		if !p.FollowUp {
			if p.Edit {
				return c.Session.InteractionResponseEdit(config.AppID, c.Interaction, &discordgo.WebhookEdit{
					Content:    p.Content.(string),
					Components: p.Components,
				})
			}
			err = c.Session.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:      flags,
					Content:    p.Content.(string),
					Components: p.Components,
				},
			})
			return
		} else {
			if p.Edit {
				return c.Session.FollowupMessageEdit(config.AppID, c.Interaction, p.ID, &discordgo.WebhookEdit{
					Content:    p.Content.(string),
					Components: p.Components,
				})
			}
			return c.Session.FollowupMessageCreate(config.AppID, c.Interaction, true, &discordgo.WebhookParams{
				Content:    p.Content.(string),
				Components: p.Components,
				Flags:      flags,
			})
		}
	case *discordgo.MessageEmbed:
		if !p.FollowUp {
			if p.Edit {
				return c.Session.InteractionResponseEdit(config.AppID, c.Interaction, &discordgo.WebhookEdit{
					Embeds: []*discordgo.MessageEmbed{
						p.Content.(*discordgo.MessageEmbed),
					},
					Components: p.Components,
				})
			}
			err = c.Session.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: flags,
					Embeds: []*discordgo.MessageEmbed{
						p.Content.(*discordgo.MessageEmbed),
					},
					Components: p.Components,
				},
			})
			return
		} else {
			if p.Edit {
				return c.Session.FollowupMessageEdit(config.AppID, c.Interaction, p.ID, &discordgo.WebhookEdit{
					Embeds: []*discordgo.MessageEmbed{
						p.Content.(*discordgo.MessageEmbed),
					},
					Components: p.Components,
				})
			}
			return c.Session.FollowupMessageCreate(config.AppID, c.Interaction, true, &discordgo.WebhookParams{
				Embeds: []*discordgo.MessageEmbed{
					p.Content.(*discordgo.MessageEmbed),
				},
				Components: p.Components,
				Flags:      flags,
			})
		}
	case *discordgo.MessageSend:
		complex := p.Content.(*discordgo.MessageSend)
		if !p.FollowUp {
			if p.Edit {
				return c.Session.InteractionResponseEdit(config.AppID, c.Interaction, &discordgo.WebhookEdit{
					Content:    complex.Content,
					Embeds:     p.embeds,
					Components: complex.Components,
					Files:      complex.Files,
				})
			}
			err = c.Session.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   flags,
					Content: complex.Content,
					Embeds: []*discordgo.MessageEmbed{
						complex.Embed,
					},
					Components: complex.Components,
					Files:      complex.Files,
				},
			})
			return
		} else {
			if p.Edit {
				return c.Session.FollowupMessageEdit(config.AppID, c.Interaction, p.ID, &discordgo.WebhookEdit{
					Content:    complex.Content,
					Embeds:     p.embeds,
					Components: complex.Components,
					Files:      complex.Files,
				})
			}
			return c.Session.FollowupMessageCreate(config.AppID, c.Interaction, true, &discordgo.WebhookParams{
				Content:    complex.Content,
				Embeds:     p.embeds,
				Components: complex.Components,
				Files:      complex.Files,
				Flags:      flags,
			})
		}
	case *discordgo.MessageEdit:
		complex := p.Content.(*discordgo.MessageEdit)
		if !p.FollowUp {
			if complex.Content == nil {
				t := ""
				complex.Content = &t
			}
			return c.Session.InteractionResponseEdit(config.AppID, c.Interaction, &discordgo.WebhookEdit{
				Content:    *complex.Content,
				Embeds:     p.embeds,
				Components: complex.Components,
				Files:      p.Files,
			})
		} else {
			if complex.Content == nil {
				t := ""
				complex.Content = &t
			}
			return c.Session.FollowupMessageEdit(config.AppID, c.Interaction, complex.ID, &discordgo.WebhookEdit{
				Content:    *complex.Content,
				Embeds:     p.embeds,
				Components: complex.Components,
				Files:      p.Files,
			})
		}
	default:
		fmt.Println("unknown")
	}
	return
}
