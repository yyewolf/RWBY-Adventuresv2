package microservice

import (
	"fmt"
	"rwby-adventures/microservices"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/yyewolf/gosf"
)

func SendMessageToBot(message *microservices.MarketMessage) {
	msg := &gosf.Message{
		Text: uuid.NewV4().String(),
	}
	msg.Body = gosf.StructToMap(message)
	fmt.Println("[MARKET] Waiting for room to be active...")
	for !roomActive {
		time.Sleep(1 * time.Second)
	}
	gosf.Broadcast("market", "sendMessage", msg)
}
