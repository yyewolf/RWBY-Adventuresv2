package commands_trade

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
)

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
	if trade.ReceiverID != ctx.Author.ID {
		ctx.Reply(discord.ReplyParams{
			Content: "Trade ID not found.",
		})
		return
	}

	sender := models.GetPlayer(trade.SenderID)
	receiver := ctx.Player

	SenderDM, _ := discord.Session.UserChannelCreate(trade.SenderID)

	errStr := fmt.Sprintf("Trade `%s` has been cancelled :", trade.ID)

	if trade.UserSends.Money < 0 || trade.UserSends.Boxes < 0 || trade.UserSends.RareBoxes < 0 || trade.UserSends.GrimmBoxes < 0 || trade.UserSends.RareGrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s there is an issue with their boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s there is an issue with your boxes.", errStr))
		trade.Delete()
		return
	}
	if trade.TargetSends.Money < 0 || trade.TargetSends.Boxes < 0 || trade.TargetSends.RareBoxes < 0 || trade.TargetSends.GrimmBoxes < 0 || trade.TargetSends.RareGrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s there is an issue with your boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s there is an issue with their boxes.", errStr))
		trade.Delete()
		return
	}
	// We check that the player has what he claims to
	if !sender.VerifyChars(trade.UserSends.Characters) || !sender.VerifyGrimms(trade.UserSends.Grimms) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s there is an issue with their personas.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s there is an issue with your personas.", errStr))
		trade.Delete()
		return
	}

	// We check that the target has what the player claims he has
	if !receiver.VerifyChars(trade.TargetSends.Characters) || !receiver.VerifyGrimms(trade.TargetSends.Grimms) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s there is an issue with your personas.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s there is an issue with their personas.", errStr))
		trade.Delete()
		return
	}
	// We check that the player has enough money
	if sender.TotalBalance() < int64(trade.UserSends.Money) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they don't have enough balance.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you don't have enough balance.", errStr))
		trade.Delete()
		return
	}
	// We check that the target has enough money
	if receiver.TotalBalance() < int64(trade.TargetSends.Money) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you don't have enough balance.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they don't have enough balance.", errStr))
		trade.Delete()
		return
	}

	PlayerResult := &models.TradeContent{
		Money:          int64(sender.TotalBalance()) + trade.TargetSends.Money - trade.UserSends.Money,
		Boxes:          int64(sender.Boxes.Boxes) + trade.TargetSends.Boxes - trade.UserSends.Boxes,
		RareBoxes:      int64(sender.Boxes.RareBoxes) + trade.TargetSends.RareBoxes - trade.UserSends.RareBoxes,
		GrimmBoxes:     int64(sender.Boxes.GrimmBoxes) + trade.TargetSends.GrimmBoxes - trade.UserSends.GrimmBoxes,
		RareGrimmBoxes: int64(sender.Boxes.RareGrimmBoxes) + trade.TargetSends.RareGrimmBoxes - trade.UserSends.RareGrimmBoxes,
	}

	TargetResult := &models.TradeContent{
		Money:          int64(receiver.TotalBalance()) - trade.TargetSends.Money + trade.UserSends.Money,
		Boxes:          int64(receiver.Boxes.Boxes) - trade.TargetSends.Boxes + trade.UserSends.Boxes,
		RareBoxes:      int64(receiver.Boxes.RareBoxes) - trade.TargetSends.RareBoxes + trade.UserSends.RareBoxes,
		GrimmBoxes:     int64(receiver.Boxes.GrimmBoxes) - trade.TargetSends.GrimmBoxes + trade.UserSends.GrimmBoxes,
		RareGrimmBoxes: int64(receiver.Boxes.RareGrimmBoxes) - trade.TargetSends.RareGrimmBoxes + trade.UserSends.RareGrimmBoxes,
	}

	if PlayerResult.Boxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they don't have enough boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you don't have enough boxes.", errStr))
		trade.Delete()
		return
	}
	if PlayerResult.RareBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they don't have enough rare boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you don't have enough rare boxes.", errStr))
		trade.Delete()
		return
	}
	if PlayerResult.GrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they don't have enough grimm boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you don't have enough grimm boxes.", errStr))
		trade.Delete()
		return
	}
	if PlayerResult.RareGrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they don't have enough rare grimm boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you don't have enough rare grimm boxes.", errStr))
		trade.Delete()
		return
	}
	if PlayerResult.Boxes+PlayerResult.RareBoxes > int64(sender.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would receive too much boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would receive too much boxes.", errStr))
		trade.Delete()
		return
	}
	if PlayerResult.RareGrimmBoxes+PlayerResult.GrimmBoxes > int64(sender.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would receive too much grimm boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would receive too much grimm boxes.", errStr))
		trade.Delete()
		return
	}

	if TargetResult.Boxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you don't have enough boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they don't have enough boxes.", errStr))
		trade.Delete()
		return
	}
	if TargetResult.RareBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you don't have enough rare boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they don't have enough rare boxes.", errStr))
		trade.Delete()
		return
	}
	if TargetResult.GrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you don't have enough grimm boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they don't have enough grimm boxes.", errStr))
		trade.Delete()
		return
	}
	if TargetResult.RareGrimmBoxes < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you don't have enough rare grimm boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they don't have enough rare grimm boxes.", errStr))
		trade.Delete()
		return
	}
	if TargetResult.Boxes+TargetResult.RareBoxes > int64(sender.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would receive too much boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would receive too much boxes.", errStr))
		trade.Delete()
		return
	}
	if TargetResult.RareGrimmBoxes+TargetResult.GrimmBoxes > int64(sender.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would receive too much grimm boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would receive too much grimm boxes.", errStr))
		trade.Delete()
		return
	}

	if len(sender.Characters)+len(trade.TargetSends.Characters)-len(trade.UserSends.Characters) > sender.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would receive too much characters.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would receive too much characters.", errStr))
		trade.Delete()
		return
	}
	if len(sender.Grimms)+len(trade.TargetSends.Grimms)-len(trade.UserSends.Grimms) > sender.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would receive too much grimms.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would receive too much grimms.", errStr))
		trade.Delete()
		return
	}
	if len(sender.Characters)+len(trade.TargetSends.Characters)-len(trade.UserSends.Characters) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would be left without any characters.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would be left without any characters.", errStr))
		trade.Delete()
		return
	}
	if len(sender.Grimms)+len(trade.TargetSends.Grimms)-len(trade.UserSends.Grimms) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would be left without any grimms.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would be left without any grimms.", errStr))
		trade.Delete()
		return
	}

	if len(receiver.Characters)-len(trade.TargetSends.Characters)+len(trade.UserSends.Characters) > receiver.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would receive too much characters.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would receive too much characters.", errStr))
		trade.Delete()
		return
	}
	if len(receiver.Grimms)-len(trade.TargetSends.Grimms)+len(trade.UserSends.Grimms) > receiver.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would receive too much grimms.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would receive too much grimms.", errStr))
		trade.Delete()
		return
	}
	if len(receiver.Characters)-len(trade.TargetSends.Characters)+len(trade.UserSends.Characters) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would be left without any characters.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would be left without any characters.", errStr))
		trade.Delete()
		return
	}
	if len(receiver.Grimms)-len(trade.TargetSends.Grimms)+len(trade.UserSends.Grimms) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would be left without any grimms.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would be left without any grimms.", errStr))
		trade.Delete()
		return
	}

	for _, CharID := range trade.UserSends.Characters {
		// Switch selected if we need to (RUNS ONLY ONCE)
		if CharID == sender.SelectedID {
			for _, char := range sender.Characters {
				var nope = false
				for _, nopeID := range trade.UserSends.Characters {
					if char.CharID == nopeID {
						nope = true
					}
				}
				if !nope {
					sender.SelectedID = char.CharID
					break
				}
			}
		}
		for _, char := range sender.Characters {
			if CharID == char.CharID {
				char.UserID = receiver.DiscordID
				config.Database.Save(&char)
				break
			}
		}
	}

	for _, GrimmID := range trade.UserSends.Grimms {
		// Switch selected if we need to (RUNS ONLY ONCE)
		if GrimmID == sender.SelectedID {
			for _, grimm := range sender.Grimms {
				var nope = false
				for _, nopeID := range trade.UserSends.Grimms {
					if grimm.GrimmID == nopeID {
						nope = true
					}
				}
				if !nope {
					sender.SelectedID = grimm.GrimmID
					break
				}
			}
		}
		for _, grimm := range sender.Grimms {
			if GrimmID == grimm.GrimmID {
				grimm.UserID = receiver.DiscordID
				config.Database.Save(&grimm)
				break
			}
		}
	}

	for _, CharID := range trade.TargetSends.Characters {
		// Switch selected if we need to (RUNS ONLY ONCE)
		if CharID == receiver.SelectedID {
			for _, char := range receiver.Characters {
				var nope = false
				for _, nopeID := range trade.TargetSends.Characters {
					if char.CharID == nopeID {
						nope = true
					}
				}
				if !nope {
					receiver.SelectedID = char.CharID
					break
				}
			}
		}
		for _, char := range receiver.Characters {
			if CharID == char.CharID {
				char.UserID = sender.DiscordID
				config.Database.Save(&char)
				break
			}
		}
	}

	for _, GrimmID := range trade.TargetSends.Grimms {
		// Switch selected if we need to (RUNS ONLY ONCE)
		if GrimmID == receiver.SelectedID {
			for _, grimm := range receiver.Grimms {
				var nope = false
				for _, nopeID := range trade.TargetSends.Grimms {
					if grimm.GrimmID == nopeID {
						nope = true
					}
				}
				if !nope {
					receiver.SelectedID = grimm.GrimmID
					break
				}
			}
		}
		for _, grimm := range receiver.Grimms {
			if GrimmID == grimm.GrimmID {
				grimm.UserID = sender.DiscordID
				config.Database.Save(&grimm)
				break
			}
		}
	}

	sender.Balance += trade.TargetSends.Money - trade.UserSends.Money
	sender.Boxes.Boxes += int(trade.TargetSends.Boxes) - int(trade.UserSends.Boxes)
	sender.Boxes.RareBoxes += int(trade.TargetSends.RareBoxes) - int(trade.UserSends.RareBoxes)
	sender.Boxes.GrimmBoxes += int(trade.TargetSends.GrimmBoxes) - int(trade.UserSends.GrimmBoxes)
	sender.Boxes.RareGrimmBoxes += int(trade.TargetSends.RareGrimmBoxes) - int(trade.UserSends.RareGrimmBoxes)

	receiver.Balance -= trade.TargetSends.Money - trade.UserSends.Money
	receiver.Boxes.Boxes -= int(trade.TargetSends.Boxes) - int(trade.UserSends.Boxes)
	receiver.Boxes.RareBoxes -= int(trade.TargetSends.RareBoxes) - int(trade.UserSends.RareBoxes)
	receiver.Boxes.GrimmBoxes -= int(trade.TargetSends.GrimmBoxes) - int(trade.UserSends.GrimmBoxes)
	receiver.Boxes.RareGrimmBoxes -= int(trade.TargetSends.RareGrimmBoxes) - int(trade.UserSends.RareGrimmBoxes)

	config.Database.Save(sender.Boxes)
	config.Database.Save(receiver.Boxes)
	config.Database.Save(sender)
	config.Database.Save(receiver)
	trade.Delete()

	ctx.Reply(discord.ReplyParams{
		Content:   fmt.Sprintf("You have accepted trade `%s` !", trade.ID),
		Ephemeral: true,
	})
	discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("Trade `%s` has been accepted by the user !", trade.ID))

}
