package repository

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountVoteRepositoryInterface interface {
	Vote(ctx context.Context, req model.CreateAccountVote) (*entity.AccountVote, error)
	GetVoteById(ctx context.Context, accountId uint) (*entity.AccountVote, error)
}

type AccountVoteRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewAccountVoteRepository(db *gorm.DB, log *logrus.Logger) AccountVoteRepositoryInterface {
	return AccountVoteRepository{db: db, log: log}
}

func (a AccountVoteRepository) Vote(ctx context.Context, req model.CreateAccountVote) (*entity.AccountVote, error) {
	accountVote := &entity.AccountVote{
		AccountId:           req.AccountId,
		Account:             req.Account,
		VotedElectionNumber: req.VotedElectionNumber,
		Candidate:           req.Candidate,
	}

	err := a.db.WithContext(ctx).Create(accountVote).Error
	if err != nil {
		a.log.Error("failed to submit the vote")
		return &entity.AccountVote{}, err
	}

	return accountVote, nil
}

func (a AccountVoteRepository) GetVoteById(ctx context.Context, accountId uint) (*entity.AccountVote, error) {
	var accountVote *entity.AccountVote

	err := a.db.WithContext(ctx).Where("ID = ?", accountId).First(&accountVote).Error
	if err != nil {
		a.log.Error("account vote id not found")
		return &entity.AccountVote{}, err
	}

	return accountVote, nil
}
