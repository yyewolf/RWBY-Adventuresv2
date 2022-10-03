package routes

import (
	"rwby-adventures/oc_contest_back/routes/auth"
	"rwby-adventures/oc_contest_back/routes/submissions"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(path *gin.RouterGroup) {
	auth.ServeAuth(path)
	submissions.ServeSubmissions(path)
}
