package discord

import (
	"time"
)

type rateLimit struct {
	lastMessage time.Time
}

func checkUser(id string) bool {
	val, err := RateLimitCache.Get(id)
	if !err {
		rl := &rateLimit{
			lastMessage: time.Now(),
		}
		RateLimitCache.Set(id, rl, 0)
		return false
	}
	if time.Now().Sub(val.(*rateLimit).lastMessage) <= time.Duration(CommandRouter.RateLimit)*time.Millisecond {
		return true
	}
	rl := &rateLimit{
		lastMessage: time.Now(),
	}
	RateLimitCache.Set(id, rl, 0)
	return false
}
