package arenas

import (
	"rwby-adventures/microservices"

	"github.com/yyewolf/gosf"
)

func CreateArena(req *microservices.CreateArena) (*gosf.Message, error) {
	body := gosf.StructToMap(req)
	message := &gosf.Message{
		Body: body,
	}

	if response, err := ArenaMicroservice.Call("createArena", message); err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
