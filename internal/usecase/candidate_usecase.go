package usecase

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type CandidateUseCaseInterface interface {
	Create(ctx context.Context, req model.CreateCandidate) (*entity.Candidate, error)
	GetById(ctx context.Context, candidateId *uint) (*entity.Candidate, error)
	Update(ctx context.Context, req model.UpdateCandidate) (*entity.Candidate, error)
	DeleteCandidate(ctx context.Context, candidateId uint) error
}

type CandidateUseCase struct {
	candidateRepo repository.CandidateRepositoryInterface
	timeOut       time.Duration
}

func NewCandidateUseCase(candidateRepo repository.CandidateRepositoryInterface, timeOut time.Duration) CandidateUseCaseInterface {
	return CandidateUseCase{candidateRepo: candidateRepo, timeOut: timeOut}
}

func (c CandidateUseCase) Create(ctx context.Context, req model.CreateCandidate) (*entity.Candidate, error) {
	candidate := &entity.Candidate{
		ElectionNumber: req.ElectionNumber,
		AccountId:      req.AccountId,
		Account:        req.Account,
		LeaderId:       req.LeaderId,
		Candidate:      req.Candidate,
	}

	if candidate.LeaderId != nil {
		leader, err := c.candidateRepo.GetCandidateById(ctx, candidate.LeaderId)
		if err != nil {
			return &entity.Candidate{}, err
		}
		vice := &entity.Candidate{
			ElectionNumber: leader.ElectionNumber,
			AccountId:      candidate.AccountId,
			Account:        candidate.Account,
			LeaderId:       &leader.ID,
			Candidate:      leader,
		}
		return vice, nil
	}

	return candidate, nil
}

func (c CandidateUseCase) GetById(ctx context.Context, candidateId *uint) (*entity.Candidate, error) {
	result, err := c.candidateRepo.GetCandidateById(ctx, candidateId)
	if err != nil {
		return &entity.Candidate{}, err
	}

	return result, nil
}

func (c CandidateUseCase) Update(ctx context.Context, req model.UpdateCandidate) (*entity.Candidate, error) {
	result, err := c.candidateRepo.Update(ctx, req)
	if err != nil {
		return &entity.Candidate{}, err
	}

	return result, err
}

func (c CandidateUseCase) DeleteCandidate(ctx context.Context, candidateId uint) error {
	err := c.candidateRepo.DeleteCandidate(ctx, candidateId)
	if err != nil {
		return err
	}

	return nil
}
