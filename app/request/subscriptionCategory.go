package request

import (
	"github.com/gookit/validate"
)

// Define the struct for the InsertSubscriptionCategory request payload
type InsertSubscriptionCategory struct {
	Name         string  `json:"name" validate:"required"`
	PriceMonthly float64 `json:"price_monthly" validate:"required"`
	PriceAnnual  float64 `json:"price_annual" validate:"required"`
}

// Custom validation error messages for InsertSubscriptionCategory
func (f InsertSubscriptionCategory) Messages() map[string]string {
	return validate.MS{
		"Name.required":         "Nama kategori langganan harus diisi.",
		"PriceMonthly.required": "Harga bulanan harus diisi.",
		"PriceAnnual.required":  "Harga tahunan harus diisi.",
	}
}

// Custom field translations for InsertSubscriptionCategory
func (f InsertSubscriptionCategory) Translates() map[string]string {
	return validate.MS{
		"Name":         "Nama",
		"PriceMonthly": "Harga Bulanan",
		"PriceAnnual":  "Harga Tahunan",
	}
}

// Define the struct for the UpdateSubscriptionCategory request payload
type UpdateSubscriptionCategory struct {
	Name         *string  `json:"name"`
	PriceMonthly *float64 `json:"price_monthly"`
	PriceAnnual  *float64 `json:"price_annual"`
}

// Custom validation error messages for UpdateSubscriptionCategory
func (f UpdateSubscriptionCategory) Messages() map[string]string {
	return validate.MS{
		"Name.required":         "Nama kategori langganan harus diisi.",
		"PriceMonthly.required": "Harga bulanan harus diisi.",
		"PriceAnnual.required":  "Harga tahunan harus diisi.",
	}
}

// Custom field translations for UpdateSubscriptionCategory
func (f UpdateSubscriptionCategory) Translates() map[string]string {
	return validate.MS{
		"Name":         "Nama",
		"PriceMonthly": "Harga Bulanan",
		"PriceAnnual":  "Harga Tahunan",
	}
}
