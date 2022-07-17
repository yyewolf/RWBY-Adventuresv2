package websocket

import (
	"rwby-adventures/dungeons_back/game"
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
	ID     string
	Game   *game.Dungeon
	Ticker *time.Ticker
	Ended  bool
	EndIt  chan int
}

var DungeonCache = cache.New(5*time.Hour, 10*time.Minute)
