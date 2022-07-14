package websocket

import (
	"fmt"
	dungeonpc "rwby-adventures/dungeons_rpc"
	"time"
)

func CreateDungeon(in *dungeonpc.CreateDungeonReq) (b bool, loots string) {
	_, exists := DungeonCache.Get(in.GetId())
	if exists {
		return true, ""
	}
	dungeon := &DungeonStruct{
		ID:  in.GetId(),
		End: endDungeon,
	}
	DungeonCache.Set(in.GetId(), dungeon, 0)
	fmt.Println("[DUNGEON] Created dungeon:", in.GetId())
	return DungeonLoop(dungeon)
}

func DungeonLoop(dungeon *DungeonStruct) (b bool, loots string) {
	//Sends data to players
	t := time.NewTicker(time.Millisecond * 100)
	dungeon.Ticker = t
	for {
		select {
		case <-dungeon.Channel:
			return dungeon.End(dungeon)
		case <-dungeon.Ticker.C:
		}
		// fmt.Println("[ARENA] Sending data to players")
		// No operations necessary if no one is in
	}
}
