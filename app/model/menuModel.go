package model

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Path      string    `json:"path" db:"path" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy *string   `json:"updated_by" db:"updated_by"`
}
