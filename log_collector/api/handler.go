package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ZicoZhou10/interview_expert_20240109/log_collector/models"
	"github.com/ZicoZhou10/interview_expert_20240109/log_collector/repository"
	"github.com/gin-gonic/gin"
)

func AddLogEntry(c *gin.Context) {
	var logEntry models.LogEntry
	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logEntry.Timestamp = time.Now()

	if err := logEntry.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := repository.DB.Create(&logEntry)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create log entry"})
		return
	}

	c.JSON(http.StatusCreated, logEntry)
}

func GetLogEntries(c *gin.Context) {
	playerID, _ := strconv.Atoi(c.Query("player_id"))
	action := c.Query("action")
	startTime, _ := time.Parse(time.RFC3339, c.Query("start_time"))
	endTime, _ := time.Parse(time.RFC3339, c.Query("end_time"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	query := repository.DB.Model(&models.LogEntry{})

	if playerID != 0 {
		query = query.Where("player_id = ?", playerID)
	}
	if action != "" {
		query = query.Where("action = ?", action)
	}
	if !startTime.IsZero() {
		query = query.Where("timestamp >= ?", startTime)
	}
	if !endTime.IsZero() {
		query = query.Where("timestamp <= ?", endTime)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}

	var logEntries []models.LogEntry
	result := query.Order("timestamp desc").Find(&logEntries)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch log entries"})
		return
	}

	c.JSON(http.StatusOK, logEntries)
}
