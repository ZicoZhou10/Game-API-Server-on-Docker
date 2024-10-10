package main

import (
	"github.com/ZicoZhou10/interview_expert_20240109/payment_system/api"
	"github.com/ZicoZhou10/interview_expert_20240109/payment_system/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDatabase()

	r := gin.Default()

	r.POST("/payments", api.ProcessPayment)
	r.GET("/payments/:id", api.GetPayment)

	r.Run(":8080")
}
