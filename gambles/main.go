package main

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pmylund/go-cache"
	"github.com/yyewolf/gosf"
)

var images = cache.New(10*time.Minute, 10*time.Minute)

func CreateMicroservice() {
	// Start the server using a basic configuration
	go gosf.Startup(map[string]interface{}{"port": config.GambleRPC})
	fmt.Println("[GAMBLES] Microservice UP.")
}

func init() {
	// Listen on an endpoint
	gosf.Listen("addImage", addImage)
}

func main() {
	r := gin.Default()
	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		img, found := images.Get(id)
		if !found {
			c.String(404, "Image not found.")
			return
		}
		c.Data(200, "image/png", img.([]byte))
	})
	r.Run(config.GamblePort)
}
