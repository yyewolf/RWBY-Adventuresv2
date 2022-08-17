package microservice

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

func init() {
	// Listen on an endpoint
	gosf.Listen("createListing", createListing)
	gosf.Listen("createAuction", createAuction)
	gosf.OnConnect(joinRoom)
}

func CreateMicroservice() {
	// Start the server using a basic configuration
	go gosf.Startup(map[string]interface{}{"port": config.MarketRPC})
	fmt.Println("[MARKET] Microservice UP.")
}
