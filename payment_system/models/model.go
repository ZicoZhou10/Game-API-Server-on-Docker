package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Payment struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	PlayerID       uint      `json:"player_id" gorm:"not null" validate:"required"`
	Amount         float64   `json:"amount" gorm:"not null" validate:"required,gt=0"`
	PaymentMethod  string    `json:"payment_method" gorm:"not null" validate:"required,oneof=credit_card bank_transfer third_party blockchain"`
	Status         string    `json:"status" gorm:"not null" validate:"required,oneof=pending completed failed"`
	TransactionID  string    `json:"transaction_id"`
	PaymentDetails string    `json:"payment_details" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (p *Payment) Validate() error {
	return validate.Struct(p)
}
