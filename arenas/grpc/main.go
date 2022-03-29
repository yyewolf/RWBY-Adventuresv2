package agrpc

import (
	"context"
	"fmt"
	"net"
	"rwby-adventures/arenapc"

	"google.golang.org/grpc"
)

var (
	// addr = "main"
	port = "50001"
)

type server struct {
	arenapc.UnimplementedArenaServer
}

func (s *server) CreateArena(ctx context.Context, in *arenapc.CreateArenaReq) (*arenapc.CreateArenaRep, error) {
	return &arenapc.CreateArenaRep{Status: 1}, nil

}

func (s *server) JoinArena(ctx context.Context, in *arenapc.JoinArenaReq) (*arenapc.JoinArenaRep, error) {
	return &arenapc.JoinArenaRep{Ok: 1}, nil
}

func Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
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
