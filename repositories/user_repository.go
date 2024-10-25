package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/rosaekapratama/go-starter/constant/str"
	"github.com/rosaekapratama/go-starter/database"
	"github.com/rosaekapratama/go-starter/log"
	repoModel "github.com/rosaekapratama/mnc-go-test2/models/repo"
	"gorm.io/gorm"
)

func NewUserRepository(ctx context.Context) UserRepository {
	gormDB, sqlDB, err := database.Manager.DB(ctx, "playground")
	if err != nil {
		log.Fatal(ctx, err, "error on database.Manager.DB(ctx, \"playground\")")
	}

	return &userRepositoryImpl{
		gormDB: gormDB,
		sqlDB:  sqlDB,
	}
}

func (r *userRepositoryImpl) BeginTx() *gorm.DB {
	return r.gormDB.Begin()
}

func (r *userRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user *repoModel.User) (err error) {
	err = tx.WithContext(ctx).Save(user).Error
	if err != nil {
		log.Error(ctx, err, "error on r.gormDB.WithContext(ctx).Save(user)")
		return
	}
	return
}

func (r *userRepositoryImpl) FindById(ctx context.Context, userId uuid.UUID) (user *repoModel.User, err error) {
	user = &repoModel.User{ID: userId}
	err = r.gormDB.WithContext(ctx).Find(user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		user = nil
		err = nil
	} else if err != nil {
		user = nil
		log.Errorf(ctx, err, "error on r.gormDB.WithContext(ctx).Find(user), userId=%s", userId.String())
	} else if user.PhoneNumber == str.Empty {
		user = nil
	}
	return
}

func (r *userRepositoryImpl) FindByPhoneNo(ctx context.Context, phoneNo string) (user *repoModel.User, err error) {
	user = &repoModel.User{}
	err = r.gormDB.WithContext(ctx).Find(user, "phone_number = ?", phoneNo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		user = nil
		err = nil
	} else if err != nil {
		user = nil
		log.Errorf(ctx, err, "error on r.gormDB.WithContext(ctx).Find(user, \"phone_no = ?\", phoneNo), phoneNo=%s", phoneNo)
	} else if user.PhoneNumber == str.Empty {
		user = nil
	}
	return
}
