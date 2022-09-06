package websocket

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

func (dungeon *DungeonStruct) End() *discordgo.MessageEmbed {
	player := models.GetPlayer(dungeon.UserID)
	if player.IsNew {
		return &discordgo.MessageEmbed{
			Title: "Dungeon Recap",
			Color: config.Botcolor,
			Description: fmt.Sprintf("You have completed the dungeon but there has been an issue, send your rewards to an admin to get them : \n"+
				"Lien: **%d**\n"+
				"Character Box(es): **%d**\n"+
				"Arm(s): **%d**\n"+
				"Minion(s): **%d**",
				dungeon.Game.Rewards.Lien,
				dungeon.Game.Rewards.CCBox,
				dungeon.Game.Rewards.Arms,
				dungeon.Game.Rewards.Minions),
		}
	}

	player.Balance += int64(dungeon.Game.Rewards.Lien)
	player.Arms += dungeon.Game.Rewards.Arms
	player.Minions += dungeon.Game.Rewards.Minions
	player.Boxes.Boxes += dungeon.Game.Rewards.CCBox

	player.Save()
	player.Boxes.Save()

	return &discordgo.MessageEmbed{
		Title: "Dungeon Recap",
		Color: config.Botcolor,
		Description: fmt.Sprintf("You have completed the dungeon and received the following rewards: \n"+
			"Lien: **%d**\n"+
			"Character Box(es): **%d**\n"+
			"Arm(s): **%d**\n"+
			"Minion(s): **%d**",
			dungeon.Game.Rewards.Lien,
			dungeon.Game.Rewards.CCBox,
			dungeon.Game.Rewards.Arms,
			dungeon.Game.Rewards.Minions),
	}
}
