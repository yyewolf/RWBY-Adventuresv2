package microservice

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

func init() {
	// Listen on an endpoint
	gosf.Listen("createDungeon", createDungeon)
	gosf.OnConnect(joinRoom)

	// d := game.NewDungeon(15, 15)
	// dungeon := &websocket.DungeonStruct{
	// 	ID:    "test",
	// 	Game:  d,
	// 	EndIt: make(chan int),
	// }
	// websocket.DungeonCache.Set("test", dungeon, 0)

	// for y := 0; y < d.Height; y++ {
	// 	for x := 0; x < d.Width; x++ {
	// 		fmt.Print(d.Grid[y][x].Type)
	// 	}
	// 	fmt.Println()
	// }

	// go func() {
	// 	dungeonLoop(dungeon)
	// }()
}

func CreateMicroservice() {
	// Start the server using a basic configuration
	go gosf.Startup(map[string]interface{}{"port": config.DungeonRPC})
	fmt.Println("[DUNGEON] Microservice UP.")
}
