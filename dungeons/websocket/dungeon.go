package websocket

import (
	"github.com/ambelovsky/gosf"
)

type WebUser struct {
	Name  string
	ID    string
	Token string
}

type DungeonUserData struct {
	Dungeon *DungeonStruct
	User    *WebUser
	Token   string
	Host    string
	Port    int
}

func DungeonConnect(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := GetToken(request)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}
	d := data.(*DungeonUserData)
	client.Join(d.Dungeon.ID)

	//fmt.Println("[WS] Arena connected!")
	return gosf.NewSuccessMessage()
}
