package routes

import (
	"net/http"
	"os"
	"rwby-adventures/auth/export"
	"rwby-adventures/config"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()

	CORS := os.Getenv("CORS")
	c := strings.Split(CORS, ",")
	c = append(c, "http://localhost:8080")

	export.Provider()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     c,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	basepath := router.Group("/")

	initializeRoutes(basepath)

	s := &http.Server{
		Addr:           config.OCPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go s.ListenAndServe()
}
