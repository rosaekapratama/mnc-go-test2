package repo

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID  `db:"id"`
	FirstName   *string    `db:"first_name"`
	LastName    *string    `db:"last_name"`
	PhoneNumber string     `db:"phone_number"`
	Address     *string    `db:"address"`
	Pin         string     `db:"pin"`
	CreatedDt   time.Time  `db:"created_dt"`
	UpdatedDt   *time.Time `db:"updated_dt"`
}

type Account struct {
	ID        uuid.UUID  `db:"id"`
	UserID    uuid.UUID  `db:"user_id"`
	Type      string     `db:"type"`
	Balance   float64    `db:"balance"`
	CreatedDt time.Time  `db:"created_dt"`
	UpdateDt  *time.Time `db:"update_dt"`
}

type Transaction struct {
	ID            uuid.UUID `db:"id"`
	AccountID     uuid.UUID `db:"account_id"`
	Type          string    `db:"type"`
	Status        string    `db:"status"`
	CR            bool      `db:"cr"`
	Amount        float64   `db:"amount"`
	BalanceBefore float64   `db:"balance_before"`
	BalanceAfter  float64   `db:"balance_after"`
	Remark        *string   `db:"remark"`
	CreatedDt     time.Time `db:"created_dt"`
}

type TransactionDetail struct {
	ID              string    `gorm:"column:id"`
	UserID          string    `gorm:"column:user_id"`
	TransactionType string    `gorm:"column:transaction_type"`
	Status          string    `gorm:"column:status"`
	Cr              bool      `gorm:"column:cr"`
	Amount          float64   `gorm:"column:amount"`
	Remarks         string    `gorm:"column:remarks"`
	BalanceBefore   float64   `gorm:"column:balance_before"`
	BalanceAfter    float64   `gorm:"column:balance_after"`
	CreatedDt       time.Time `gorm:"column:created_dt"`
}
