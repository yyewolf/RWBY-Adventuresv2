package websocket

import "rwby-adventures/microservices"

func (dungeon *DungeonStruct) End() *microservices.DungeonEndResponse {
	return &microservices.DungeonEndResponse{
		Rewards: dungeon.Game.Rewards,
		Win:     dungeon.Game.Win,
	}
}
