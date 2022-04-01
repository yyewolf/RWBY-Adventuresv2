package rwby_grpc

import (
	"context"
	"flag"
	"fmt"
	"rwby-adventures/arenapc"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "127.0.0.1:50002"
)

var ArenaServer arenapc.ArenaClient

func ConnectToRPC() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("[ARENA] Cannot connect to arena grpc")
	}
	//defer conn.Close()
	ArenaServer = arenapc.NewArenaClient(conn)
	fmt.Println("[ARENA] Connected to arena grpc")

	go RPCWatchdog()
}

func RPCWatchdog() {
	for {
		time.Sleep(time.Second * 5)
		_, err := ArenaServer.Ping(context.Background(), &arenapc.PingReq{})
		if err != nil {
			fmt.Println("[ARENA] Reconnecting to RPC")
			ConnectToRPC()
			return
		}
	}
}
