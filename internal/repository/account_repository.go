package repository

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountRepositoryInterface interface {
	Create(ctx context.Context, req model.CreateAccount) (*entity.Account, error)
	GetById(ctx context.Context, accountId uint) (*entity.Account, error)
	GetByUsername(ctx context.Context, username string) (*entity.Account, error)
	Update(ctx context.Context, req model.UpdateAccount) (*entity.Account, error)
	DeleteAccount(ctx context.Context, accountId uint) error
}

type AccountRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewAccountRepository(db *gorm.DB, log *logrus.Logger) AccountRepositoryInterface {
	return &AccountRepository{db: db, log: log}
}

func (a AccountRepository) Create(ctx context.Context, req model.CreateAccount) (*entity.Account, error) {
	account := &entity.Account{
		Avatar:    req.Avatar,
		Username:  req.Username,
		BirthDate: req.Birthdate,
	}

	err := a.db.WithContext(ctx).Create(account).Error
	if err != nil {
		a.log.Error("cannot insert data")
		return &entity.Account{}, err
	}

	return account, nil
}

func (a AccountRepository) GetById(ctx context.Context, accountId uint) (*entity.Account, error) {
	account := entity.Account{}
	err := a.db.WithContext(ctx).Where("ID = ?", accountId).First(&account).Error
	if err != nil {
		a.log.Error("account not found")
		return &entity.Account{}, err
	}

	return &account, nil
}

func (a AccountRepository) GetByUsername(ctx context.Context, username string) (*entity.Account, error) {
	account := &entity.Account{}
	err := a.db.WithContext(ctx).Where("username = ?", username).First(&account).Error
	if err != nil {
		a.log.Error("account not found")
		return &entity.Account{}, err
	}

	return account, nil
}

func (a AccountRepository) Update(ctx context.Context, req model.UpdateAccount) (*entity.Account, error) {
	account, err := a.GetById(ctx, req.ID)
	if err != nil {
		a.log.Error(err)
		return &entity.Account{}, err
	}

	arg := &entity.Account{
		Avatar: req.Avatar,
	}

	err = a.db.WithContext(ctx).Model(&account).Update("avatar", arg.Avatar).Error
	if err != nil {
		a.log.Error("cannot update data")
		return &entity.Account{}, err
	}

	return arg, nil
}

func (a AccountRepository) DeleteAccount(ctx context.Context, accountId uint) error {
	var account entity.Account

	err := a.db.WithContext(ctx).Delete(&account, accountId).Error
	if err != nil {
		a.log.Error("cannot delete account")
		return err
	}

	return nil
}
