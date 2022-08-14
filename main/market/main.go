package market

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/ambelovsky/gosf"
)

var MarketMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("market", "127.0.0.1", config.MarketRPC, false)
	MarketMicroservice = gosf.GetMicroservice("market")
	go watchdog()
	fmt.Println("[MARKET] Initialized market microservice.")
}

func watchdog() {
	t := time.NewTicker(time.Second * 5)
	// check every ticks
	for {
		select {
		case <-t.C:
			// check if the microservice is up
			if !MarketMicroservice.Connected() {
				MarketMicroservice.Connect()
			}
		}
	}
}
