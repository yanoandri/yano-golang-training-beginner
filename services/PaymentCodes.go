package services

import (
	"errors"

	"github.com/yanoandri/yano-golang-training-beginner/model"
	"gorm.io/gorm"
)

type IPaymentCodeService interface {
	CreatePaymentCode(payment model.PaymentCodes) (model.PaymentCodes, error)
	GetPaymentCodeById(id string) (model.PaymentCodes, error)
}

type Repository struct {
	Database *gorm.DB
}

func (conn Repository) CreatePaymentCode(payment model.PaymentCodes) (model.PaymentCodes, error) {
	result := conn.Database.Create(&payment)
	if result.RowsAffected == 0 {
		return model.PaymentCodes{}, errors.New("payment data not found")
	}
	return payment, nil
}

func (conn Repository) GetPaymentCodeById(id string) (model.PaymentCodes, error) {
	var payment model.PaymentCodes
	result := conn.Database.First(&payment, "id = ?", id)
	if result.RowsAffected == 0 {
		return model.PaymentCodes{}, errors.New("payment data not found")
	}
	return payment, nil
}

func NewPaymentCodeService(conn *gorm.DB) Repository {
	return Repository{Database: conn}
}
