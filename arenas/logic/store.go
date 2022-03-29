package logic

import (
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"
)

type Player struct {

	// Useful for fast response
	Socket    *websocket.Conn
	LastClick time.Time

	*sync.Mutex
}

type ArenaStruct struct {
	Players map[string]*Player

	Special    bool
	ChannelID  string
	GuildID    string
	Active     bool
	Name       string
	NameText   string
	DamageText string
	Img        string
	MaxHealth  int
	CurHealth  int
	Channel    chan int
	SpawnedAt  time.Time
	End        func(*ArenaStruct)
	state      *discordgo.Session
}
