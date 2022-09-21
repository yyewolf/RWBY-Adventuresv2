package main

import (
	"fmt"
	"net/http"
	"os"
	"rwby-adventures/auth/api"
	"rwby-adventures/config"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	APP_PORT := config.AuthPort

	router := gin.Default()

	CORS := os.Getenv("CORS")
	c := strings.Split(CORS, ",")
	c = append(c, "http://localhost:8080")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     c,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	basepath := router.Group("/api")

	api.LoadAPI(basepath)

	s := &http.Server{
		Addr:           APP_PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	fmt.Println("Error: ", err)
}
