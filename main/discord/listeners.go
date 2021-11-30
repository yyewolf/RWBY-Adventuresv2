package discord

import (
	"rwby-adventures/models"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pmylund/go-cache"
)

type Listener struct {
	Cache    *cache.Cache
	Callback func(*ListenerContext)
}

type ListenerContext struct {
	s *discordgo.Session

	ID        string
	ChannelID string
	GuildID   string

	Author *discordgo.User
	Player *models.Player

	Data interface{}

	Message *discordgo.Message
}

func isChannelInCache(id string) bool {
	_, found := ListenersCache.Get(id)
	return found
}

func getDataFromCache(id string) (interface{}, time.Time, bool, func(*ListenerContext)) {
	val, found := ListenersCache.Get(id)
	if !found {
		return val, time.Time{}, found, nil
	}
	d, expire, found := val.(Listener).Cache.GetWithExpiration(id)
	return d, expire, found, val.(Listener).Callback
}

func addDataToCache(id string, cache *cache.Cache, callback func(*ListenerContext), data interface{}) {
	ListenersCache.Set(id, Listener{
		Cache:    cache,
		Callback: callback,
	}, 0)
	cache.Set(id, data, 0)
}

func (l *ListenerContext) reply(p ReplyParams) (st *discordgo.Message, err error) {
	ctx := &CmdContext{
		Session:   l.s,
		ID:        l.ID,
		ChannelID: l.ChannelID,
		GuildID:   l.GuildID,
		Author:    l.Author,
		Player:    l.Player,
		Message:   l.Message,
	}
	return ctx.Reply(p)
}
