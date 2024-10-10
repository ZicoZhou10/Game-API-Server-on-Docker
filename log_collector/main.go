package main

import (
	"github.com/ZicoZhou10/interview_expert_20240109/log_collector/api"
	"github.com/ZicoZhou10/interview_expert_20240109/log_collector/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDatabase()

	r := gin.Default()

	r.POST("/logs", api.AddLogEntry)
	r.GET("/logs", api.GetLogEntries)

	r.Run(":8080")
}
