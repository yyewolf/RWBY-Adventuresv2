package websocket

import (
	"time"

	"github.com/ambelovsky/gosf"
)

type WebUser struct {
	Name  string
	ID    string
	Token string
}

type ArenaUserData struct {
	Arena *ArenaStruct
	User  *WebUser
	Token string
	Host  string
	Port  int
}

func ArenaConnect(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := GetToken(request)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}
	d := data.(*ArenaUserData)
	client.Join(d.Arena.ID)

	if val, ok := d.Arena.Players[d.User.ID]; ok {
		val.Client.Disconnect()
	}
	d.Arena.Players[d.User.ID] = &Player{
		Client:    client,
		Data:      d,
		LastClick: time.Now(),
	}
	//fmt.Println("[WS] Arena connected!")
	return gosf.NewSuccessMessage()
}

func ArenaHit(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := GetToken(request)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}
	d := data.(*ArenaUserData)
	if d.Arena.CurHealth > 0 {
		d.Arena.CurHealth -= 50
	}
	if d.Arena.CurHealth <= 0 {
		d.Arena.CurHealth = 0
		time.Sleep(100 * time.Millisecond)
		// We stop the loop
		d.Arena.Channel <- 1
	}
	//fmt.Println("[WS] Arena hit!")
	return gosf.NewSuccessMessage()
}
