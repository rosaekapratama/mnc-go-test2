package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/rosaekapratama/go-starter/database"
	"github.com/rosaekapratama/go-starter/log"
	repoModel "github.com/rosaekapratama/mnc-go-test2/models/repo"
	"gorm.io/gorm"
)

func NewTransactionRepository(ctx context.Context) TransactionRepository {
	gormDB, sqlDB, err := database.Manager.DB(ctx, "playground")
	if err != nil {
		log.Fatal(ctx, err, "error on database.Manager.DB(ctx, \"playground\")")
	}

	return &transactionRepositoryImpl{
		gormDB: gormDB,
		sqlDB:  sqlDB,
	}
}

func (r *transactionRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, transaction *repoModel.Transaction) (err error) {
	err = tx.WithContext(ctx).Save(transaction).Error
	if err != nil {
		log.Error(ctx, err, "error on r.gormDB.WithContext(ctx).Save(transaction)")
		return
	}
	return
}

func (r *transactionRepositoryImpl) FindAllByUserId(ctx context.Context, userId uuid.UUID) (transactionList []*repoModel.TransactionDetail, err error) {
	err = r.gormDB.WithContext(ctx).Raw(queryFindAllTransactionByUserId, userId).Scan(&transactionList).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		return
	}
	if err != nil {
		log.Error(ctx, err, "error on r.gormDB.WithContext(ctx).Save(transaction)")
		return
	}
	return
}
