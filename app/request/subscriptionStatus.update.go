package request

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

// Define the struct for the UpdateSubscriptionStatus request payload
type UpdateSubscriptionStatus struct {
	CategoryBenefitID uuid.UUID `json:"category_benefit_id" validate:"required"`
	UserID            uuid.UUID `json:"user_id" validate:"required"`
	Qty               int       `json:"qty" validate:"required|min:1"`
	RemainingQty      int       `json:"remaining_qty" validate:"required|min:0"`
}

// Custom validation error messages
func (f UpdateSubscriptionStatus) Messages() map[string]string {
	return validate.MS{
		"CategoryBenefitID.required": "ID kategori manfaat harus diisi.",
		"UserID.required":            "ID pengguna harus diisi.",
		"Qty.required":               "Jumlah harus diisi.",
		"Qty.min":                    "Jumlah harus minimal 1.",
		"RemainingQty.required":      "Jumlah tersisa harus diisi.",
		"RemainingQty.min":           "Jumlah tersisa tidak boleh kurang dari 0.",
	}
}

// Custom field translations
func (f UpdateSubscriptionStatus) Translates() map[string]string {
	return validate.MS{
		"CategoryBenefitID": "ID Kategori Manfaat",
		"UserID":            "ID Pengguna",
		"Qty":               "Jumlah",
		"RemainingQty":      "Jumlah Tersisa",
	}
}
