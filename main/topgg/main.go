package market

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/yyewolf/gosf"
)

var TopggMicroservice *gosf.Microservice

func StartTopGG() {
	gosf.RegisterMicroservice("topgg", "127.0.0.1", config.TopGGRPC, false)
	TopggMicroservice = gosf.GetMicroservice("topgg")
	TopggMicroservice.Listen("sendMessage", sendMessage)
	go watchdog()
	fmt.Println("[TOPGG] Initialized microservice.")
}

func watchdog() {
	t := time.NewTicker(time.Second * 10)
	// check every ticks
	for <-t.C; ; {
		// check if the microservice is up
		if !TopggMicroservice.Connected() {
			TopggMicroservice.Connect()
		}
	}
}
