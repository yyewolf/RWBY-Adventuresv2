package gambles

import (
	"rwby-adventures/microservices"

	"github.com/yyewolf/gosf"
)

func UploadImage(req *microservices.GambleUpload) (*gosf.Message, error) {
	body := gosf.StructToMap(req)
	message := &gosf.Message{
		Body: body,
	}

	if response, err := GambleMicroservice.Call("addImage", message); err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
