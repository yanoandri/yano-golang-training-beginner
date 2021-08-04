package services

import (
	"errors"
	"time"

	"github.com/yanoandri/yano-golang-training-beginner/model"
	"gorm.io/gorm"
)

type IPaymentCodeService interface {
	CreatePaymentCode(payment model.PaymentCodes) (model.PaymentCodes, error)
	GetPaymentCodeById(id string) (model.PaymentCodes, error)
	ExpirePaymentCode() int64
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

func (conn Repository) ExpirePaymentCode() int64 {
	result := conn.Database.Model(model.PaymentCodes{}).Where("expiration_date < ? and status = ?", time.Now().Format(time.RFC3339), "ACTIVE").Updates(model.PaymentCodes{Status: "INACTIVE"})
	return result.RowsAffected
}

func NewPaymentCodeService(conn *gorm.DB) Repository {
	return Repository{Database: conn}
}
