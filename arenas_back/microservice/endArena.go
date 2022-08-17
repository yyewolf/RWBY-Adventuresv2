package microservice

import (
	"fmt"
	"math/rand"
	"rwby-adventures/arenas_back/cache"
	"rwby-adventures/config"
	"rwby-adventures/models"
	"time"
)

func EndClassicArena(arena *cache.Arena) (loots string) {
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
		money := rand.Intn(300) + 54
		if lvl {
			rewards += fmt.Sprintf("`%s has earned %d XP (Level Up !) and %dⱠ.`\n", player.Data.User.Name, XP, money)
		} else {
			rewards += fmt.Sprintf("`%s has earned %d XP and %dⱠ.`\n", player.Data.User.Name, XP, money)
		}
		// We give the player the reward
		config.Database.Save(p.Status)
		p.Balance += int64(money)
		config.Database.Save(p.Balance)
		if p.SelectedChar != nil {
			config.Database.Save(p.SelectedChar)
		}
		if p.SelectedGrimm != nil {
			config.Database.Save(p.SelectedGrimm)
		}
	}

	// We delete the arena from the cache
	cache.Arenas.Delete(arena.ID)

	return rewards
}
