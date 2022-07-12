package websocket

import (
	"time"

	"github.com/ambelovsky/gosf"
	"github.com/pmylund/go-cache"
)

type Token struct {
	Secret string
	UserID string
}

func init() {
	Tokens = cache.New(5*time.Minute, 10*time.Minute)
}

func GetToken(request *gosf.Request) (data interface{}, found bool) {
	token, ok := GetString(request, "token")
	if !ok {
		return nil, false
	}
	return Tokens.Get(token)
}
