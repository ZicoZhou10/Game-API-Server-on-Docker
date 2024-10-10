package services

import (
	"errors"
	"fmt"

	"github.com/ZicoZhou10/interview_expert_20240109/payment_system/models"
	"github.com/ZicoZhou10/interview_expert_20240109/payment_system/repository"
)

type PaymentService struct{}

func (s *PaymentService) ProcessPayment(payment *models.Payment) error {
	var err error

	switch payment.PaymentMethod {
	case "credit_card":
		err = s.processCreditCardPayment(payment)
	case "bank_transfer":
		err = s.processBankTransferPayment(payment)
	case "third_party":
		err = s.processThirdPartyPayment(payment)
	case "blockchain":
		err = s.processBlockchainPayment(payment)
	default:
		return errors.New("invalid payment method")
	}

	if err != nil {
		payment.Status = "failed"
		repository.DB.Save(payment)
		return err
	}

	payment.Status = "completed"
	repository.DB.Save(payment)
	return nil
}

func (s *PaymentService) processCreditCardPayment(payment *models.Payment) error {
	// Simulate credit card payment processing
	fmt.Println("Processing credit card payment...")
	payment.TransactionID = "CC-" + fmt.Sprintf("%d", payment.ID)
	return nil
}

func (s *PaymentService) processBankTransferPayment(payment *models.Payment) error {
	// Simulate bank transfer payment processing
	fmt.Println("Processing bank transfer payment...")
	payment.TransactionID = "BT-" + fmt.Sprintf("%d", payment.ID)
	return nil
}

func (s *PaymentService) processThirdPartyPayment(payment *models.Payment) error {
	// Simulate third-party payment processing
	fmt.Println("Processing third-party payment...")
	payment.TransactionID = "TP-" + fmt.Sprintf("%d", payment.ID)
	return nil
}

func (s *PaymentService) processBlockchainPayment(payment *models.Payment) error {
	// Simulate blockchain payment processing
	fmt.Println("Processing blockchain payment...")
	payment.TransactionID = "BC-" + fmt.Sprintf("%d", payment.ID)
	return nil
}
