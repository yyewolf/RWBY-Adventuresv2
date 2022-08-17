package microservice

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

func init() {
	// Listen on an endpoint
	gosf.Listen("createArena", createArena)
	// gosf.Listen("createAuction", createAuction)
}

func CreateMicroservice() {
	// Start the server using a basic configuration
	go gosf.Startup(map[string]interface{}{"port": config.ArenaRPC})
	fmt.Println("[ARENA] Microservice UP.")
}
