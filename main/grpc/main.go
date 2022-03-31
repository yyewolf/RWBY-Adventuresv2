package rwby_grpc

import (
	"flag"
	"fmt"
	"rwby-adventures/arenapc"

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
}
