package api

import (
	"net/http"
	"strconv"

	"github.com/ZicoZhou10/interview_expert_20240109/room_management/models"
	"github.com/ZicoZhou10/interview_expert_20240109/room_management/repository"
	"github.com/gin-gonic/gin"
)

func ListRooms(c *gin.Context) {
	var rooms []models.Room
	repository.DB.Find(&rooms)
	c.JSON(http.StatusOK, rooms)
}

func CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := room.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Create(&room)
	c.JSON(http.StatusCreated, room)
}

func GetRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var room models.Room
	if err := repository.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	c.JSON(http.StatusOK, room)
}

func UpdateRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var room models.Room
	if err := repository.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := room.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Save(&room)
	c.JSON(http.StatusOK, room)
}

func DeleteRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var room models.Room
	if err := repository.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	repository.DB.Delete(&room)
	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}
