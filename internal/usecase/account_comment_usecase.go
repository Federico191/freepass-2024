package usecase

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type AccountCommentUseCaseInterface interface {
	Create(ctx context.Context, req model.CreateAccountComment) (*entity.AccountComment, error)
	Get(ctx context.Context, req model.GetDelAccountComment) (*entity.AccountComment, error)
	Update(ctx context.Context, req model.UpdateAccountComment) (*entity.AccountComment, error)
	Delete(ctx context.Context, req model.GetDelAccountComment) error
}

type AccountCommentUseCase struct {
	accountCommentRepo repository.AccountCommentRepositoryInterface
	timeOut            time.Duration
}

func NewAccountCommentUseCase(accountCommentRepo repository.AccountCommentRepositoryInterface, timeOut time.Duration) AccountCommentUseCaseInterface {
	return AccountCommentUseCase{accountCommentRepo: accountCommentRepo, timeOut: timeOut}
}

func (a AccountCommentUseCase) Create(ctx context.Context, req model.CreateAccountComment) (*entity.AccountComment, error) {
	comment, err := a.accountCommentRepo.Create(ctx, req)
	if err != nil {
		return &entity.AccountComment{}, err
	}

	return comment, nil
}

func (a AccountCommentUseCase) Get(ctx context.Context, req model.GetDelAccountComment) (*entity.AccountComment, error) {
	comment, err := a.accountCommentRepo.Get(ctx, req)
	if err != nil {
		return &entity.AccountComment{}, err
	}

	return comment, nil
}

func (a AccountCommentUseCase) Update(ctx context.Context, req model.UpdateAccountComment) (*entity.AccountComment, error) {
	comment, err := a.accountCommentRepo.Update(ctx, req)
	if err != nil {
		return &entity.AccountComment{}, err
	}

	req.UpdatedAt = time.Now()
	return comment, nil
}

func (a AccountCommentUseCase) Delete(ctx context.Context, req model.GetDelAccountComment) error {
	err := a.accountCommentRepo.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
