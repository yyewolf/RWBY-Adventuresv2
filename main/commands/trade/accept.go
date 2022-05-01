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

	player := models.GetPlayer(trade.SenderID)
	target := ctx.Player

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
	if !player.VerifyChars(trade.UserSends.Characters) || !player.VerifyGrimms(trade.UserSends.Grimms) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s there is an issue with their personas.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s there is an issue with your personas.", errStr))
		trade.Delete()
		return
	}

	// We check that the target has what the player claims he has
	if !target.VerifyChars(trade.TargetSends.Characters) || !target.VerifyGrimms(trade.TargetSends.Grimms) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s there is an issue with your personas.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s there is an issue with their personas.", errStr))
		trade.Delete()
		return
	}

	if player.TotalBalance() < int64(trade.UserSends.Money) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they don't have enough balance.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you don't have enough balance.", errStr))
		trade.Delete()
		return
	}
	if target.TotalBalance() < int64(trade.TargetSends.Money) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you don't have enough balance.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they don't have enough balance.", errStr))
		trade.Delete()
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
	if PlayerResult.Boxes+PlayerResult.RareBoxes > int64(player.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would receive too much boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would receive too much boxes.", errStr))
		trade.Delete()
		return
	}
	if PlayerResult.RareGrimmBoxes+PlayerResult.GrimmBoxes > int64(player.MaxChar()) {
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
	if TargetResult.Boxes+TargetResult.RareBoxes > int64(player.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would receive too much boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would receive too much boxes.", errStr))
		trade.Delete()
		return
	}
	if TargetResult.RareGrimmBoxes+TargetResult.GrimmBoxes > int64(player.MaxChar()) {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would receive too much grimm boxes.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would receive too much grimm boxes.", errStr))
		trade.Delete()
		return
	}

	if len(player.Characters)+len(trade.TargetSends.Characters)-len(trade.UserSends.Characters) > player.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would receive too much characters.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would receive too much characters.", errStr))
		trade.Delete()
		return
	}
	if len(player.Grimms)+len(trade.TargetSends.Grimms)-len(trade.UserSends.Grimms) > player.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would receive too much grimms.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would receive too much grimms.", errStr))
		trade.Delete()
		return
	}
	if len(player.Characters)+len(trade.TargetSends.Characters)-len(trade.UserSends.Characters) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would be left without any characters.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would be left without any characters.", errStr))
		trade.Delete()
		return
	}
	if len(player.Grimms)+len(trade.TargetSends.Grimms)-len(trade.UserSends.Grimms) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s they would be left without any grimms.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s you would be left without any grimms.", errStr))
		trade.Delete()
		return
	}

	if len(target.Characters)-len(trade.TargetSends.Characters)+len(trade.UserSends.Characters) > target.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would receive too much characters.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would receive too much characters.", errStr))
		trade.Delete()
		return
	}
	if len(target.Grimms)-len(trade.TargetSends.Grimms)+len(trade.UserSends.Grimms) > target.CharLimit {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would receive too much grimms.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would receive too much grimms.", errStr))
		trade.Delete()
		return
	}
	if len(target.Characters)-len(trade.TargetSends.Characters)+len(trade.UserSends.Characters) < 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s you would be left without any characters.", errStr),
			Ephemeral: true,
		})
		discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("%s they would be left without any characters.", errStr))
		trade.Delete()
		return
	}
	if len(target.Grimms)-len(trade.TargetSends.Grimms)+len(trade.UserSends.Grimms) < 0 {
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
				config.Database.Save(&char)
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
				config.Database.Save(&grimm)
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
				config.Database.Save(&char)
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
				config.Database.Save(&grimm)
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
	trade.Delete()

	ctx.Reply(discord.ReplyParams{
		Content:   fmt.Sprintf("You have accepted trade `%s` !", trade.ID),
		Ephemeral: true,
	})
	discord.Session.ChannelMessageSend(SenderDM.ID, fmt.Sprintf("Trade `%s` has been accepted by the user !", trade.ID))

}
