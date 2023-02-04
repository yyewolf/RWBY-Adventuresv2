package arenas

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

var ArenaMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("arenas", config.ArenaRPCHost, config.ArenaRPC, false)
	ArenaMicroservice = gosf.GetMicroservice("arenas")
	// go watchdog()
	fmt.Println("[ARENAS] Initialized microservice.")
}

func watchdog() {
	// t := time.NewTicker(time.Second * 10)
	// // check every ticks
	// for <-t.C; ; {
	// 	// check if the microservice is up
	// 	if !ArenaMicroservice.Connected() {
	// 		ArenaMicroservice.Connect()
	// 	}
	// }
}
