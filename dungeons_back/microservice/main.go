package microservice

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/ambelovsky/gosf"
)

func init() {
	// Listen on an endpoint
	gosf.Listen("createDungeon", createDungeon)
}

func CreateMicroservice() {
	// Start the server using a basic configuration
	go gosf.Startup(map[string]interface{}{"port": config.DungeonRPC})
	fmt.Println("[DUNGEON] Microservice UP.")
}
