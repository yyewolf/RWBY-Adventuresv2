package microservice

import (
	"fmt"
	"math/rand"
	"rwby-adventures/arenas_back/cache"
	"rwby-adventures/arenas_back/websocket"
	"rwby-adventures/microservices"

	uuid "github.com/satori/go.uuid"
	"github.com/yyewolf/gosf"
)

func createArena(client *gosf.Client, request *gosf.Request) *gosf.Message {
	client.Join("arena")

	var req microservices.CreateArena
	gosf.MapToStruct(request.Message.Body, &req)

	_, exists := cache.Arenas.Get(req.ID)
	if exists {
		return gosf.NewFailureMessage("already exists")
	}
	health := ((rand.Intn(25)+25)*50 + 50) * 5
	arena := &cache.Arena{
		Players:   make(map[string]*cache.Player),
		ID:        req.ID,
		Name:      "Grimm",
		Img:       "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/8bc5f32c-8fbf-4166-8729-b7d4d62cb35d/ddbvlct-86de9502-195a-4185-a02d-41432f9d19e9.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcLzhiYzVmMzJjLThmYmYtNDE2Ni04NzI5LWI3ZDRkNjJjYjM1ZFwvZGRidmxjdC04NmRlOTUwMi0xOTVhLTQxODUtYTAyZC00MTQzMmY5ZDE5ZTkucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.593heLTyYUe6rYeftcflrOjPyGVBnopJCEl_gfllkAo",
		MaxHealth: health,
		CurHealth: health,
		End:       EndClassicArena,
		Channel:   make(chan int),
	}
	cache.Arenas.Set(req.ID, arena, 0)
	fmt.Println("[ARENA] Created arena:", req.ID)

	go func() {
		loots := websocket.ArenaLoop(arena)

		msg := gosf.NewSuccessMessage()
		msg.Body = gosf.StructToMap(microservices.EndArena{
			ChannelID: req.ChannelID,
			Message:   loots,
		})
		msg.Text = uuid.NewV4().String()
		gosf.Broadcast("arena", "endArena", msg)
	}()

	return gosf.NewSuccessMessage("starting the arena")
}
