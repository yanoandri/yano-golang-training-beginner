package request

type (
	PaymentCodes struct {
		PaymentCode string `json:"payment_code" validate:"required"`
		Name        string `json:"name" validate:"required"`
	}
)
