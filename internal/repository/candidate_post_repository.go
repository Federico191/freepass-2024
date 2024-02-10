package repository

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CandidatePostRepositoryInterface interface {
	CreatePost(ctx context.Context, req model.CreateCandidatePost) (*entity.CandidatePost, error)
	GetPost(ctx context.Context, req model.GetDelCandidatePost) (*entity.CandidatePost, error)
	DeleteCandidatePost(ctx context.Context, req model.GetDelCandidatePost) error
}

type CandidatePostRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewCandidatePostRepository(db *gorm.DB, log *logrus.Logger) CandidatePostRepositoryInterface {
	return CandidatePostRepository{db: db, log: log}
}

func (c CandidatePostRepository) CreatePost(ctx context.Context, req model.CreateCandidatePost) (*entity.CandidatePost, error) {
	candidatePost := &entity.CandidatePost{
		PostId:      req.PostId,
		Post:        req.Post,
		CandidateId: req.CandidateId,
		Candidate:   req.Candidate,
	}

	err := c.db.WithContext(ctx).Create(candidatePost).Error
	if err != nil {
		c.log.Error("cannot insert data")
		return &entity.CandidatePost{}, err
	}

	return candidatePost, nil
}

func (c CandidatePostRepository) GetPost(ctx context.Context, req model.GetDelCandidatePost) (*entity.CandidatePost, error) {
	var candidatePost entity.CandidatePost

	err := c.db.WithContext(ctx).Where("candidateId = ? AND postId = ? ", req.CandidateId, req.PostId).First(&candidatePost).Error
	if err != nil {
		c.log.Error("candidate post not found")
		return &entity.CandidatePost{}, err
	}

	return &candidatePost, nil
}

func (c CandidatePostRepository) DeleteCandidatePost(ctx context.Context, req model.GetDelCandidatePost) error {
	var candidatePost *entity.CandidatePost

	err := c.db.WithContext(ctx).Where("candidateId = ? AND postId = ? ", req.CandidateId, req.PostId).Delete(&candidatePost).Error
	if err != nil {
		c.log.Error("cannot delete candidate`s post")
		return err
	}

	return nil
}
