package submissions

import (
	"os"
	"rwby-adventures/config"
	"rwby-adventures/models"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	s := models.GetSubmission(id)
	if s == nil {
		c.JSON(404, gin.H{
			"error": "Submission not found",
		})
		return
	}

	u := c.MustGet("User").(*goth.User)

	if s.DiscordID != u.UserID {
		c.JSON(403, gin.H{
			"error": "You are not allowed to delete this submission",
		})
		return
	}

	// delete files first
	for _, f := range s.Files {
		os.Remove(f.Path)
	}

	config.Database.Delete(s)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
