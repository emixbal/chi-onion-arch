package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/gookit/validate"
)

// Define the struct for the UpdateUserProfile request payload
type UpdateUserProfile struct {
	FirstName      string    `json:"first_name" validate:"required"`
	LastName       string    `json:"last_name" validate:"required"`
	Email          string    `json:"email" validate:"required|email"`
	Phone          string    `json:"phone" validate:"required"`
	RoleCode       string    `json:"role_code" validate:"required"`
	IsActive       bool      `json:"is_active"`
	IsVerification bool      `json:"is_verification"`
	UpdatedBy      uuid.UUID `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Custom validation error messages
func (f UpdateUserProfile) Messages() map[string]string {
	return validate.MS{
		"FirstName.required": "Nama depan harus diisi.",
		"LastName.required":  "Nama belakang harus diisi.",
		"Email.required":     "Alamat email harus diisi.",
		"Email.email":        "Alamat email tidak valid.",
		"Phone.required":     "Nomor telepon harus diisi.",
		"RoleCode.required":  "Kode peran harus diisi.",
	}
}

// Custom field translations
func (f UpdateUserProfile) Translates() map[string]string {
	return validate.MS{
		"FirstName":      "Nama Depan",
		"LastName":       "Nama Belakang",
		"Email":          "Alamat Email",
		"Phone":          "Nomor Telepon",
		"RoleCode":       "Kode Peran",
		"IsActive":       "Status Aktif",
		"IsVerification": "Status Verifikasi",
		"UpdatedBy":      "Diperbarui Oleh",
		"UpdatedAt":      "Waktu Pembaruan",
	}
}
