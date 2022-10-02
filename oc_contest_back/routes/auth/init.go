package auth

import (
	"rwby-adventures/oc_contest_back/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func ServeAuth(path *gin.RouterGroup) {
	subpath := path.Group("/auth")

	subpath.GET("/logout", Logout)
	subpath.GET("/login", Login)
	subpath.GET("/callback", Callback)
	subpath.GET("/status", middlewares.GetStatus(), Status)
}
