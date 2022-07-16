package websocket

import (
	"rwby-adventures/dungeons_back/game"
	"rwby-adventures/microservices"
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
	ID   string
	Game *game.Dungeon

	End     func(*DungeonStruct) *microservices.DungeonEndResponse
	Ticker  *time.Ticker
	Channel chan int
}

var DungeonCache = cache.New(5*time.Hour, 10*time.Minute)
