package websocket

import "rwby-adventures/microservices"

func EndDungeon(dungeon *DungeonStruct) *microservices.DungeonEndResponse {
	return &microservices.DungeonEndResponse{
		Loots: "",
	}
}
