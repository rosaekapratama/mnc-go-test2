package repositories

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/rosaekapratama/mnc-go-test2/models/repo"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user *repo.User) (err error)
	FindById(ctx context.Context, userId uuid.UUID) (user *repo.User, err error)
	FindByPhoneNo(ctx context.Context, phoneNo string) (user *repo.User, err error)
}

type AccountRepository interface {
	FindById(ctx context.Context, accountId uuid.UUID) (account *repo.Account, err error)
	FindSavingByUserId(ctx context.Context, userId uuid.UUID) (account *repo.Account, err error)
	CreateSaving(ctx context.Context, tx *gorm.DB, account *repo.Account) (err error)
	Topup(ctx context.Context, tx *gorm.DB, accountId uuid.UUID, amount float64) (account *repo.Account, err error)
}

type TransactionRepository interface {
	Topup(ctx context.Context, tx *gorm.DB, transaction *repo.Transaction) (err error)
}

type userRepositoryImpl struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

type accountRepositoryImpl struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

type transactionRepositoryImpl struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}
