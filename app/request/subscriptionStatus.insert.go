package request

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

type BulkInsertSubscriptionStatus struct {
	Subscriptions []InsertSubscriptionStatus `json:"subscriptions" validate:"required"`
}

type InsertSubscriptionStatus struct {
	CategoryBenefitID uuid.UUID `json:"category_benefit_id" validate:"required"`
	CategoryID        uuid.UUID `json:"category_id" validate:"required"`
	UserID            uuid.UUID `json:"user_id" validate:"required"`
	Qty               int       `json:"qty" validate:"required|min:1"`
	RemainingQty      int       `json:"remaining_qty" validate:"required|min:0"`
}

// Custom validation error messages for InsertSubscriptionStatus
func (f InsertSubscriptionStatus) Messages() map[string]string {
	return validate.MS{
		"CategoryBenefitID.required": "ID kategori manfaat harus diisi.",
		"CategoryID.required":        "ID kategori harus diisi.",
		"UserID.required":            "ID pengguna harus diisi.",
		"Qty.required":               "Jumlah harus diisi.",
		"Qty.min":                    "Jumlah harus minimal 1.",
		"RemainingQty.required":      "Jumlah tersisa harus diisi.",
		"RemainingQty.min":           "Jumlah tersisa tidak boleh kurang dari 0.",
	}
}

func (f BulkInsertSubscriptionStatus) Messages() map[string]string {
	return validate.MS{
		"Subscriptions.required": "List subscription harus diisi. Berbentuk array.",
	}
}
