package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/rosaekapratama/go-starter/constant/location"
	"github.com/rosaekapratama/go-starter/database"
	"github.com/rosaekapratama/go-starter/log"
	repoModel "github.com/rosaekapratama/mnc-go-test2/models/repo"
	"gorm.io/gorm"
	"time"
)

func NewAccountRepository(ctx context.Context) AccountRepository {
	gormDB, sqlDB, err := database.Manager.DB(ctx, "playground")
	if err != nil {
		log.Fatal(ctx, err, "error on database.Manager.DB(ctx, \"playground\")")
	}

	return &accountRepositoryImpl{
		gormDB: gormDB,
		sqlDB:  sqlDB,
	}
}

func (r *accountRepositoryImpl) FindById(ctx context.Context, accountId uuid.UUID) (account *repoModel.Account, err error) {
	account = &repoModel.Account{ID: accountId}
	err = r.gormDB.WithContext(ctx).Find(account).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		account = nil
		err = nil
	} else if err != nil {
		account = nil
		log.Errorf(ctx, err, "error on r.gormDB.WithContext(ctx).Find(user), accountId=%s", accountId.String())
	} else if account.ID == uuid.Nil {
		account = nil
	}
	return
}

func (r *accountRepositoryImpl) FindSavingByUserId(ctx context.Context, userId uuid.UUID) (account *repoModel.Account, err error) {
	account = &repoModel.Account{}
	where := &repoModel.Account{UserID: userId, Type: "SAVING"}
	err = r.gormDB.WithContext(ctx).Where(where, "user_id").Find(account).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		account = nil
		err = nil
	} else if err != nil {
		account = nil
		log.Errorf(ctx, err, "error on r.gormDB.WithContext(ctx).FindSavingByUserId(user), userId=%s", userId.String())
	} else if account.ID == uuid.Nil {
		account = nil
	}
	return
}

func (r *accountRepositoryImpl) CreateSaving(ctx context.Context, tx *gorm.DB, account *repoModel.Account) (err error) {
	err = tx.WithContext(ctx).Save(account).Error
	if err != nil {
		log.Error(ctx, err, "error on r.gormDB.WithContext(ctx).Save(account)")
		return
	}
	return
}

func (r *accountRepositoryImpl) Credit(ctx context.Context, tx *gorm.DB, accountId uuid.UUID, amount float64) (account *repoModel.Account, err error) {
	account, err = r.FindById(ctx, accountId)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	if account == nil {
		return
	}

	now := time.Now().In(location.AsiaJakarta)
	account.Balance += amount
	account.UpdateDt = &now
	err = tx.WithContext(ctx).Save(account).Error
	if err != nil {
		log.Error(ctx, err, "error on tx.WithContext(ctx).Save(account)")
		return
	}
	return
}

func (r *accountRepositoryImpl) Debit(ctx context.Context, tx *gorm.DB, accountId uuid.UUID, amount float64) (account *repoModel.Account, err error) {
	account, err = r.FindById(ctx, accountId)
	if err != nil {
		log.Error(ctx, err)
		return
	}

	if account == nil {
		return
	}

	if account.Balance < amount {
		return account, errors.New("not enough balance")
	}

	now := time.Now().In(location.AsiaJakarta)
	account.Balance -= amount
	account.UpdateDt = &now
	err = tx.WithContext(ctx).Save(account).Error
	if err != nil {
		log.Error(ctx, err, "error on tx.WithContext(ctx).Save(account)")
		return
	}
	return
}
