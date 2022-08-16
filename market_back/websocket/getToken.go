package websocket

import (
	"crypto/sha256"
	"fmt"

	"github.com/ambelovsky/gosf"
	"github.com/pmylund/go-cache"
	uuid "github.com/satori/go.uuid"
)

func getToken(client *gosf.Client, request *gosf.Request) *gosf.Message {
	token := fmt.Sprintf("%x", sha256.Sum256(uuid.NewV4().Bytes()))
	msg := gosf.NewSuccessMessage()
	msg.Body = map[string]interface{}{
		"token": token,
	}
	t := &Token{
		Empty: true,
		Token: token,
	}
	Tokens.Set(token, t, cache.DefaultExpiration)
	return msg
}
