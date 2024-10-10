package api

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/ZicoZhou10/interview_expert_20240109/challenge_system/models"
	"github.com/ZicoZhou10/interview_expert_20240109/challenge_system/repository"
	"github.com/gin-gonic/gin"
)

var lastChallengeTime = make(map[uint]time.Time)

func JoinChallenge(c *gin.Context) {
	var challenge models.Challenge
	if err := c.ShouldBindJSON(&challenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := challenge.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if player can join challenge (1 minute cooldown)
	lastTime, exists := lastChallengeTime[challenge.PlayerID]
	if exists && time.Since(lastTime) < time.Minute {
		c.JSON(http.StatusTooEarly, gin.H{"error": "You can only join a challenge once per minute"})
		return
	}

	// Process challenge
	challenge.CreatedAt = time.Now()
	challenge.Won = rand.Float32() < 0.01 // 1% chance of winning

	result := repository.DB.Create(&challenge)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to join challenge"})
		return
	}

	lastChallengeTime[challenge.PlayerID] = challenge.CreatedAt

	// Simulate 30 second challenge duration
	time.Sleep(30 * time.Second)

	if challenge.Won {
		c.JSON(http.StatusOK, gin.H{"message": "Congratulations! You won the challenge!", "won": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Sorry, you didn't win this time. Try again!", "won": false})
	}
}

func GetChallengeResults(c *gin.Context) {
	var challenges []models.Challenge
	result := repository.DB.Order("created_at desc").Limit(10).Find(&challenges)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch challenge results"})
		return
	}
	c.JSON(http.StatusOK, challenges)
}
