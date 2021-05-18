package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	Active   = "ACTIVE"
	Inactive = "INACTIVE"
	Expired  = "EXPIRED"
)

type PaymentCodes struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;" json:"id"`
	PaymentCode    string    `json:"payment_code"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	ExpirationDate string    `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
