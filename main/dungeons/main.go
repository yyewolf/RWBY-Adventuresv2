package dungeons

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/ambelovsky/gosf"
)

var DungeonsMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("dungeons", "127.0.0.1", config.DungeonRPC, false)
	DungeonsMicroservice = gosf.GetMicroservice("dungeons")
	go watchdog()
	fmt.Println("[DUNGEONS] Initialized dungeons microservice.")
}

func watchdog() {
	t := time.NewTicker(time.Second * 5)
	// check every ticks
	for {
		select {
		case <-t.C:
			// check if the microservice is up
			if !DungeonsMicroservice.Connected() {
				DungeonsMicroservice.Connect()
			}
		}
	}
}
