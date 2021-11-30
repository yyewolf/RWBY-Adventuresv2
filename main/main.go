package main

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/models"
)

func main() {
	fmt.Println("Hello")
	p := models.Player{
		DiscordID: "5",
	}
	config.Database.Create(&p)

	u := &models.Player{
		Status: models.PlayerStatus{},
	}
	config.Database.Preload("Profile").Take(u)
	fmt.Println("User bal :")
	fmt.Printf("%v", u.Status.Voted)
}
