package websocket

import (
	"time"

	"github.com/pmylund/go-cache"
)

var Tokens *cache.Cache

type Token struct {
	Secret string
	UserID string
}

func init() {
	Tokens = cache.New(5*time.Minute, 10*time.Minute)
}
