package usecase

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type CandidatePostUseCaseInterface interface {
	CreatePost(ctx context.Context, req model.CreateCandidatePost) (*entity.CandidatePost, error)
	GetPost(ctx context.Context, req model.GetDelCandidatePost) (*entity.CandidatePost, error)
	DeleteCandidatePost(ctx context.Context, req model.GetDelCandidatePost) error
}

type CandidatePostUseCase struct {
	candidatePostRepo repository.CandidatePostRepositoryInterface
	timeOut           time.Duration
}

func NewCandidatePostUseCase(candidatePostRepo repository.CandidatePostRepositoryInterface, timeOut time.Duration) CandidatePostUseCaseInterface {
	return CandidatePostUseCase{candidatePostRepo: candidatePostRepo, timeOut: timeOut}
}

func (c CandidatePostUseCase) CreatePost(ctx context.Context, req model.CreateCandidatePost) (*entity.CandidatePost, error) {
	result, err := c.candidatePostRepo.CreatePost(ctx, req)
	if err != nil {
		return &entity.CandidatePost{}, err
	}

	return result, nil
}

func (c CandidatePostUseCase) GetPost(ctx context.Context, req model.GetDelCandidatePost) (*entity.CandidatePost, error) {
	candidatePost, err := c.candidatePostRepo.GetPost(ctx, req)
	if err != nil {
		return &entity.CandidatePost{}, err
	}

	return candidatePost, nil
}

func (c CandidatePostUseCase) DeleteCandidatePost(ctx context.Context, req model.GetDelCandidatePost) error {
	err := c.candidatePostRepo.DeleteCandidatePost(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
