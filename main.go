package main

import (
	"github.com/mikio815/scoregame-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/submit", handlers.SubmitScore)
	r.Run()
}
