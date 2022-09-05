package main

import (
	"rwby-adventures/models"
)

func main() {
	// Gets all ids from database
	r, _ := database.DB.Query("SELECT id FROM players")
	i := 0
	for r.Next() {
		var id string
		r.Scan(&id)
		if id != "374018098673614860" {
			continue
		}
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
