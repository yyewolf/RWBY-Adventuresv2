package web

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("[TOPGG] Failed authorization from top.gg")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var info topggreq
	err := decoder.Decode(&info)
	if err != nil {
		fmt.Println("[TOPGG] Failed to decode info from top.gg :", err)
		return
	}
	p := models.GetPlayer(info.User)
	if p.IsNew {
		fmt.Println("[TOPGG] Player is new and can't receive daily.")
		return
	}

	fmt.Println("[TOPGG] Received vote from : " + info.User)

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
