package handles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubmitRequest struct {
	PlayerID string `json:"player_id"`
	Score    int    `json:"score"`
}

type SubmitResponse struct {
	Score int    `json:"score"`
	Rank  string `json:"rank"`
}

func SubmitScore(c *gin.Context) {
	var req SubmitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	rank := "B"
	if req.Score >= 100000 {
		rank = "S"
	} else if req.Score >= 50000 {
		rank = "A"
	}

	resp := SubmitResponse{
		Score: req.Score,
		Rank:  rank,
	}

	c.JSON(http.StatusOK, resp)
}
