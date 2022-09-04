package main

import "rwby-adventures/models"

func main() {
	// Gets all ids from database
	r, _ := database.DB.Query("SELECT id FROM players")
	for r.Next() {
		var id string
		r.Scan(&id)

		p := getPlayerInv(id)

		newP := models.GetPlayer(id)
		newP.Save()

		newP = playerConverter(p)
		newP.Save()
	}
}
