package request

import (
	"github.com/gookit/validate"
)

// Define the struct for the RoleInsert request payload
type RoleInsert struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}

// Custom validation error messages
func (f RoleInsert) Messages() map[string]string {
	return validate.MS{
		"Name.required": "Nama harus diisi.",
		"Code.required": "Code harus diisi.",
	}
}

// Custom field translations
func (f RoleInsert) Translates() map[string]string {
	return validate.MS{
		"Name": "Nama Role",
		"Code": "Code Role",
	}
}

// Define the struct for the RoleUpdate request payload
type RoleUpdate struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}

// Custom validation error messages
func (f RoleUpdate) Messages() map[string]string {
	return validate.MS{
		"Name.required": "Nama harus diisi.",
		"Code.required": "Code harus diisi.",
	}
}

// Custom field translations
func (f RoleUpdate) Translates() map[string]string {
	return validate.MS{
		"Name": "Nama Role",
		"Code": "Code Role",
	}
}
