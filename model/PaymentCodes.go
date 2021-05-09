package model

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Active   Status = "ACTIVE"
	Inactive Status = "INACTIVE"
	Expired  Status = "EXPIRED"
)

type PaymentCodes struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	PaymentCode    string
	Name           string
	Status         Status
	ExpirationDate string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
