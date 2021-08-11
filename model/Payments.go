package model

import "github.com/google/uuid"

type Payments struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;" json:"id"`
	TransactionId string    `gorm:"unique;" json:"transaction_id"`
	PaymentCode   string    `json:"payment_code"`
	Name          string    `json:"name"`
	Amount        string    `json:"amount"`
}
