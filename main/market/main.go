package market

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/yyewolf/gosf"
)

var MarketMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("market", "127.0.0.1", config.MarketRPC, false)
	MarketMicroservice = gosf.GetMicroservice("market")
	MarketMicroservice.Listen("sendMessage", sendMessage)
	go watchdog()
	fmt.Println("[MARKET] Initialized microservice.")
}

func watchdog() {
	t := time.NewTicker(time.Second * 10)
	// check every ticks
	for <-t.C; ; {
		// check if the microservice is up
		if !MarketMicroservice.Connected() {
			MarketMicroservice.Connect()
		}
	}
}
