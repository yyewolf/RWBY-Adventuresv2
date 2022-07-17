package microservice

import (
	"fmt"
	"rwby-adventures/dungeons_back/game"
	"rwby-adventures/dungeons_back/websocket"
	"rwby-adventures/microservices"
	"time"

	"github.com/ambelovsky/gosf"
)

func createDungeon(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// Convert to our nicer format
	var req microservices.DungeonCreateRequest
	gosf.MapToStruct(request.Message.Body, &req)

	ID := req.ID

	_, exists := websocket.DungeonCache.Get(ID)
	if exists {
		return gosf.NewFailureMessage("already exists")
	}
	d := game.NewDungeon(15, 15)
	dungeon := &websocket.DungeonStruct{
		ID:    ID,
		Game:  d,
		EndIt: make(chan int),
	}
	websocket.DungeonCache.Set(ID, dungeon, 0)

	fmt.Println("[DUNGEONS] Created dungeon with ID:", ID)

	data := dungeonLoop(dungeon)

	msg := gosf.NewSuccessMessage("finished")
	msg.Body = gosf.StructToMap(data)

	return msg
}

func dungeonLoop(dungeon *websocket.DungeonStruct) *microservices.DungeonEndResponse {
	//Sends data to players
	t := time.NewTicker(time.Millisecond * 100)
	dungeon.Ticker = t
	for {
		select {
		case <-dungeon.EndIt:
			return dungeon.End()
		case <-dungeon.Ticker.C:
		}
	}
}
