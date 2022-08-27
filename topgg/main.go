package main

import (
	"os"
	"os/signal"
	"rwby-adventures/topgg/microservice"
	"syscall"
)

func main() {
	microservice.CreateMicroservice()
	// websocket.CreateArena(&arenapc.CreateArenaReq{
	// 	Id: "main",
	// })

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
