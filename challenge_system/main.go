package main

import (
	"github.com/ZicoZhou10/interview_expert_20240109/challenge_system/api"
	"github.com/ZicoZhou10/interview_expert_20240109/challenge_system/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDatabase()

	r := gin.Default()

	r.POST("/challenges", api.JoinChallenge)
	r.GET("/challenges/results", api.GetChallengeResults)

	r.Run(":8080")
}
