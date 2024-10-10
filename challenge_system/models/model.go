package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Challenge struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PlayerID  uint      `json:"player_id" gorm:"not null" validate:"required"`
	Amount    float64   `json:"amount" gorm:"not null" validate:"required,eq=20.01"`
	Won       bool      `json:"won" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (c *Challenge) Validate() error {
	return validate.Struct(c)
}
