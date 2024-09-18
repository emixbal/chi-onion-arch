package request

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

// Define the struct for the InsertMenu request payload
type InsertMenu struct {
	Name         string     `json:"name" validate:"required"`
	Path         string     `json:"path" validate:"required"`
	Platform     string     `json:"platform" validate:"required|in:website,portal-admin,portal-member,application"`
	IsChild      bool       `json:"is_child"`
	ParentID     *uuid.UUID `json:"parent_id"`
	Title        string     `json:"title"`
	MetaOrder    int        `json:"meta_order"`
	RequiresAuth bool       `json:"requires_auth" validate:"required"`
}

// Custom validation error messages
func (f InsertMenu) Messages() map[string]string {
	return validate.MS{
		"Name.required":         "Nama menu harus diisi.",
		"Path.required":         "Path menu harus diisi.",
		"Platform.required":     "Platform harus diisi.",
		"Platform.in":           "Platform harus salah satu dari: website, portal-admin, portal-member, application.",
		"RequiresAuth.required": "Status requires_auth harus diisi.",
	}
}

// Custom field translations
func (f InsertMenu) Translates() map[string]string {
	return validate.MS{
		"Name":         "Nama Menu",
		"Path":         "Path Menu",
		"Platform":     "Platform",
		"IsChild":      "Status Child Menu",
		"ParentID":     "Parent ID",
		"Title":        "Judul Menu",
		"MetaOrder":    "Urutan Menu",
		"RequiresAuth": "Memerlukan Autentikasi",
	}
}
