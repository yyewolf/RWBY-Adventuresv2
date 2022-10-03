package auth

import (
	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	logged := c.MustGet("Logged").(bool)
	if !logged {
		c.JSON(200, gin.H{
			"logged": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"logged": true,
		"user":   c.MustGet("User"),
	})
}
