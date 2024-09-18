package request

import (
	"github.com/gookit/validate"
)

// Define the struct for the InsertPermission request payload
type InsertPermission struct {
	UserID string `json:"user_id" validate:"required"`
	MenuID string `json:"menu_id" validate:"required"`
}

// Custom validation error messages for InsertPermission
func (f InsertPermission) Messages() map[string]string {
	return validate.MS{
		"UserID.required": "User ID permission harus diisi.",
		"MenuID.required": "Menu ID permission harus diisi.",
	}
}

// Custom field translations for InsertPermission
func (f InsertPermission) Translates() map[string]string {
	return validate.MS{
		"UserID": "User ID",
		"MenuID": "Menu ID",
	}
}

// Define the struct for the UpdatePermission request payload
type UpdatePermission struct {
	UserID string `json:"user_id" validate:"required"`
	MenuID string `json:"menu_id" validate:"required"`
}

// Custom validation error messages for UpdatePermission
func (f UpdatePermission) Messages() map[string]string {
	return validate.MS{
		"UserID.required": "User ID permission harus diisi.",
		"MenuID.required": "Menu ID permission harus diisi.",
	}
}

// Custom field translations for UpdatePermission
func (f UpdatePermission) Translates() map[string]string {
	return validate.MS{
		"UserID": "User ID",
		"MenuID": "Menu ID",
	}
}
