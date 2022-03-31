package websocket

import (
	"fmt"
	"rwby-adventures/arenapc"
	"time"

	"github.com/ambelovsky/gosf"
)

func CreateArena(in *arenapc.CreateArenaReq) (b bool, loots string) {
	_, exists := ArenaCache.Get(in.GetId())
	if exists {
		return true, ""
	}
	arena := &ArenaStruct{
		Players:   make(map[string]*Player),
		ID:        in.GetId(),
		Name:      "Wumpus",
		Img:       "https://vultam.net/img/background/wumpus.png",
		MaxHealth: 100,
		CurHealth: 100,
		End:       EndClassicArena,
		Channel:   make(chan int),
	}
	ArenaCache.Set(in.GetId(), arena, 0)
	fmt.Println("[ARENA] Created arena:", in.GetId())
	return ArenaLoop(arena)
}

func ArenaLoop(arena *ArenaStruct) (b bool, loots string) {
	//Sends data to players
	t := time.NewTicker(time.Millisecond * 100)
	arena.Ticker = t
	for {
		select {
		case <-arena.Channel:
			return arena.End(arena)
		case <-arena.Ticker.C:
			break
		}
		//fmt.Println("[ARENA] Sending data to players")
		//No operations necessary if no one is in
		go gosf.Broadcast(arena.ID, "arenaLoop", &gosf.Message{
			Body: map[string]interface{}{
				"h": arena.CurHealth,
				"n": len(arena.Players),
			},
		})
	}
}
