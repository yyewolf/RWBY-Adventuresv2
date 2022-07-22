package commands_settings

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// TO DO : store this in a database

var ReportCommand = &discord.Command{
	Name:        "report",
	Description: "Report an issue.",
	Menu:        discord.ConfigurationMenu,
	Call:        report,
}

type ReportData struct {
	SenderID   string
	OriginalID string
	Text       string
	Images     string
}

func report(ctx *discord.CmdContext) {
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: ReportModal(ctx.ID),
		},
	})

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          ReportCallback,
	}, 0)
}

func ReportCallback(ctx *discord.CmdContext) {
	if ctx.Author.ID != ctx.Menu.SourceContext.Author.ID {
		return
	}
	data := &ReportData{
		SenderID: ctx.Author.ID,
	}
	for _, comp := range ctx.ModalData.Components {
		c := comp.(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
		switch c.CustomID {
		case "report-text":
			data.Text = c.Value
		case "report-images":
			data.Images = c.Value
		}
	}
	ctx.Reply(discord.ReplyParams{
		Content:   "Thank you for your report. We will review it as soon as possible. You may also be contacted be in DMs by the support team.",
		FollowUp:  true,
		Ephemeral: true,
	})

	fields := []*discordgo.MessageEmbedField{
		{Name: "Reply", Value: data.Text},
	}
	if data.Images != "" {
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Images", Value: data.Images})
	}
	ctx.IsInteraction = false
	msg, err := ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:  "New ticket",
			Fields: fields,
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("Reported by %s%s", ctx.Author.Username, ctx.Author.Discriminator),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL(""),
			},
			Color: config.Botcolor,
		},
		Components: TicketButtons(ctx.ID),
		ChannelID:  config.ReportChannel,
	})
	if err != nil {
		return
	}
	data.OriginalID = msg.ID
	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          ReportReplyCallback,
		Data:          data,
		Modal:         true,
	}, 0)
}

func ReportReplyCallback(ctx *discord.CmdContext) {
	data := ctx.Menu.Data.(*ReportData)
	split := strings.Split(ctx.ComponentData.CustomID, "-")
	switch split[1] {
	case "reply":
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseModal,
				Data: ReportReplyModal(ctx.ID),
			},
		})

		discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
			MenuID:        ctx.ID,
			SourceContext: ctx,
			Call:          ReplyCallback,
			Data:          data,
		}, 0)
	case "delete":
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredMessageUpdate,
			},
		})

		ctx.Reply(discord.ReplyParams{
			ID:        data.OriginalID,
			ChannelID: config.ReportChannel,
			Delete:    true,
		})
	}
}

func ReplyCallback(ctx *discord.CmdContext) {
	data := ctx.Menu.Data.(*ReportData)
	var resp string
	for _, comp := range ctx.ModalData.Components {
		c := comp.(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
		switch c.CustomID {
		case "reply-text":
			resp = c.Value
		}
	}

	fields := []*discordgo.MessageEmbedField{
		{Name: "Problem", Value: data.Text},
	}
	if data.Images != "" {
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Images", Value: data.Images})
	}
	fields = append(fields, &discordgo.MessageEmbedField{Name: "Reply", Value: resp})

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:  "New ticket - Already replied to",
			Fields: fields,
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("Reported by %s%s and replied by %s%s", ctx.Author.Username, ctx.Author.Discriminator, ctx.Author.Username, ctx.Author.Discriminator),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ctx.Author.AvatarURL(""),
			},
			Color: config.Botcolor,
		},
		Edit: true,
	})

	ctx.IsInteraction = false
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title: "Ticket Reply",
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Original problem", Value: data.Text},
				{Name: "Response", Value: resp},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("Replied by %s%s", ctx.Author.Username, ctx.Author.Discriminator),
			},
			Color: config.Botcolor,
		},
		ID: data.SenderID,
		DM: true,
	})
}

func ReportModal(menuID string) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Title: "Report an issue",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.TextInput{
						CustomID:    "report-text",
						Label:       "Report",
						Style:       discordgo.TextInputParagraph,
						Placeholder: "Please describe your issue.",
						Required:    true,
						MaxLength:   600,
						MinLength:   10,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.TextInput{
						CustomID:    "report-images",
						Label:       "Images",
						Style:       discordgo.TextInputParagraph,
						Placeholder: "Please provide any images (link) you can that might help explain your issue.",
						MaxLength:   300,
						MinLength:   0,
					},
				},
			},
		},
		CustomID: menuID,
	}
}

func ReportReplyModal(menuID string) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Title: "Replying",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.TextInput{
						CustomID:    "reply-text",
						Label:       "Reply",
						Style:       discordgo.TextInputParagraph,
						Placeholder: "Please help them.",
						Required:    true,
						MaxLength:   2000,
						MinLength:   10,
					},
				},
			},
		},
		CustomID: menuID,
	}
}

func TicketButtons(menuID string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label:    "Reply",
					Style:    discordgo.PrimaryButton,
					CustomID: fmt.Sprintf("%s-reply", menuID),
				},
				&discordgo.Button{
					Label:    "Delete",
					Style:    discordgo.DangerButton,
					CustomID: fmt.Sprintf("%s-delete", menuID),
				},
			},
		},
	}
}
