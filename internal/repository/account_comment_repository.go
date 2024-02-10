package repository

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountCommentRepositoryInterface interface {
	Create(ctx context.Context, req model.CreateAccountComment) (*entity.AccountComment, error)
	Get(ctx context.Context, req model.GetDelAccountComment) (*entity.AccountComment, error)
	Update(ctx context.Context, req model.UpdateAccountComment) (*entity.AccountComment, error)
	Delete(ctx context.Context, req model.GetDelAccountComment) error
}

type AccountCommentRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewAccountCommentRepository(db *gorm.DB, log *logrus.Logger) AccountCommentRepositoryInterface {
	return AccountCommentRepository{db: db, log: log}
}

func (a AccountCommentRepository) Create(ctx context.Context, req model.CreateAccountComment) (*entity.AccountComment, error) {
	accountComment := &entity.AccountComment{
		PostId:      req.PostId,
		Post:        req.Post,
		CandidateId: req.CandidateId,
		Candidate:   req.Candidate,
		AccountId:   req.AccountId,
		Account:     req.Account,
		Comment:     req.Comment,
	}

	err := a.db.WithContext(ctx).Create(accountComment).Error
	if err != nil {
		a.log.Error("cannot insert account comment")
		return &entity.AccountComment{}, err
	}

	return accountComment, nil
}

func (a AccountCommentRepository) Get(ctx context.Context, req model.GetDelAccountComment) (*entity.AccountComment, error) {
	var accountComment *entity.AccountComment

	err := a.db.WithContext(ctx).
		Where("account_id = ? AND candidate_id = ? AND post_id = ?", req.AccountId, req.CandidateId, req.PostId).
		First(&accountComment).Error
	if err != nil {
		a.log.Error("id not found")
		return &entity.AccountComment{}, err
	}

	return accountComment, nil
}

func (a AccountCommentRepository) Update(ctx context.Context, req model.UpdateAccountComment) (*entity.AccountComment, error) {
	reqGet := model.GetDelAccountComment{
		PostId:      req.PostId,
		CandidateId: req.CandidateId,
		AccountId:   req.AccountId,
	}

	accountComment, err := a.Get(ctx, reqGet)
	if err != nil {
		a.log.Error("post id not found")
		return &entity.AccountComment{}, err
	}

	arg := entity.AccountComment{
		Comment: req.Comment,
	}

	err = a.db.WithContext(ctx).Model(&accountComment).Update("comment", arg).Error
	if err != nil {
		a.log.Error("cannot update account comment")
		return &entity.AccountComment{}, err
	}

	return accountComment, nil
}

func (a AccountCommentRepository) Delete(ctx context.Context, req model.GetDelAccountComment) error {
	var accountComment *entity.AccountComment

	err := a.db.WithContext(ctx).Delete(&accountComment, req).Error
	if err != nil {
		a.log.Error("cannot delete account comment")
		return err
	}

	return nil
}
