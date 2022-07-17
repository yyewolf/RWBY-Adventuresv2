package microservice

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/dungeons_back/game"
	"rwby-adventures/dungeons_back/websocket"

	"github.com/ambelovsky/gosf"
)

func init() {
	// Listen on an endpoint
	gosf.Listen("createDungeon", createDungeon)

	fmt.Println("[WS] Started.")

	d := game.NewDungeon(15, 15)
	dungeon := &websocket.DungeonStruct{
		ID:    "test",
		Game:  d,
		EndIt: make(chan int),
	}
	websocket.DungeonCache.Set("test", dungeon, 0)

	go func() {
		dungeonLoop(dungeon)
	}()
}

func CreateMicroservice() {
	// Start the server using a basic configuration
	go gosf.Startup(map[string]interface{}{"port": config.DungeonRPC})
	fmt.Println("[DUNGEON] Microservice UP.")
}
