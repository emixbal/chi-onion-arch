package request

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

// Define the struct for the SubscriptionHistoryCreate request payload
type SubscriptionHistoryCreate struct {
	UserID          uuid.UUID `json:"user_id" validate:"required"`
	CategoryID      uuid.UUID `json:"category_id" validate:"required"`
	PaymentStatus   string    `json:"payment_status" validate:"required"`
	SubscribeStatus string    `json:"subscribe_status" validate:"required"`
	ExpiredTime     string    `json:"expired_time" validate:"required"`
	Price           float64   `json:"price" validate:"required"`
}

// Custom validation error messages
func (f SubscriptionHistoryCreate) Messages() map[string]string {
	return validate.MS{
		"UserID.required":          "User ID harus diisi.",
		"CategoryID.required":      "Category ID harus diisi.",
		"PaymentStatus.required":   "Status pembayaran harus diisi.",
		"SubscribeStatus.required": "Status langganan harus diisi.",
		"ExpiredTime.required":     "Waktu kadaluarsa harus diisi.",
		"ExpiredTime.date_format":  "Format waktu kadaluarsa harus DD-MM-YYYY.",
		"Price.required":           "Harga harus diisi.",
	}
}

// Custom field translations
func (f SubscriptionHistoryCreate) Translates() map[string]string {
	return validate.MS{
		"UserID":          "User ID",
		"CategoryID":      "Category ID",
		"PaymentStatus":   "Status Pembayaran",
		"SubscribeStatus": "Status Langganan",
		"ExpiredTime":     "Waktu Kadaluarsa",
		"Price":           "Harga",
	}
}

// Define the struct for the SubscriptionHistoryUpdate request payload
type SubscriptionHistoryUpdate struct {
	PaymentStatus   *string `json:"payment_status"`
	SubscribeStatus *string `json:"subscribe_status"`
}

// Custom validation error messages for Update
func (f SubscriptionHistoryUpdate) Messages() map[string]string {
	return validate.MS{
		"PaymentStatus.oneof":   "Status pembayaran harus salah satu dari [paid, unpaid].",
		"SubscribeStatus.oneof": "Status langganan harus salah satu dari [progress, cancel, success, expired].",
	}
}

// Custom field translations for Update
func (f SubscriptionHistoryUpdate) Translates() map[string]string {
	return validate.MS{
		"PaymentStatus":   "Status Pembayaran",
		"SubscribeStatus": "Status Langganan",
	}
}
