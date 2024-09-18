package request

import (
	"github.com/gookit/validate"
)

// Define the struct for the FormUpdateMenu request payload
type FormUpdateMenu struct {
	Name string `json:"name" validate:"required"`
	Path string `json:"path" validate:"required"`
}

// Custom validation error messages
func (f FormUpdateMenu) Messages() map[string]string {
	return validate.MS{
		"Name.required": "Nama menu harus diisi.",
		"Path.required": "Path menu harus diisi.",
	}
}

// Custom field translations
func (f FormUpdateMenu) Translates() map[string]string {
	return validate.MS{
		"Name": "Nama Menu",
		"Path": "Path Menu",
	}
}
