package microservice

import (
	"github.com/yyewolf/gosf"
)

func joinRoom(client *gosf.Client, request *gosf.Request) {
	client.Join("market")
	roomActive = true
}
