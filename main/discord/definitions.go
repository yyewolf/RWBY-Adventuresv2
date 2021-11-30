package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/bwmarrin/snowflake"
	"github.com/pmylund/go-cache"
)

// Useful
var StartTime = time.Now()

// DGO
var Session *discordgo.Session
var CommandRouter router
var DefaultFooter *discordgo.MessageEmbedFooter

// Caches
var RateLimitCache *cache.Cache
var ActiveMenus *cache.Cache
var ListenersCache *cache.Cache

// UUID
var Node *snowflake.Node

func init() {
	RateLimitCache = cache.New(5*time.Minute, 10*time.Minute)
	ActiveMenus = cache.New(5*time.Minute, 10*time.Minute)
	ListenersCache = cache.New(5*time.Minute, 10*time.Minute)
	HelpMenus = make(map[string][]*Command)
	Node, _ = snowflake.NewNode(1)
}
