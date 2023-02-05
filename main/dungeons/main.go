package dungeons

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/yyewolf/gosf"
)

var DungeonsMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("dungeons", config.DungeonRPCHost, config.DungeonRPC, false)
	DungeonsMicroservice = gosf.GetMicroservice("dungeons")
	DungeonsMicroservice.Listen("sendMessage", sendMessage)
	go watchdog()
	fmt.Println("[DUNGEONS] Initialized microservice.")
}

func watchdog() {
	t := time.NewTicker(time.Second * 5)
	// check every ticks
	for {
		select {
		case <-t.C:
			// check if the microservice is still alive
			if !DungeonsMicroservice.Connected() {
				DungeonsMicroservice.Connect()
			}
		}
	}
}
