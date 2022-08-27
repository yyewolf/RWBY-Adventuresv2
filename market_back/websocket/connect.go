package websocket

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

func connect(client *gosf.Client, request *gosf.Request) *gosf.Message {
	token, exists := GetToken(request)
	if !exists {
		return gosf.NewFailureMessage("Invalid token")
	}
	t := token.(*Token)

	msg := gosf.NewSuccessMessage()
	if t.Empty {
		link := fmt.Sprintf("http://%s%s/login/%s", config.MarketHost, config.MarketPort, t.Token)
		msg.Body = map[string]interface{}{
			"link":      link,
			"connected": false,
		}
	} else {
		msg.Body = map[string]interface{}{
			"connected": true,
		}
	}
	return msg
}
