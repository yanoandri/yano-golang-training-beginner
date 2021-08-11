package services

import (
	"errors"

	"github.com/yanoandri/yano-golang-training-beginner/model"
	"gorm.io/gorm"
)

type IInquiriesService interface {
	CreateInquiry(inquiry model.Inquiries) (model.Inquiries, error)
	GetInquiryByTransactionId(id string) (model.Inquiries, error)
}

func (conn Repository) CreateInquiry(inquiry model.Inquiries) (model.Inquiries, error) {
	result := conn.Database.Create(&inquiry)
	if result.RowsAffected == 0 {
		return model.Inquiries{}, errors.New("inquiry data not created")
	}
	return inquiry, nil
}

func (conn Repository) GetInquiryByTransactionId(id string) (model.Inquiries, error) {
	var inquiry model.Inquiries
	result := conn.Database.First(&inquiry, "transaction_id = ?", id)
	if result.RowsAffected == 0 {
		return model.Inquiries{}, errors.New("inquiry data not found")
	}
	return inquiry, nil
}

func NewInquiryService(conn *gorm.DB) Repository {
	return Repository{Database: conn}
}
