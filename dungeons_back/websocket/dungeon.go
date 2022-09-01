package websocket

import (
	"rwby-adventures/dungeons_back/game"
	"rwby-adventures/microservices"

	"github.com/yyewolf/gosf"
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

	data, found := Tokens.Get(token)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}
	d := data.(*DungeonUserData)

	client.Join(d.Dungeon.ID)

	smallGrid := d.Dungeon.Game.GetSmallGrid(3, 3)

	if d.Dungeon.Ended {
		return &gosf.Message{
			Success: true,
			Body: gosf.StructToMap(&WebsocketUpdate{
				Grid:    smallGrid,
				Health:  d.Dungeon.Game.Health,
				Ended:   true,
				Rewards: d.Dungeon.Game.Rewards,
			}),
		}
	}

	//fmt.Println("[WS] Arena connected!")
	return &gosf.Message{
		Success: true,
		Body: gosf.StructToMap(&WebsocketUpdate{
			Grid:   smallGrid,
			Health: d.Dungeon.Game.Health,
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

	token, found := GetString(request, "token")
	if !found {
		return gosf.NewFailureMessage("f")
	}

	data, found := Tokens.Get(token)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}

	direction, found := GetInteger(request, "direction")
	if !found {
		return gosf.NewFailureMessage("no dir found")
	}

	d := data.(*DungeonUserData)

	end := d.Dungeon.Game.MovePlayer(direction)
	smallGrid := d.Dungeon.Game.GetSmallGrid(3, 3)
	if end {
		d.Dungeon.EndIt <- 1
		d.Dungeon.Ended = true

		if d.Dungeon.Game.Health <= 0 {
			d.Dungeon.Game.Rewards.Lien /= 2
		}

		return &gosf.Message{
			Success: true,
			Body: gosf.StructToMap(&WebsocketUpdate{
				Grid:    smallGrid,
				Health:  d.Dungeon.Game.Health,
				Ended:   true,
				Rewards: d.Dungeon.Game.Rewards,
				Win:     d.Dungeon.Game.Win,
			}),
		}
	}

	return &gosf.Message{
		Success: true,
		Body: gosf.StructToMap(&WebsocketUpdate{
			Grid:   smallGrid,
			Health: d.Dungeon.Game.Health,
		}),
	}
}

func AmbrosiusChoice(client *gosf.Client, request *gosf.Request) *gosf.Message {
	token, found := GetString(request, "token")
	if !found {
		return gosf.NewFailureMessage("f")
	}

	data, found := Tokens.Get(token)
	if !found {
		client.Disconnect()
		return gosf.NewFailureMessage("f")
	}
	d := data.(*DungeonUserData)
	if d.Dungeon.Ended {
		return gosf.NewFailureMessage("f")
	}

	choice, found := GetInteger(request, "choice")
	if !found {
		return gosf.NewFailureMessage("no choice found")
	}

	return d.Dungeon.Game.MakeChoice(choice)
}
