package microservice

import (
	"fmt"

	"github.com/yyewolf/gosf"
)

func joinRoom(client *gosf.Client, request *gosf.Request) {
	fmt.Println("[DUNGEONS] Client connected to dungeons")
	client.Join("dungeons")
	roomActive = true
}
