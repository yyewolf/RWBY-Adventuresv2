package main

import (
	"fmt"
	"rwby-adventures/models"
)

func main() {
	// Gets all ids from database
	r, _ := database.DB.Query("SELECT id FROM players")
	i := 0
	for r.Next() {
		var id string
		r.Scan(&id)
		fmt.Println(i)
		p := getPlayerInv(id)
		if p.IsNew {
			continue
		}

		newP := models.GetPlayer(id)
		newP.Save()

		newP = playerConverter(p)
		newP.Save()
		i++
	}
}
