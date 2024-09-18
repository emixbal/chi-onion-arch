package request

import (
	"github.com/gookit/validate"
)

// Define the struct for the UpdateUserPassword request payload
type UpdateUserPassword struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

// Custom validation error messages
func (f UpdateUserPassword) Messages() map[string]string {
	return validate.MS{
		"OldPassword.required": "Password lama harus diisi.",
		"NewPassword.required": "Password baru harus diisi.",
	}
}

// Custom field translations
func (f UpdateUserPassword) Translates() map[string]string {
	return validate.MS{
		"OldPassword": "OldPassword",
		"NewPassword": "NewPassword",
	}
}
