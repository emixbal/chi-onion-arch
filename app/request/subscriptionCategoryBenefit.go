package request

import (
	"github.com/gookit/validate"
)

// Define the struct for the InsertCategoryBenefit request payload
type InsertCategoryBenefit struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
	Qty  int    `json:"qty" validate:"required|min:1"`
}

// Custom validation error messages
func (f InsertCategoryBenefit) Messages() map[string]string {
	return validate.MS{
		"Code.required": "Kode benefit harus diisi.",
		"Name.required": "Nama benefit harus diisi.",
		"Qty.required":  "Jumlah (qty) harus diisi.",
		"Qty.min":       "Jumlah (qty) minimal adalah 1.",
	}
}

// Custom field translations
func (f InsertCategoryBenefit) Translates() map[string]string {
	return validate.MS{
		"Code": "Kode Benefit",
		"Name": "Nama Benefit",
		"Qty":  "Jumlah",
	}
}

// Define the struct for the UpdateCategoryBenefit request payload
type UpdateCategoryBenefit struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
	Qty  int    `json:"qty" validate:"required|min:1"`
}

// Custom validation error messages
func (f UpdateCategoryBenefit) Messages() map[string]string {
	return validate.MS{
		"Code.required": "Kode benefit harus diisi.",
		"Name.required": "Nama benefit harus diisi.",
		"Qty.required":  "Jumlah (qty) harus diisi.",
		"Qty.min":       "Jumlah (qty) minimal adalah 1.",
	}
}

// Custom field translations
func (f UpdateCategoryBenefit) Translates() map[string]string {
	return validate.MS{
		"Code": "Kode Benefit",
		"Name": "Nama Benefit",
		"Qty":  "Jumlah",
	}
}
