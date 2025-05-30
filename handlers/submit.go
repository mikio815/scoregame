package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubmitRequest struct {
	PlayerID string  `json:"player_id"`
	Score    int     `json:"score"`
	Combo    int     `json:"combo"`
	PlayTime float64 `json:"play_time"`
}

type SubmitResponse struct {
	Score      int      `json:"score"`
	Rank       string   `json:"rank"`
	Suspicious bool     `json:"suspicious"`
	Warning    []string `json:"warnings"`
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
		Score:      req.Score,
		Rank:       rank,
		Suspicious: false,
		Warning:    []string{"Score Correct"},
	}

	if req.PlayTime < 2.0 && req.Score > 100000 {
		resp.Suspicious = true
		resp.Warning = []string{"Play time too short for the score!"}
	}

	c.JSON(http.StatusOK, resp)
}
