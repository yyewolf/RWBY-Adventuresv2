package main

import (
	"os"
	"os/signal"
	agrpc "rwby-adventures/arenas/grpc"
	"rwby-adventures/arenas/web"
	"rwby-adventures/arenas/websocket"
	"syscall"
)

func main() {
	go agrpc.Serve()
	web.StartWeb()
	websocket.StartWebsocket()

	// websocket.CreateArena(&arenapc.CreateArenaReq{
	// 	Id: "main",
	// })

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
