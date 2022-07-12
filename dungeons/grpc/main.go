package dgrpc

import (
	"context"
	"fmt"
	"net"
	"rwby-adventures/config"
	"rwby-adventures/dungeonpc"
	"rwby-adventures/dungeons/websocket"

	"google.golang.org/grpc"
)

type server struct {
	dungeonpc.UnimplementedDungeonServer
}

func (s *server) CreateDungeon(ctx context.Context, in *dungeonpc.CreateDungeonReq) (*dungeonpc.CreateDungeonRep, error) {
	err, reward := websocket.CreateDungeon(in)
	if err {
		return &dungeonpc.CreateDungeonRep{Status: 1}, nil
	} else {
		return &dungeonpc.CreateDungeonRep{Status: 0, Loots: reward}, nil
	}
}

func (s *server) Ping(ctx context.Context, in *dungeonpc.PingReq) (*dungeonpc.PingRep, error) {
	return &dungeonpc.PingRep{}, nil
}

func Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.DungeonRPC))
	if err != nil {
		fmt.Printf("[DUNGEON] Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	dungeonpc.RegisterDungeonServer(s, &server{})
	fmt.Printf("[DUNGEON] Listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("[DUNGEON] Failed to serve: %v\n", err)
	}
}
