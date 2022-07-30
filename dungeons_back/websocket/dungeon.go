package websocket

import (
	"rwby-adventures/config"
	"rwby-adventures/dungeons_back/game"
	"rwby-adventures/microservices"

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

type WebsocketUpdate struct {
	Grid    [][]*game.DungeonCell        `json:"g"`
	Health  int                          `json:"h"`
	Ended   bool                         `json:"e"`
	Win     bool                         `json:"w"`
	Rewards *microservices.DungeonReward `json:"r"`
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

	if dungeon.Ended {
		return &gosf.Message{
			Success: true,
			Body: gosf.StructToMap(&WebsocketUpdate{
				Grid:    smallGrid,
				Health:  dungeon.Game.Health,
				Ended:   true,
				Rewards: dungeon.Game.Rewards,
			}),
		}
	}

	//fmt.Println("[WS] Arena connected!")
	return &gosf.Message{
		Success: true,
		Body: gosf.StructToMap(&WebsocketUpdate{
			Grid:   smallGrid,
			Health: dungeon.Game.Health,
		}),
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
	if dungeon.Ended {
		return gosf.NewFailureMessage("f")
	}

	direction, found := GetInteger(request, "direction")
	if !found {
		return gosf.NewFailureMessage("no dir found")
	}

	end := dungeon.Game.MovePlayer(direction)
	smallGrid := dungeon.Game.GetSmallGrid(3, 3)
	if end {
		dungeon.EndIt <- 1
		dungeon.Ended = true

		if dungeon.Game.Health <= 0 {
			dungeon.Game.Rewards.Lien /= 2
		}

		return &gosf.Message{
			Success: true,
			Body: gosf.StructToMap(&WebsocketUpdate{
				Grid:    smallGrid,
				Health:  dungeon.Game.Health,
				Ended:   true,
				Rewards: dungeon.Game.Rewards,
				Win:     dungeon.Game.Win,
			}),
		}
	}

	return &gosf.Message{
		Success: true,
		Body: gosf.StructToMap(&WebsocketUpdate{
			Grid:   smallGrid,
			Health: dungeon.Game.Health,
		}),
	}
}

func AmbrosiusChoice(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := DungeonCache.Get("test")
	if !found {
		return gosf.NewFailureMessage("f")
	}
	dungeon := data.(*DungeonStruct)
	if dungeon.Ended {
		return gosf.NewFailureMessage("f")
	}

	choice, found := GetInteger(request, "choice")
	if !found {
		return gosf.NewFailureMessage("no choice found")
	}

	return dungeon.Game.MakeChoice(choice)
}
