package websocket

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/models"
	"time"
)

func EndClassicArena(arena *ArenaStruct) (b bool, loots string) {
	// We disconnect everyone from the websocket
	for _, player := range arena.Players {
		player.Client.Disconnect()
	}

	rewards := ""
	// We give rewards to all players
	for _, player := range arena.Players {
		p := models.GetPlayer(player.Data.User.ID)
		p.Status.LastXP = time.Now().Unix()
		XP := p.CalcSelectedXP(3, p.Shop.XPBoost)
		lvl := p.GiveSelectedXP(XP)
		if lvl {
			rewards += fmt.Sprintf("<@%s> has earned %d XP (Level Up !).\n", p.DiscordID, XP)
		} else {
			rewards += fmt.Sprintf("<@%s> has earned %d XP.\n", p.DiscordID, XP)
		}
		// We give the player the reward
		config.Database.Save(p.Status)
		if p.SelectedChar != nil {
			config.Database.Save(p.SelectedChar)
		}
		if p.SelectedGrimm != nil {
			config.Database.Save(p.SelectedGrimm)
		}
	}

	// We delete the arena from the cache
	ArenaCache.Delete(arena.Name)

	return false, rewards
}
