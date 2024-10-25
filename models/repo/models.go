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

type Account struct {
	ID        uuid.UUID  `json:"id"`         // Corresponds to the 'id' field
	UserID    uuid.UUID  `json:"user_id"`    // Corresponds to the 'user_id' field
	Type      string     `json:"type"`       // Corresponds to the 'type' field
	Balance   float64    `json:"balance"`    // Corresponds to the 'balance' field (numeric(20,3) can be represented as float64)
	CreatedDt time.Time  `json:"created_dt"` // Corresponds to the 'created_dt' field (timestamptz)
	UpdateDt  *time.Time `json:"update_dt"`  // Corresponds to the 'update_dt' field (timestamptz) - can be NULL
}

type Transaction struct {
	ID        uuid.UUID `json:"id"`         // Corresponds to the 'id' field
	AccountID uuid.UUID `json:"account_id"` // Corresponds to the 'account_id' field (foreign key to accounts table)
	Type      string    `json:"type"`       // Corresponds to the 'type' field (transaction type)
	CR        bool      `json:"cr"`         // Corresponds to the 'cr' field (true for credit, false for debit)
	Amount    float64   `json:"amount"`     // Corresponds to the 'amount' field (numeric(20,3) is represented as float64)
	Balance   float64   `json:"balance"`    // Corresponds to the 'balance' field (numeric(20,3) is represented as float64)
	CreatedDt time.Time `json:"created_dt"` // Corresponds to the 'created_dt' field (timestamptz)
}
