package microservice

import (
	"fmt"

	"github.com/yyewolf/gosf"
)

func joinRoom(client *gosf.Client, request *gosf.Request) {
	fmt.Println("[TOPGG] Client connected to TOPGG")
	client.Join("topgg")
	roomActive = true
}
