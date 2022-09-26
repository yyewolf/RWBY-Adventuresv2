package login

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
)

type CodeGrant struct {
	Code     string `json:"code"`
	Password string `json:"password"`
}

func codeGrant(c *gin.Context) {
	var json CodeGrant
	c.Bind(&json)
	if json.Password != os.Getenv("AUTH_KEY") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	if json.Code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	d, found := codes.Get(json.Code)
	if !found {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	u := d.(goth.User)

	fmt.Println(u)
	c.JSON(200, gin.H{
		"code": json.Code,
		"user": u,
	})
}
