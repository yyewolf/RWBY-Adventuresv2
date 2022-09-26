package login

import "github.com/gin-gonic/gin"

func LoadLoginRoutes(g *gin.RouterGroup) {
	rg := g.Group("/login")
	rg.GET("/", startLogin)
	rg.GET("/callback", callback)
	rg.POST("/code", codeGrant)
}
