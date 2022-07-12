package agrpc

import (
	"context"
	"fmt"
	"net"
	"rwby-adventures/arenapc"
	"rwby-adventures/arenas/websocket"
	"rwby-adventures/config"

	"google.golang.org/grpc"
)

type server struct {
	arenapc.UnimplementedArenaServer
}

func (s *server) CreateArena(ctx context.Context, in *arenapc.CreateArenaReq) (*arenapc.CreateArenaRep, error) {
	err, reward := websocket.CreateArena(in)
	if err {
		return &arenapc.CreateArenaRep{Status: 1}, nil
	} else {
		return &arenapc.CreateArenaRep{Status: 0, Loots: reward}, nil
	}
}

func (s *server) Ping(ctx context.Context, in *arenapc.PingReq) (*arenapc.PingRep, error) {
	return &arenapc.PingRep{}, nil
}

func Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.ArenaRPC))
	if err != nil {
		fmt.Printf("[ARENA] Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	arenapc.RegisterArenaServer(s, &server{})
	fmt.Printf("[ARENA] Listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("[ARENA] Failed to serve: %v\n", err)
	}
}
