package commands_trade

import (
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

var TradesCommand = &discord.Command{
	Name:        "trade",
	Description: "All commands regarding trades.",
	SubCommands: []*discord.Command{
		{
			Name:        "accept",
			Description: "Accept a traade.",
			Menu:        discord.InventoryMenu,
			Call:        AcceptTrade,
			Args: []discord.Arg{
				{
					Name:        "id",
					Description: "Identification number of your persona.",
					Size:        1,
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
	},
}

func AcceptTrade(ctx *discord.CmdContext) {
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
	if trade.ReceiverID != id {
		ctx.Reply(discord.ReplyParams{
			Content: "Trade ID not found.",
		})
		return
	}

	player := models.GetPlayer(trade.SenderID)
	target := ctx.Player

	if trade.UserSends.Money < 0 || trade.UserSends.Boxes < 0 || trade.UserSends.RareBoxes < 0 || trade.UserSends.GrimmBoxes < 0 || trade.UserSends.RareGrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "There is an issue with your boxes.",
			Ephemeral: true,
		})
		return
	}
	if trade.TargetSends.Money < 0 || trade.TargetSends.Boxes < 0 || trade.TargetSends.RareBoxes < 0 || trade.TargetSends.GrimmBoxes < 0 || trade.TargetSends.RareGrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "There is an issue with their boxes.",
			Ephemeral: true,
		})
		return
	}
	// We check that the player has what he claims to
	if !player.VerifyChars(trade.UserSends.Characters) || !player.VerifyGrimms(trade.UserSends.Grimms) {
		ctx.Reply(discord.ReplyParams{
			Content:   "There is an issue with your personas.",
			Ephemeral: true,
		})
		return
	}

	// We check that the target has what the player claims he has
	if !target.VerifyChars(trade.TargetSends.Characters) || !target.VerifyGrimms(trade.TargetSends.Grimms) {
		ctx.Reply(discord.ReplyParams{
			Content:   "There is an issue with your target's personas.",
			Ephemeral: true,
		})
		return
	}

	if player.TradeSent >= 5 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You have too much trade on hold.",
			Ephemeral: true,
		})
		return
	}
	if target.TradeReceived >= 8 {
		ctx.Reply(discord.ReplyParams{
			Content:   "This person already has too much trades on hold.",
			Ephemeral: true,
		})
		return
	}

	if player.TotalBalance() < int64(trade.UserSends.Money) {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have enough balance.",
			Ephemeral: true,
		})
		return
	}
	if target.TotalBalance() < int64(trade.TargetSends.Money) {
		ctx.Reply(discord.ReplyParams{
			Content:   "This person doesn't have enough balance.",
			Ephemeral: true,
		})
		return
	}

	PlayerResult := &models.TradeContent{
		Money:          int64(player.TotalBalance()) + trade.TargetSends.Money - trade.UserSends.Money,
		Boxes:          int64(player.Boxes.Boxes) + trade.TargetSends.Boxes - trade.UserSends.Boxes,
		RareBoxes:      int64(player.Boxes.RareBoxes) + trade.TargetSends.RareBoxes - trade.UserSends.RareBoxes,
		GrimmBoxes:     int64(player.Boxes.GrimmBoxes) + trade.TargetSends.GrimmBoxes - trade.UserSends.GrimmBoxes,
		RareGrimmBoxes: int64(player.Boxes.RareGrimmBoxes) + trade.TargetSends.RareGrimmBoxes - trade.UserSends.RareGrimmBoxes,
	}

	TargetResult := &models.TradeContent{
		Money:          int64(target.TotalBalance()) - trade.TargetSends.Money + trade.UserSends.Money,
		Boxes:          int64(target.Boxes.Boxes) - trade.TargetSends.Boxes + trade.UserSends.Boxes,
		RareBoxes:      int64(target.Boxes.RareBoxes) - trade.TargetSends.RareBoxes + trade.UserSends.RareBoxes,
		GrimmBoxes:     int64(target.Boxes.GrimmBoxes) - trade.TargetSends.GrimmBoxes + trade.UserSends.GrimmBoxes,
		RareGrimmBoxes: int64(target.Boxes.RareGrimmBoxes) - trade.TargetSends.RareGrimmBoxes + trade.UserSends.RareGrimmBoxes,
	}

	if PlayerResult.Money < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have enough balance.",
			Ephemeral: true,
		})
		return
	}
	if PlayerResult.Boxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have enough boxes.",
			Ephemeral: true,
		})
		return
	}
	if PlayerResult.RareBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have enough rare boxes.",
			Ephemeral: true,
		})
		return
	}
	if PlayerResult.GrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have enough grimm boxes.",
			Ephemeral: true,
		})
		return
	}
	if PlayerResult.RareGrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You don't have enough rare grimm boxes.",
			Ephemeral: true,
		})
		return
	}
	if PlayerResult.Boxes+PlayerResult.RareBoxes > int64(player.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   "You would receive too much boxes.",
			Ephemeral: true,
		})
		return
	}
	if PlayerResult.RareGrimmBoxes+PlayerResult.GrimmBoxes > int64(player.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   "You would receive too much grimm boxes.",
			Ephemeral: true,
		})
		return
	}

	if TargetResult.Money < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "They don't have enough balance.",
			Ephemeral: true,
		})
		return
	}
	if TargetResult.Boxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "They don't have enough boxes.",
			Ephemeral: true,
		})
		return
	}
	if TargetResult.RareBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "They don't have enough rare boxes.",
			Ephemeral: true,
		})
		return
	}
	if TargetResult.GrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "They don't have enough grimm boxes.",
			Ephemeral: true,
		})
		return
	}
	if TargetResult.RareGrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "They don't have enough rare grimm boxes.",
			Ephemeral: true,
		})
		return
	}
	if TargetResult.Boxes+TargetResult.RareBoxes > int64(player.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   "They would receive too much boxes.",
			Ephemeral: true,
		})
		return
	}
	if TargetResult.RareGrimmBoxes+TargetResult.GrimmBoxes > int64(player.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   "They would receive too much grimm boxes.",
			Ephemeral: true,
		})
		return
	}

	if len(player.Characters)+len(trade.TargetSends.Characters)-len(trade.UserSends.Characters) > player.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   "You would receive too much characters.",
			Ephemeral: true,
		})
		return
	}
	if len(player.Grimms)+len(trade.TargetSends.Grimms)-len(trade.UserSends.Grimms) > player.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   "You would receive too much grimms.",
			Ephemeral: true,
		})
		return
	}
	if len(player.Characters)+len(trade.TargetSends.Characters)-len(trade.UserSends.Characters) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You would be left without any characters.",
			Ephemeral: true,
		})
		return
	}
	if len(player.Grimms)+len(trade.TargetSends.Grimms)-len(trade.UserSends.Grimms) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "You would be left without any grimms.",
			Ephemeral: true,
		})
		return
	}

	if len(target.Characters)-len(trade.TargetSends.Characters)+len(trade.UserSends.Characters) > target.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   "They would receive too much characters.",
			Ephemeral: true,
		})
		return
	}
	if len(target.Grimms)-len(trade.TargetSends.Grimms)+len(trade.UserSends.Grimms) > target.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   "They would receive too much grimms.",
			Ephemeral: true,
		})
		return
	}
	if len(target.Characters)-len(trade.TargetSends.Characters)+len(trade.UserSends.Characters) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "They would be left without any characters.",
			Ephemeral: true,
		})
		return
	}
	if len(target.Grimms)-len(trade.TargetSends.Grimms)+len(trade.UserSends.Grimms) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "They would be left without any grimms.",
			Ephemeral: true,
		})
		return
	}

	for _, CharID := range trade.UserSends.Characters {
		// Switch selected if we need to (RUNS ONLY ONCE)
		if CharID == player.SelectedID {
			for _, char := range player.Characters {
				var nope = false
				for _, nopeID := range trade.UserSends.Characters {
					if char.CharID == nopeID {
						nope = true
					}
				}
				if !nope {
					player.SelectedID = char.CharID
					break
				}
			}
		}
		for _, char := range player.Characters {
			if CharID == char.CharID {
				char.UserID = target.DiscordID
				config.Database.Save(char)
				break
			}
		}
	}

	for _, GrimmID := range trade.UserSends.Grimms {
		// Switch selected if we need to (RUNS ONLY ONCE)
		if GrimmID == player.SelectedID {
			for _, grimm := range player.Grimms {
				var nope = false
				for _, nopeID := range trade.UserSends.Grimms {
					if grimm.GrimmID == nopeID {
						nope = true
					}
				}
				if !nope {
					player.SelectedID = grimm.GrimmID
					break
				}
			}
		}
		for _, grimm := range player.Grimms {
			if GrimmID == grimm.GrimmID {
				grimm.UserID = target.DiscordID
				config.Database.Save(grimm)
				break
			}
		}
	}

	for _, CharID := range trade.TargetSends.Characters {
		// Switch selected if we need to (RUNS ONLY ONCE)
		if CharID == target.SelectedID {
			for _, char := range target.Characters {
				var nope = false
				for _, nopeID := range trade.TargetSends.Characters {
					if char.CharID == nopeID {
						nope = true
					}
				}
				if !nope {
					target.SelectedID = char.CharID
					break
				}
			}
		}
		for _, char := range target.Characters {
			if CharID == char.CharID {
				char.UserID = player.DiscordID
				config.Database.Save(char)
				break
			}
		}
	}

	for _, GrimmID := range trade.TargetSends.Grimms {
		// Switch selected if we need to (RUNS ONLY ONCE)
		if GrimmID == target.SelectedID {
			for _, grimm := range target.Grimms {
				var nope = false
				for _, nopeID := range trade.TargetSends.Grimms {
					if grimm.GrimmID == nopeID {
						nope = true
					}
				}
				if !nope {
					target.SelectedID = grimm.GrimmID
					break
				}
			}
		}
		for _, grimm := range target.Grimms {
			if GrimmID == grimm.GrimmID {
				grimm.UserID = player.DiscordID
				config.Database.Save(grimm)
				break
			}
		}
	}

	player.Balance += trade.TargetSends.Money - trade.UserSends.Money
	player.Boxes.Boxes += int(trade.TargetSends.Boxes) - int(trade.UserSends.Boxes)
	player.Boxes.RareBoxes += int(trade.TargetSends.RareBoxes) - int(trade.UserSends.RareBoxes)
	player.Boxes.GrimmBoxes += int(trade.TargetSends.GrimmBoxes) - int(trade.UserSends.GrimmBoxes)
	player.Boxes.RareGrimmBoxes += int(trade.TargetSends.RareGrimmBoxes) - int(trade.UserSends.RareGrimmBoxes)

	target.Balance -= trade.TargetSends.Money - trade.UserSends.Money
	target.Boxes.Boxes -= int(trade.TargetSends.Boxes) - int(trade.UserSends.Boxes)
	target.Boxes.RareBoxes -= int(trade.TargetSends.RareBoxes) - int(trade.UserSends.RareBoxes)
	target.Boxes.GrimmBoxes -= int(trade.TargetSends.GrimmBoxes) - int(trade.UserSends.GrimmBoxes)
	target.Boxes.RareGrimmBoxes -= int(trade.TargetSends.RareGrimmBoxes) - int(trade.UserSends.RareGrimmBoxes)

	config.Database.Save(player.Boxes)
	config.Database.Save(target.Boxes)
	config.Database.Save(player)
	config.Database.Save(target)
	config.Database.Delete(trade.TargetSends)
	config.Database.Delete(trade.UserSends)
	config.Database.Delete(trade)
	config.Database.Commit()
}
