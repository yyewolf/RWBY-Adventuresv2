package main

import (
	"encoding/base64"
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/microservices"
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
	CreateMicroservice()

	r := gin.Default()
	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		img, found := images.Get(id)
		if !found {
			c.String(404, "Image not found.")
			return
		}
		b := img.(*microservices.GambleUpload)
		data, err := base64.StdEncoding.DecodeString(b.Image)
		if err != nil {
			c.String(500, "Error decoding image.")
			return
		}
		c.Data(200, "image/png", data)
	})
	r.Run(config.GamblePort)
}
