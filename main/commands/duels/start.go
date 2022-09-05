package commands_duels

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func DuelCreate(ctx *discord.CmdContext) {
	arg := ctx.Arguments.GetArg("user", 0, "")
	if !arg.Found {
		ctx.Reply(discord.ReplyParams{
			Content: "You need to specify an opponent.",
		})
		return
	}
	opponent := arg.Raw.UserValue(ctx.Session)

	// if opponent.ID == ctx.Author.ID {
	// 	ctx.Reply(discord.ReplyParams{
	// 		Content: "You can't duel yourself.",
	// 	})
	// 	return
	// }

	if !ctx.Player.SelectedChar.Valid() && !ctx.Player.SelectedGrimm.Valid() {
		ctx.Reply(discord.ReplyParams{
			Content: "You need to select a persona.",
		})
		return
	}

	var persona *BattlePersona
	if ctx.Player.SelectedType == models.CharType {
		ctx.Player.SelectedChar.CalcStats()
		persona = CharToPersona(ctx.Player.SelectedChar)
	} else {
		ctx.Player.SelectedGrimm.CalcStats()
		persona = GrimmToPersona(ctx.Player.SelectedGrimm)
	}

	NewBattle := &BattleStruct{
		ID: ctx.ID,
		Players: []*BattlePlayer{
			{
				User:         ctx.Author,
				SelectedID:   ctx.Player.SelectedID,
				SelectedType: ctx.Player.SelectedType,
			},
			{
				User: &discordgo.User{
					ID: opponent.ID,
				},
			},
		},
		Chars: []*BattlePersona{
			persona,
		},
		PlayersID:  make(map[string]int),
		ChannelID:  ctx.ChannelID,
		StartTime:  time.Now(),
		HasStarted: false,
		TurnNumber: 1,
		Finished:   false,
		Original:   ctx,
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Duel Request",
		Description: ctx.Author.Username + " will be playing " + persona.name + " ! \n" + opponent.Username + ", to accept this duel enter the command r!accept.",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: persona.iconURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "And special thanks to @Sonya#2665 for the background!",
		},
		Color: persona.RarityToColor(),
	}

	NewBattle.PlayersID[ctx.Author.ID] = 0
	NewBattle.PlayersID[opponent.ID] = 1

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          AcceptDuel,
		Data:          NewBattle,
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content:    embed,
		Components: AcceptComponent(ctx.ID),
	})
}

func AcceptDuel(ctx *discord.CmdContext) {
	battle := ctx.Menu.Data.(*BattleStruct)
	if battle == nil {
		ctx.Reply(discord.ReplyParams{
			Content: "No duel found.",
		})
		return
	}
	if battle.Finished {
		ctx.Reply(discord.ReplyParams{
			Content: "The duel has finished.",
		})
		return
	}
	if battle.HasStarted {
		ctx.Reply(discord.ReplyParams{
			Content: "The duel has already started.",
		})
		return
	}
	if ctx.Author.ID != battle.Players[1].User.ID {
		ctx.Reply(discord.ReplyParams{
			Content: "You can't accept this duel.",
		})
		return
	}

	var persona *BattlePersona
	if ctx.Player.SelectedChar.Valid() {
		ctx.Player.SelectedChar.CalcStats()
		persona = CharToPersona(ctx.Player.SelectedChar)
	} else {
		ctx.Player.SelectedGrimm.CalcStats()
		persona = GrimmToPersona(ctx.Player.SelectedGrimm)
	}
	battle.Players[1].User = ctx.Author
	battle.Chars = append(battle.Chars, persona)

	battle.Original.IsInteraction = false
	battle.Original.Reply(discord.ReplyParams{
		Content: duelEmbed(battle, 0, 1),
		ID:      battle.Original.Author.ID,
		DM:      true,
	})

	ctx.IsInteraction = false
	ctx.Reply(discord.ReplyParams{
		Content: duelEmbed(battle, 1, 0),
		ID:      ctx.Author.ID,
		DM:      true,
	})

	battle.Original.IsInteraction = true
	ctx.IsInteraction = true

	battle.HasStarted = true
	battle.Original = ctx

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          duelAttack,
		Data:          battle,
	}, 0)

	msg, _ := ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageSend{
			Files: []*discordgo.File{
				{
					Reader: createDuelImage(battle.Chars[0], battle.Chars[1]),
					Name:   "vs.png",
				},
			},
			Embed: &discordgo.MessageEmbed{
				Title:       battle.Players[0].User.Username + " VS " + battle.Players[1].User.Username,
				Description: "Turn : 1",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Latest event :",
						Value: "Nothing yet.",
					},
					{
						Name:   battle.Players[0].User.Username + "'s HP :",
						Value:  strconv.Itoa(battle.Chars[0].stats.Health),
						Inline: true,
					},
					{
						Name:   battle.Players[1].User.Username + "'s HP :",
						Value:  strconv.Itoa(battle.Chars[1].stats.Health),
						Inline: true,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://vs.png",
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: "Click the buttons to attack !",
				},
				Color: config.Botcolor,
			},
			Components: newDuelComponent(ctx.ID),
		},
		FollowUp: true,
	})
	battle.BattleMessage = msg
}

func AcceptComponent(menuID string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label:    "Accept",
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-accept",
				},
			},
		},
	}
}
