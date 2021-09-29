package services

import (
	"errors"

	"github.com/yanoandri/yano-golang-training-beginner/model"
	"gorm.io/gorm"
)

type IPaymentsService interface {
	CreatePayment(payment model.Payments) (model.Payments, error)
	GetPaymentById(id string) (model.Payments, error)
}

func (conn Repository) CreatePayment(payment model.Payments) (model.Payments, error) {
	result := conn.Database.Create(&payment)
	if result.RowsAffected == 0 {
		return model.Payments{}, errors.New("payment data not created")
	}
	return payment, nil
}

func (conn Repository) GetPaymentById(id string) (model.Payments, error) {
	var payment model.Payments
	result := conn.Database.First(&payment, "id = ?", id)
	if result.RowsAffected == 0 {
		return model.Payments{}, errors.New("payment data not found")
	}
	return payment, nil
}

func NewPaymentService(conn *gorm.DB) Repository {
	return Repository{Database: conn}
}
