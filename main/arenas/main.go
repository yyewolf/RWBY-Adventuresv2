package arenas

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

var ArenaMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("arenas", "127.0.0.1", config.ArenaRPC, false)
	ArenaMicroservice = gosf.GetMicroservice("arenas")
	fmt.Println("[ARENAS] Initialized microservice.")
}
