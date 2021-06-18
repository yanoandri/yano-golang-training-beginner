package services

import (
	"github.com/yanoandri/yano-golang-training-beginner/model"
	"github.com/yanoandri/yano-golang-training-beginner/repository"
)

type IPaymentCodeRepository interface {
	Create(payment model.PaymentCodes) (model.PaymentCodes, error)
	GetPaymentById(id string) (model.PaymentCodes, error)
}

type PaymentCodeRepository struct {
	PaymentCodeRepository repository.PaymentCodeRepository
}

func (repo PaymentCodeRepository) CreatePaymentCode(payment model.PaymentCodes) (model.PaymentCodes, error) {
	return repo.PaymentCodeRepository.Create(payment)
}

func (repo PaymentCodeRepository) Get(id string) (model.PaymentCodes, error) {
	return repo.PaymentCodeRepository.GetPaymentById(id)
}

func NewPaymentCodeService(repo repository.PaymentCodeRepository) PaymentCodeRepository {
	return PaymentCodeRepository{
		PaymentCodeRepository: repo,
	}
}
