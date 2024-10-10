package main

import (
	"github.com/ZicoZhou10/interview_expert_20240109/player_management/api"
	"github.com/ZicoZhou10/interview_expert_20240109/player_management/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDatabase()

	r := gin.Default()

	r.GET("/players", api.ListPlayers)
	r.POST("/players", api.CreatePlayer)
	r.GET("/players/:id", api.GetPlayer)
	r.PUT("/players/:id", api.UpdatePlayer)
	r.DELETE("/players/:id", api.DeletePlayer)

	r.Run(":8080")
}
