package main

import (
	"github.com/ZicoZhou10/interview_expert_20240109/room_management/api"
	"github.com/ZicoZhou10/interview_expert_20240109/room_management/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDatabase()

	r := gin.Default()

	r.GET("/rooms", api.ListRooms)
	r.POST("/rooms", api.CreateRoom)
	r.GET("/rooms/:id", api.GetRoom)
	r.PUT("/rooms/:id", api.UpdateRoom)
	r.DELETE("/rooms/:id", api.DeleteRoom)

	r.Run(":8080")
}
