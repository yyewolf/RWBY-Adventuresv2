package websocket

import (
	"rwby-adventures/config"

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
	// data, found := GetToken(request)
	// if !found {
	// 	client.Disconnect()
	// 	return gosf.NewFailureMessage("f")
	// }
	// d := data.(*DungeonUserData)

	token, found := GetString(request, "token")
	if !found {
		return gosf.NewFailureMessage("f")
	}

	data, found := DungeonCache.Get(token)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}
	dungeon := data.(*DungeonStruct)

	d := &DungeonUserData{
		Dungeon: dungeon,
		User: &WebUser{
			Name: "Yewolf",
			ID:   "144472011924570113",
		},
		Token: "test",
		Host:  config.DungeonHost,
		Port:  config.DungeonWebsocket,
	}
	client.Join(d.Dungeon.ID)

	smallGrid := dungeon.Game.GetSmallGrid(3, 3)

	//fmt.Println("[WS] Arena connected!")
	return &gosf.Message{
		Success: true,
		Body: map[string]interface{}{
			"g": smallGrid,
		},
	}
}

func DungeonMove(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// data, found := GetToken(request)
	// if !found {
	// 	client.Disconnect()
	// 	return gosf.NewFailureMessage("f")
	// }
	// d := data.(*DungeonUserData)

	data, found := DungeonCache.Get("test")
	if !found {
		return gosf.NewFailureMessage("f")
	}
	dungeon := data.(*DungeonStruct)
	direction, found := GetInteger(request, "direction")
	if !found {
		return gosf.NewFailureMessage("no dir found")
	}

	dungeon.Game.MovePlayer(direction)

	smallGrid := dungeon.Game.GetSmallGrid(3, 3)

	return &gosf.Message{
		Success: true,
		Body: map[string]interface{}{
			"g": smallGrid,
		},
	}
}
