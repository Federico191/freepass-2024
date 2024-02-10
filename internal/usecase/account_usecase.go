package usecase

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type AccountUseCaseInterface interface {
	Create(ctx context.Context, req model.CreateAccount) (*entity.Account, error)
	GetById(ctx context.Context, accountId uint) (*entity.Account, error)
	GetByUsername(ctx context.Context, username string) (*entity.Account, error)
	Update(ctx context.Context, req model.UpdateAccount) (*entity.Account, error)
	DeleteAccount(ctx context.Context, accountId uint) error
}

type AccountUseCase struct {
	AccountRepo repository.AccountRepositoryInterface
	TimeOut     time.Duration
}

func NewAccountUseCase(AccountRepo repository.AccountRepositoryInterface, timeOut time.Duration) AccountUseCaseInterface {
	return &AccountUseCase{AccountRepo: AccountRepo, TimeOut: timeOut}
}

func (a AccountUseCase) Create(ctx context.Context, req model.CreateAccount) (*entity.Account, error) {
	account, err := a.AccountRepo.Create(ctx, req)
	if err != nil {
		return &entity.Account{}, err
	}

	return account, nil
}

func (a AccountUseCase) GetById(ctx context.Context, accountId uint) (*entity.Account, error) {
	account, err := a.AccountRepo.GetById(ctx, accountId)
	if err != nil {
		return &entity.Account{}, err
	}

	return account, nil
}
func (a AccountUseCase) GetByUsername(ctx context.Context, username string) (*entity.Account, error) {
	account, err := a.AccountRepo.GetByUsername(ctx, username)
	if err != nil {
		return &entity.Account{}, nil
	}

	return account, nil
}

func (a AccountUseCase) Update(ctx context.Context, req model.UpdateAccount) (*entity.Account, error) {
	_, err := a.AccountRepo.GetById(ctx, req.ID)
	if err != nil {
		return &entity.Account{}, err
	}

	req.UpdatedAt = time.Now()
	updatedAccount, err := a.AccountRepo.Update(ctx, req)
	if err != nil {
		return &entity.Account{}, err
	}

	return updatedAccount, nil
}

func (a AccountUseCase) DeleteAccount(ctx context.Context, accountId uint) error {
	err := a.AccountRepo.DeleteAccount(ctx, accountId)
	if err != nil {
		return err
	}

	return nil
}
