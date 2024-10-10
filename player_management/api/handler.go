package api

import (
	"net/http"
	"strconv"

	"github.com/ZicoZhou10/interview_expert_20240109/player_management/models"
	"github.com/ZicoZhou10/interview_expert_20240109/player_management/repository"
	"github.com/gin-gonic/gin"
)

func ListPlayers(c *gin.Context) {
	var players []models.Player
	repository.DB.Find(&players)
	c.JSON(http.StatusOK, players)
}

func CreatePlayer(c *gin.Context) {
	var player models.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Create(&player)
	c.JSON(http.StatusCreated, player)
}

func GetPlayer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var player models.Player
	if err := repository.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func UpdatePlayer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var player models.Player
	if err := repository.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Save(&player)
	c.JSON(http.StatusOK, player)
}

func DeletePlayer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var player models.Player
	if err := repository.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	repository.DB.Delete(&player)
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}
