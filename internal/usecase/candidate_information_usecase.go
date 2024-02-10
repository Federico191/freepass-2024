package usecase

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type CandidateInformationInterface interface {
	Create(ctx context.Context, req model.CreateCandidateInformation) (*entity.CandidateInformation, error)
	GetByElectionNumber(ctx context.Context, electionNumber uint) (*entity.CandidateInformation, error)
	Update(ctx context.Context, arg model.UpdateCandidateInformation) (*entity.CandidateInformation, error)
	Delete(ctx context.Context, electionNumber uint) error
}

type CandidateInformationUseCase struct {
	candidateInformationRepo repository.CandidateInformationRepositoryInterface
	timeOut                  time.Duration
}

func NewCandidateInformationUseCase(candidateInformationRepo repository.CandidateInformationRepositoryInterface, timeOut time.Duration) CandidateInformationInterface {
	return CandidateInformationUseCase{candidateInformationRepo: candidateInformationRepo, timeOut: timeOut}
}

func (c CandidateInformationUseCase) Create(ctx context.Context, req model.CreateCandidateInformation) (*entity.CandidateInformation, error) {
	result, err := c.candidateInformationRepo.Create(ctx, req)
	if err != nil {
		return &entity.CandidateInformation{}, err
	}

	return result, nil
}

func (c CandidateInformationUseCase) GetByElectionNumber(ctx context.Context, electionNumber uint) (*entity.CandidateInformation, error) {
	candidate, err := c.candidateInformationRepo.GetByElectionNumber(ctx, electionNumber)
	if err != nil {
		return &entity.CandidateInformation{}, err
	}

	return candidate, nil
}

func (c CandidateInformationUseCase) Update(ctx context.Context, arg model.UpdateCandidateInformation) (*entity.CandidateInformation, error) {
	candidate, err := c.candidateInformationRepo.Update(ctx, arg)
	if err != nil {
		return &entity.CandidateInformation{}, err
	}

	arg.UpdatedAt = time.Now()
	return candidate, nil
}

func (c CandidateInformationUseCase) Delete(ctx context.Context, electionNumber uint) error {
	err := c.candidateInformationRepo.Delete(ctx, electionNumber)
	if err != nil {
		return err
	}

	return nil
}
