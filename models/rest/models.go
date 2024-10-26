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

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"` // Phone number
	Pin         string `json:"pin"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TopupRequest struct {
	Amount float64 `json:"amount"` // Representing the 'amount' field
}

type TopUpResponse struct {
	TopupID       string  `json:"top_up_id"`      // Representing the 'top_up_id' field as a string (UUID as a string)
	AmountTopup   float64 `json:"amount_top_up"`  // Representing the 'amount_top_up' field as an integer
	BalanceBefore float64 `json:"balance_before"` // Representing the 'balance_before' field
	BalanceAfter  float64 `json:"balance_after"`  // Representing the 'balance_after' field
	CreatedDate   string  `json:"created_date"`   // Representing the 'created_date' field as a timestamp
}

// PaymentRequest represents the incoming payment request
type PaymentRequest struct {
	Amount  float64 `json:"amount"`
	Remarks string  `json:"remarks"`
}

// PaymentResponse represents the details of the payment after processing
type PaymentResponse struct {
	PaymentID     string  `json:"payment_id"`
	Amount        float64 `json:"amount"`
	Remarks       string  `json:"remarks"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedDate   string  `json:"created_date"`
}

type TransferRequest struct {
	TargetUser string  `json:"target_user"`
	Amount     float64 `json:"amount"`
	Remarks    string  `json:"remarks"`
}

type TransferResponse struct {
	TransferID    string  `json:"transfer_id"`
	Amount        float64 `json:"amount"`
	Remarks       string  `json:"remarks"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedDate   string  `json:"created_date"`
}

type TransactionDetailResponse struct {
	ID              string  `json:"id"`
	UserID          string  `json:"user_id"`
	TransactionType string  `json:"transaction_type"`
	Status          string  `json:"status"`
	Cr              bool    `json:"cr"`
	Amount          float64 `json:"amount"`
	Remarks         string  `json:"remarks"`
	BalanceBefore   float64 `json:"balance_before"`
	BalanceAfter    float64 `json:"balance_after"`
	CreatedDate     string  `json:"created_date"`
}
