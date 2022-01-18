package websocket

import (
	"fmt"

	"github.com/ambelovsky/gosf"
)

func echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	return gosf.NewSuccessMessage(request.Message.Text)
}

func StartWebsocket() {
	// Start the server using a basic configuration
	gosf.Listen("echo", echo)
	go gosf.Startup(map[string]interface{}{"port": 9999})
	fmt.Println("[WS] Started.")
}
