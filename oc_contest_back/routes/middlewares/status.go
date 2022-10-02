package middlewares

import (
	"rwby-adventures/auth/export"

	"github.com/gin-gonic/gin"
)

func GetStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := export.GetUser(c)
		if err != nil {
			c.Set("Logged", false)
			return
		}

		c.Set("Logged", true)
		c.Set("User", u)
	}
}
