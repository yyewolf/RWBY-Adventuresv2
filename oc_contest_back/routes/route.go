package routes

import (
	"rwby-adventures/oc_contest_back/routes/auth"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(path *gin.RouterGroup) {
	auth.ServeAuth(path)
}
