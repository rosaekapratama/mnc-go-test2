package rest

import (
	"github.com/google/uuid"
)

type BaseResponse[T any] struct {
	Status  string `json:"status,omitempty"` // Response status
	Message string `json:"message,omitempty"`
	Result  T      `json:"result,omitempty"`
}

type RegisterRequest struct {
	FirstName   string `json:"first_name"`   // First name
	LastName    string `json:"last_name"`    // Last name
	PhoneNumber string `json:"phone_number"` // Phone number
	Address     string `json:"address"`      // Address
	Pin         string `json:"pin"`          // PIN
}

type RegisterResponse struct {
	UserID      uuid.UUID `json:"user_id"`      // User ID (UUID)
	FirstName   string    `json:"first_name"`   // First name
	LastName    string    `json:"last_name"`    // Last name
	PhoneNumber string    `json:"phone_number"` // Phone number
	Address     string    `json:"address"`      // Address
	CreatedDate string    `json:"created_date"` // Created date
}
