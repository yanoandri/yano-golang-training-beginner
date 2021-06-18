package repository

import (
	"errors"

	"github.com/yanoandri/yano-golang-training-beginner/model"
	"gorm.io/gorm"
)

type PaymentCodeRepository struct {
	Connection *gorm.DB
}

func (db PaymentCodeRepository) Create(payment model.PaymentCodes) (model.PaymentCodes, error) {
	result := db.Connection.Create(&payment)
	if result.RowsAffected == 0 {
		return model.PaymentCodes{}, errors.New("Payment data not found")
	}
	return payment, nil
}

func (db PaymentCodeRepository) GetPaymentById(id string) (model.PaymentCodes, error) {
	var payment model.PaymentCodes
	result := db.Connection.First(&payment, "id = ?", id)
	if result.RowsAffected == 0 {
		return model.PaymentCodes{}, errors.New("Payment data not found")
	}
	return payment, nil
}

func NewPaymentCodeRepository(connection *gorm.DB) *PaymentCodeRepository {
	return &PaymentCodeRepository{Connection: connection}
}
