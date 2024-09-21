package model

type Transaction struct {
	OrderID       string `json:"order_id"`
	GrossAmount   int64  `json:"gross_amount"`
	PaymentMethod string `json:"payment_method"`
}
