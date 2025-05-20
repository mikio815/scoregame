package main

import (
	"scoregame-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/submit", handlers.SubmitScore)
	r.Run()
}
