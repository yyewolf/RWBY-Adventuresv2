package microservice

import (
	"fmt"
	"rwby-adventures/microservices"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/yyewolf/gosf"
)

var waitAfterCon bool
var roomActive bool

func SendMessageToBot(message *microservices.TopGGMessage) {
	msg := &gosf.Message{
		Text: uuid.NewV4().String(),
	}
	msg.Body = gosf.StructToMap(message)
	fmt.Println("[TOPGG] Waiting for room to be active...")
	for !roomActive {
		time.Sleep(1 * time.Second)
		waitAfterCon = true
	}
	if waitAfterCon {
		time.Sleep(5 * time.Second)
	}
	gosf.Broadcast("market", "sendMessage", msg)
}
