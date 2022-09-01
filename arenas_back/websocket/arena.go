package websocket

import (
	"fmt"
	"rwby-adventures/arenas_back/cache"
	"time"

	"github.com/yyewolf/gosf"
)

func ArenaConnect(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := GetToken(request)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}
	d := data.(*cache.User)
	client.Join(d.Arena.ID)
	fmt.Println("[WS] Client connected!")

	// if val, ok := d.Arena.Players[d.User.ID]; ok {
	// 	val.Client.Disconnect()
	// }
	d.Arena.Players[d.User.ID] = &cache.Player{
		Client:    client,
		Data:      d,
		LastClick: time.Now(),
	}
	//fmt.Println("[WS] Arena connected!")
	msg := gosf.NewSuccessMessage()
	msg.Body = map[string]interface{}{
		"username": d.User.Name,
		"arena":    d.Arena,
	}
	return msg
}

func ArenaHit(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := GetToken(request)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}
	d := data.(*cache.User)
	if d.Arena.CurHealth > 0 {
		d.Arena.CurHealth -= 50
	}
	if d.Arena.CurHealth <= 0 {
		d.Arena.CurHealth = 0
		time.Sleep(100 * time.Millisecond)
		// We stop the loop
		d.Arena.Channel <- 1
	}
	//fmt.Println("[WS] Arena hit!")
	return gosf.NewSuccessMessage()
}

func ArenaLoop(arena *cache.Arena) (loots string) {
	//Sends data to players
	t := time.NewTicker(time.Millisecond * 100)
	arena.Ticker = t
	for {
		select {
		case <-arena.Channel:
			return arena.End(arena)
		case <-arena.Ticker.C:
		}
		//fmt.Println("[ARENA] Sending data to players")
		//No operations necessary if no one is in
		go gosf.Broadcast(arena.ID, "arenaLoop", &gosf.Message{
			Body: map[string]interface{}{
				"h": arena.CurHealth * 100.0 / arena.MaxHealth,
				"n": len(arena.Players),
			},
		})
	}
}
