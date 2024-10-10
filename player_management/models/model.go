package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Player struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:255;not null" validate:"required,min=2,max=255"`
	Level     int       `json:"level" gorm:"not null" validate:"required,min=1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Level struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:255;not null"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (p *Player) Validate() error {
	return validate.Struct(p)
}
