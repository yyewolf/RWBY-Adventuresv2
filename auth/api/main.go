package api

import (
	"rwby-adventures/auth/api/login"
	"rwby-adventures/auth/api/logout"
	"rwby-adventures/auth/api/user"

	"github.com/gin-gonic/gin"
)

func LoadAPI(g *gin.RouterGroup) {
	user.LoadUserRoutes(g)
	login.LoadLoginRoutes(g)
	logout.LoadLogoutRoutes(g)
}
