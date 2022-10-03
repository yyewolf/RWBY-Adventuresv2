package submissions

import (
	"rwby-adventures/models"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
)

func Current(c *gin.Context) {
	u := c.MustGet("User").(*goth.User)
	s := models.GetUserSubmissions(u.UserID)

	c.JSON(200, gin.H{
		"submissions": s,
	})
}
