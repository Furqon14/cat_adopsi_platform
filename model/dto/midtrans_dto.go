package dto

import (
	"time"

	"github.com/veritrans/go-midtrans"
)

type ChargeMidtrans struct {
	PaymentType        string `json:"payment_type"`
	TransactionDetails TransactionDetails
	CustomerDetails    CustomerDetails
	CustomField1       string `json:"custom_field1"`
	CustomField2       string `json:"custom_field2"`
	CustomField3       string `json:"custom_field3"`
	CustomExpiry       CustomExpiry
	Metadata           Metadata
}

type TransactionDetails struct {
	OrderID     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
}

type CustomerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type CustomExpiry struct {
	ExpiryDuration time.Duration `json:"expiry_duration"`
	Unit           string        `json:"unit"`
}

type Metadata struct {
	You       string `json:"you"`
	Put       string `json:"put"`
	Parameter string `json:"parameter"`
}

type Data struct {
	Tipe  string                `json:"tipe,omitempty"`
	Items []midtrans.ItemDetail `json:"items,omitempty"`
}

func (d *Data) GetTotal() int64 {
	var total int64
	for _, v := range d.Items {
		total += v.Price * int64(v.Qty)
	}
	return total
}
