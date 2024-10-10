package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type LogEntry struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PlayerID  uint      `json:"player_id" gorm:"not null" validate:"required"`
	Action    string    `json:"action" gorm:"not null" validate:"required,oneof=register login logout enter_room exit_room join_challenge challenge_result"`
	Details   string    `json:"details" gorm:"type:text"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (l *LogEntry) Validate() error {
	return validate.Struct(l)
}
