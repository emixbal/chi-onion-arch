package request

import (
	"github.com/gookit/validate"
)

// Define the struct for the FormInsertMenu request payload
type FormInsertMenu struct {
	Name string `json:"name" validate:"required"`
	Path string `json:"path" validate:"required"`
}

// Custom validation error messages
func (f FormInsertMenu) Messages() map[string]string {
	return validate.MS{
		"Name.required": "Nama menu harus diisi.",
		"Path.required": "Path menu harus diisi.",
	}
}

// Custom field translations
func (f FormInsertMenu) Translates() map[string]string {
	return validate.MS{
		"Name": "Nama Menu",
		"Path": "Path Menu",
	}
}
