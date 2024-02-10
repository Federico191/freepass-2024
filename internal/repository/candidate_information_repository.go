package repository

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CandidateInformationRepositoryInterface interface {
	Create(ctx context.Context, req model.CreateCandidateInformation) (*entity.CandidateInformation, error)
	GetByElectionNumber(ctx context.Context, electionNumber uint) (*entity.CandidateInformation, error)
	Update(ctx context.Context, req model.UpdateCandidateInformation) (*entity.CandidateInformation, error)
	Delete(ctx context.Context, electionNumber uint) error
}

type CandidateInformationRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewCandidateInformationRepository(db *gorm.DB, log *logrus.Logger) CandidateInformationRepositoryInterface {
	return CandidateInformationRepository{db: db, log: log}
}

func (c CandidateInformationRepository) Create(ctx context.Context, req model.CreateCandidateInformation) (*entity.CandidateInformation, error) {
	candidateInformation := &entity.CandidateInformation{
		ElectionNumber: req.ElectionNumber,
		Candidate:      req.Candidate,
		Vision:         req.Vision,
		Mission:        req.Mission,
		Achievement:    req.Achievement,
		Experience:     req.Experience,
	}

	err := c.db.WithContext(ctx).Create(candidateInformation).Error
	if err != nil {
		c.log.Error("cannot insert data :", err)
		return &entity.CandidateInformation{}, err
	}

	return candidateInformation, nil

}

func (c CandidateInformationRepository) GetByElectionNumber(ctx context.Context, electionNumber uint) (*entity.CandidateInformation, error) {
	var candidateInformation *entity.CandidateInformation

	err := c.db.WithContext(ctx).Where("election_number = ?", electionNumber).First(&candidateInformation).Error
	if err != nil {
		c.log.Error("election number not found")
		return &entity.CandidateInformation{}, err
	}

	return candidateInformation, nil
}

func (c CandidateInformationRepository) Update(ctx context.Context, req model.UpdateCandidateInformation) (*entity.CandidateInformation, error) {
	candidateInformation, err := c.GetByElectionNumber(ctx, req.ElectionNumber)
	if err != nil {
		c.log.Error("election number not found")
		return &entity.CandidateInformation{}, err
	}

	arg := &entity.CandidateInformation{
		ElectionNumber: candidateInformation.ElectionNumber,
		Candidate:      req.Candidate,
		Vision:         req.Vision,
		Mission:        req.Mission,
		Achievement:    req.Achievement,
		Experience:     req.Achievement,
	}

	err = c.db.WithContext(ctx).Model(&arg).Updates(arg).Error
	if err != nil {
		c.log.Error("cannot update candidate information")
		return &entity.CandidateInformation{}, err
	}

	return candidateInformation, nil
}

func (c CandidateInformationRepository) Delete(ctx context.Context, electionNumber uint) error {
	var candidateInformation entity.CandidateInformation

	err := c.db.WithContext(ctx).Delete(&candidateInformation, electionNumber).Error
	if err != nil {
		c.log.Error("cannot delete post")
		return err
	}

	return nil
}
