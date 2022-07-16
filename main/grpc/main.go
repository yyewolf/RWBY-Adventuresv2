package rwby_grpc

import (
	"context"
	"flag"
	"fmt"
	arenapc "rwby-adventures/arenas_rpc"
	"rwby-adventures/config"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addrArena = fmt.Sprintf("%s:%s", config.ArenaHost, config.ArenaRPC)
)

var ArenaServer arenapc.ArenaClient

func ConnectToRPC() {
	go ConnectToArena()
}

func ConnectToArena() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(addrArena, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("[ARENA] Cannot connect to arena grpc")
	}
	//defer conn.Close()
	ArenaServer = arenapc.NewArenaClient(conn)
	//fmt.Println("[ARENA] Connected to arena grpc")

	go arenaWatchdog()
}

func arenaWatchdog() {
	for {
		time.Sleep(time.Second * 5)
		_, err := ArenaServer.Ping(context.Background(), &arenapc.PingReq{})
		if err != nil {
			fmt.Println("[ARENA] Reconnecting to RPC")
			ConnectToArena()
			return
		}
	}
}
