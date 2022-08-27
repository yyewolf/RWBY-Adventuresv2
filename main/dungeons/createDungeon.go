package dungeons

import (
	"rwby-adventures/microservices"

	"github.com/yyewolf/gosf"
)

func CreateDungeon(req *microservices.DungeonCreateRequest) (*gosf.Message, error) {
	body := gosf.StructToMap(req)
	message := &gosf.Message{
		Body: body,
	}

	if response, err := DungeonsMicroservice.Call("createDungeon", message); err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
