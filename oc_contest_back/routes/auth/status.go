package auth

import (
	"rwby-adventures/models"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
)

func Status(c *gin.Context) {
	logged := c.MustGet("Logged").(bool)
	if !logged {
		c.JSON(200, gin.H{
			"logged": false,
		})
		return
	}

	u := c.MustGet("User").(*goth.User)
	voteAmount := models.GetAmountOfVote(u.UserID)

	c.JSON(200, gin.H{
		"logged": true,
		"user":   u,
		"votes":  voteAmount,
	})
}
