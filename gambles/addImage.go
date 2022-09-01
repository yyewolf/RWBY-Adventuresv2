package main

import (
	"rwby-adventures/microservices"

	"github.com/pmylund/go-cache"
	"github.com/yyewolf/gosf"
)

func addImage(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req microservices.GambleUpload
	gosf.MapToStruct(request.Message.Body, &req)

	_, found := images.Get(req.UUID)

	if found {
		images.Set(req.UUID, req.Image, cache.DefaultExpiration)
	}

	return gosf.NewSuccessMessage()
}
