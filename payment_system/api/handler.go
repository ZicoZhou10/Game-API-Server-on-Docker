package api

import (
	"net/http"
	"strconv"

	"github.com/ZicoZhou10/interview_expert_20240109/payment_system/models"
	"github.com/ZicoZhou10/interview_expert_20240109/payment_system/repository"
	"github.com/ZicoZhou10/interview_expert_20240109/payment_system/services"
	"github.com/gin-gonic/gin"
)

var paymentService = &services.PaymentService{}

func ProcessPayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := payment.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.Status = "pending"
	result := repository.DB.Create(&payment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment record"})
		return
	}

	err := paymentService.ProcessPayment(&payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment", "payment_id": payment.ID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully", "payment_id": payment.ID, "status": payment.Status, "transaction_id": payment.TransactionID})
}

func GetPayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var payment models.Payment
	result := repository.DB.First(&payment, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}
