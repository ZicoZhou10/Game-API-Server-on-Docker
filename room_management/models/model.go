package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Room struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:255;not null;unique" validate:"required,min=2,max=255"`
	Description string    `json:"description" gorm:"type:text"`
	Status      string    `json:"status" gorm:"size:20;not null" validate:"required,oneof=available occupied maintenance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Reservation struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	RoomID    uint      `json:"room_id" gorm:"not null" validate:"required"`
	Room      Room      `json:"room" gorm:"foreignKey:RoomID"`
	PlayerID  uint      `json:"player_id" gorm:"not null" validate:"required"`
	Date      time.Time `json:"date" gorm:"not null" validate:"required"`
	StartTime time.Time `json:"start_time" gorm:"not null" validate:"required"`
	EndTime   time.Time `json:"end_time" gorm:"not null" validate:"required,gtfield=StartTime"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (r *Room) Validate() error {
	return validate.Struct(r)
}

func (r *Reservation) Validate() error {
	return validate.Struct(r)
}
