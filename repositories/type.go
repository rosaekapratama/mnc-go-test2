package repositories

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/rosaekapratama/mnc-go-test2/models/repo"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, user *repo.User) (err error)
	FindById(ctx context.Context, id uuid.UUID) (user *repo.User, err error)
	FindByPhoneNo(ctx context.Context, phoneNo string) (user *repo.User, err error)
}

type userRepositoryImpl struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}
