package usecase

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type ElectionPeriodUseCaseInterface interface {
	Create(ctx context.Context, req model.CreateElectionPeriod) (*entity.ElectionPeriod, error)
	Get(ctx context.Context, id uint) (*entity.ElectionPeriod, error)
	Update(ctx context.Context, req model.UpdateElectionPeriod) (*entity.ElectionPeriod, error)
	Delete(ctx context.Context, id uint) error
}

type ElectionPeriodUseCase struct {
	electionPeriodRepo repository.ElectionPeriodRepositoryInterface
	timeOut            time.Duration
}

func NewElectionPeriodUseCase(electionPeriodRepo repository.ElectionPeriodRepositoryInterface, timeOut time.Duration) ElectionPeriodUseCaseInterface {
	return ElectionPeriodUseCase{electionPeriodRepo: electionPeriodRepo, timeOut: timeOut}
}

func (e ElectionPeriodUseCase) Create(ctx context.Context, req model.CreateElectionPeriod) (*entity.ElectionPeriod, error) {
	result, err := e.electionPeriodRepo.Create(ctx, req)
	if err != nil {
		return &entity.ElectionPeriod{}, err
	}

	return result, nil
}

func (e ElectionPeriodUseCase) Get(ctx context.Context, id uint) (*entity.ElectionPeriod, error) {
	result, err := e.electionPeriodRepo.Get(ctx, id)
	if err != nil {
		return &entity.ElectionPeriod{}, err
	}

	return result, nil
}

func (e ElectionPeriodUseCase) Update(ctx context.Context, req model.UpdateElectionPeriod) (*entity.ElectionPeriod, error) {
	result, err := e.electionPeriodRepo.Update(ctx, req)
	if err != nil {
		return &entity.ElectionPeriod{}, err
	}

	return result, nil
}

func (e ElectionPeriodUseCase) Delete(ctx context.Context, id uint) error {
	err := e.electionPeriodRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
