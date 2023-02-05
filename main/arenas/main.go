package arenas

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/yyewolf/gosf"
)

var ArenaMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("arenas", config.ArenaRPCHost, config.ArenaRPC, false)
	ArenaMicroservice = gosf.GetMicroservice("arenas")
	go watchdog()
	fmt.Println("[ARENAS] Initialized microservice.")
}

func watchdog() {
	t := time.NewTicker(time.Second * 10)
	// check every ticks
	for {
		select {
		case <-t.C:
			if !ArenaMicroservice.Connected() {
				ArenaMicroservice.Connect()
			}
		}
	}
}
