package user

import "github.com/gin-gonic/gin"

func LoadUserRoutes(g *gin.RouterGroup) {
	rg := g.Group("/user")
	rg.GET("/", GetUser)
}
