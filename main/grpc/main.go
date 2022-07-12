package rwby_grpc

import (
	"context"
	"flag"
	"fmt"
	"rwby-adventures/arenapc"
	"rwby-adventures/config"
	"rwby-adventures/dungeonpc"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addrArena   = fmt.Sprintf("%s:%s", config.ArenaHost, config.ArenaRPC)
	addrDungeon = fmt.Sprintf("%s:%s", config.DungeonHost, config.DungeonRPC)
)

var ArenaServer arenapc.ArenaClient
var DungeonServer dungeonpc.DungeonClient

func ConnectToRPC() {
	go ConnectToArena()
	go ConnectToDungeon()
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

func ConnectToDungeon() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(addrDungeon, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("[DUNGEON] Cannot connect to dungeon grpc")
	}
	//defer conn.Close()
	DungeonServer = dungeonpc.NewDungeonClient(conn)
	//fmt.Println("[DUNGEON] Connected to dungeon grpc")

	go dungeonWatchdog()
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

func dungeonWatchdog() {
	for {
		time.Sleep(time.Second * 5)
		_, err := DungeonServer.Ping(context.Background(), &dungeonpc.PingReq{})
		if err != nil {
			fmt.Println("[DUNGEON] Reconnecting to RPC")
			ConnectToDungeon()
			return
		}
	}
}
