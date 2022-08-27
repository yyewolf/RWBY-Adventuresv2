package microservice

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

func init() {
	gosf.OnConnect(joinRoom)
}

func CreateMicroservice() {
	// Start the server using a basic configuration
	go gosf.Startup(map[string]interface{}{"port": config.TopGGRPC})
	fmt.Println("[TOPGG] Microservice UP.")
}
