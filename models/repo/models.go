package repo

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID  `db:"id" json:"id"`                     // UUID
	FirstName   *string    `db:"first_name" json:"first_name"`     // Nullable
	LastName    *string    `db:"last_name" json:"last_name"`       // Nullable
	PhoneNumber string     `db:"phone_number" json:"phone_number"` // Nullable
	Address     *string    `db:"address" json:"address"`           // Nullable
	Pin         string     `db:"pin" json:"pin"`                   // Not nullable
	CreatedDt   time.Time  `db:"created_dt" json:"created_dt"`     // Not nullable
	UpdatedDt   *time.Time `db:"updated_dt" json:"updated_dt"`     // Nullable
}
