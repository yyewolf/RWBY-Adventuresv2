package gambles

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/yyewolf/gosf"
)

var GambleMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("gambles", config.GambleRPCHost, config.GambleRPC, false)
	GambleMicroservice = gosf.GetMicroservice("gambles")
	go watchdog()
	fmt.Println("[GAMBLES] Initialized microservice.")
}

func watchdog() {
	t := time.NewTicker(time.Second * 10)
	// check every ticks
	for <-t.C; ; {
		// check if the microservice is up
		if !GambleMicroservice.Connected() {
			GambleMicroservice.Connect()
		}
	}
}
