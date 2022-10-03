package auth

import (
	"fmt"
	"rwby-adventures/auth/export"
	"rwby-adventures/config"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//export.Login(c, fmt.Sprintf("%slogin/callback", config.MarketHost))
	export.Login(c, fmt.Sprintf("%sauth/callback", config.OCHost))
}
