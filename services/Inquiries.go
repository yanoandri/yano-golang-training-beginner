package services

import (
	"errors"

	"github.com/yanoandri/yano-golang-training-beginner/model"
	"gorm.io/gorm"
)

type InquiryCodes struct {
	ID            string `json:"id"`
	TransactionId string `json:"transaction_id"`
	PaymentCode   string `json:"payment_code"`
	Name          string `json:"name"`
	Status        string `json:"status"`
}

type IInquiriesService interface {
	CreateInquiry(inquiry model.Inquiries) (model.Inquiries, error)
	GetInquiryByTransactionId(id string) (InquiryCodes, error)
}

func (conn Repository) CreateInquiry(inquiry model.Inquiries) (model.Inquiries, error) {
	result := conn.Database.Create(&inquiry)
	if result.RowsAffected == 0 {
		return model.Inquiries{}, errors.New("inquiry data not created")
	}
	return inquiry, nil
}

func (conn Repository) GetInquiryByTransactionId(id string) (InquiryCodes, error) {
	var inquiry InquiryCodes
	result := conn.Database.Model(&model.Inquiries{}).Select("inquiries.id, inquiries.transaction_id, inquiries.payment_code, payment_codes.name, payment_codes.status").Joins("left join payment_codes on payment_codes.payment_code = inquiries.payment_code").Where("transaction_id = ?", id).Scan(&inquiry)
	if result.RowsAffected == 0 {
		return InquiryCodes{}, errors.New("inquiry data not found")
	}
	return inquiry, nil
}

func NewInquiryService(conn *gorm.DB) Repository {
	return Repository{Database: conn}
}
