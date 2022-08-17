package microservice

import (
	"fmt"
	"rwby-adventures/arenas_back/cache"
	"rwby-adventures/arenas_back/websocket"
	"rwby-adventures/microservices"

	"github.com/yyewolf/gosf"
)

func createArena(client *gosf.Client, request *gosf.Request) *gosf.Message {
	client.Join("arena")

	var req microservices.CreateArena
	gosf.MapToStruct(request.Message.Body, &req)

	_, exists := cache.Arenas.Get(req.ID)
	if exists {
		return gosf.NewFailureMessage("already exists")
	}
	arena := &cache.Arena{
		Players:   make(map[string]*cache.Player),
		ID:        req.ID,
		Name:      "Wumpus",
		Img:       "https://vultam.net/img/background/wumpus.png",
		MaxHealth: 100,
		CurHealth: 100,
		End:       EndClassicArena,
		Channel:   make(chan int),
	}
	cache.Arenas.Set(req.ID, arena, 0)
	fmt.Println("[ARENA] Created arena:", req.ID)

	go func() {
		loots := websocket.ArenaLoop(arena)

		msg := gosf.NewSuccessMessage()
		msg.Body = gosf.StructToMap(microservices.EndArena{
			ChannelID: req.ChannelID,
			Message:   loots,
		})
		gosf.Broadcast("arena", "endArena", msg)
	}()

	return gosf.NewSuccessMessage("starting the arena")
}
