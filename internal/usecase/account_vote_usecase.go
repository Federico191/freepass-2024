package usecase

import (
	"context"
	"errors"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type AccountVoteInterface interface {
	Vote(ctx context.Context, req model.CreateAccountVote) (*entity.AccountVote, error)
}
type AccountVoteUseCase struct {
	accountVoteRepo repository.AccountVoteRepositoryInterface
	timeOut         time.Duration
}

func NewAccountVoteUseCase(accountVoteRepo repository.AccountVoteRepositoryInterface, timeOut time.Duration) AccountVoteInterface {
	return AccountVoteUseCase{accountVoteRepo: accountVoteRepo, timeOut: timeOut}
}

func (a AccountVoteUseCase) Vote(ctx context.Context, req model.CreateAccountVote) (*entity.AccountVote, error) {
	check, err := a.accountVoteRepo.GetVoteById(ctx, req.AccountId)
	if err != nil {
		return &entity.AccountVote{}, err
	}

	isVoted := *check.IsVoted

	if isVoted == true {
		return &entity.AccountVote{}, errors.New("account already voted")
	}

	vote, err := a.accountVoteRepo.Vote(ctx, req)
	if err != nil {
		return &entity.AccountVote{}, err
	}

	return vote, nil
}
