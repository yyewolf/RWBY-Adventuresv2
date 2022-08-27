package web

import (
	"encoding/json"
	"net/http"
	"rwby-adventures/config"
	"rwby-adventures/microservices"
	"rwby-adventures/models"
	"rwby-adventures/topgg/microservice"

	"github.com/bwmarrin/discordgo"
)

type topggreq struct {
	Bot       string `json:"bot"`
	User      string `json:"user"`
	Type      string `json:"type"`
	IsWeekend bool   `json:"isWeekend"`
	Query     string `json:"query?"`
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != config.TopGG {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var info topggreq
	err := decoder.Decode(&info)
	if err != nil {
		return
	}
	p := models.GetPlayer(info.User)
	if p.IsNew {
		return
	}

	microservice.SendMessageToBot(&microservices.TopGGMessage{
		UserID: info.User,
		Message: &discordgo.MessageEmbed{
			Title:       "Daily Reward :",
			Description: "Thank you for your vote !\nYou can claim your reward with : `/daily`.",
			Color:       config.Botcolor,
		},
	})

	p.Status.Voted = true
	p.Status.Save()
	w.WriteHeader(200)
}