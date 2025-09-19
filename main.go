package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type timeStruct struct {
	Time time.Time `json:"time"`
}

func newTimeNow() timeStruct { return newTime(time.Now()) }

func newTime(time time.Time) timeStruct { return timeStruct{time} }

func getHealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, newTimeNow())
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/healthcheck", getHealthCheck)
	}
	return router
}

func main() {
	server := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	_ = server.Run(":" + port)
}
