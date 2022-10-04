package submissions

import (
	"rwby-adventures/oc_contest_back/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func ServeSubmissions(path *gin.RouterGroup) {
	subpath := path.Group("/submissions")

	subpath.GET("/current", middlewares.GetStatus(), middlewares.IsLogged(), Current)
	subpath.POST("/create", middlewares.GetStatus(), middlewares.IsLogged(), Create)
	subpath.GET("/delete/:id", middlewares.GetStatus(), middlewares.IsLogged(), Delete)
	subpath.GET("/vote/:id", middlewares.GetStatus(), middlewares.IsLogged(), Vote)
	subpath.GET("/all/:page", All)
}
