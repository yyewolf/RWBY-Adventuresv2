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
	Data      *DungeonUserData

	*sync.Mutex
}

type DungeonStruct struct {
	ID string

	End     func(*DungeonStruct) (bool, string)
	Ticker  *time.Ticker
	Channel chan int
	// TO DO
}

var DungeonCache = cache.New(5*time.Hour, 10*time.Minute)
