package cache

import (
	"sync"
	"time"

	"github.com/pmylund/go-cache"
	"github.com/yyewolf/gosf"
)

type WebUser struct {
	Name  string
	ID    string
	Token string
}

type User struct {
	Arena *Arena
	User  *WebUser
	Token string
}

type Player struct {
	// Useful for fast response
	Client    *gosf.Client
	LastClick time.Time
	Data      *User

	*sync.Mutex
}

type Arena struct {
	Players map[string]*Player `json:"-"`
	ID      string             `json:"-"`

	Name      string              `json:"title"`
	Img       string              `json:"image"`
	MaxHealth int                 `json:"-"`
	CurHealth int                 `json:"-"`
	Channel   chan int            `json:"-"`
	End       func(*Arena) string `json:"-"`
	Ticker    *time.Ticker        `json:"-"`
}

var Arenas *cache.Cache

func init() {
	Arenas = cache.New(8*time.Hour, 1*time.Hour)
}
