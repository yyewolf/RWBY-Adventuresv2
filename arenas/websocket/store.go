package websocket

import (
	"sync"
	"time"

	"github.com/ambelovsky/gosf"
	"github.com/pmylund/go-cache"
)

type Player struct {

	// Useful for fast response
	Client    *gosf.Client
	LastClick time.Time
	Data      *ArenaUserData

	*sync.Mutex
}

type ArenaStruct struct {
	Players map[string]*Player
	ID      string

	Name      string
	Img       string
	MaxHealth int
	CurHealth int
	Channel   chan int
	End       func(*ArenaStruct) (bool, string)
}

var ArenaCache = cache.New(5*time.Hour, 10*time.Minute)
