package microservice

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/ambelovsky/gosf"
)

func init() {
	// Listen on an endpoint
	gosf.Listen("createListing", createListing)
	gosf.Listen("createAuction", createAuction)

	fmt.Println("[WS] Started.")
}

func CreateMicroservice() {
	// Start the server using a basic configuration
	go gosf.Startup(map[string]interface{}{"port": config.MarketRPC})
	fmt.Println("[MARKET] Microservice UP.")
}
