package logout

import "github.com/gin-gonic/gin"

func LoadLogoutRoutes(g *gin.RouterGroup) {
	rg := g.Group("/logout")
	rg.GET("/", logout)
}
