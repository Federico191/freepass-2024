package repository

import (
	"context"
	"errors"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type ElectionPeriodRepositoryInterface interface {
	Create(ctx context.Context, req model.CreateElectionPeriod) (*entity.ElectionPeriod, error)
	Get(ctx context.Context, id uint) (*entity.ElectionPeriod, error)
	Update(ctx context.Context, req model.UpdateElectionPeriod) (*entity.ElectionPeriod, error)
	Delete(ctx context.Context, id uint) error
	GetActiveElectionPeriod(ctx context.Context) (*entity.ElectionPeriod, error)
}

type ElectionPeriodRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewElectionPeriodRepository(db *gorm.DB, log *logrus.Logger) ElectionPeriodRepositoryInterface {
	return ElectionPeriodRepository{db: db, log: log}
}

func (e ElectionPeriodRepository) Create(ctx context.Context, req model.CreateElectionPeriod) (*entity.ElectionPeriod, error) {
	electionPeriod := entity.ElectionPeriod{
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		AdminUsername: req.AdminUsername,
		Admin:         req.Admin,
	}

	err := e.db.WithContext(ctx).Create(electionPeriod).Error
	if err != nil {
		e.log.Error("cannot insert data")
		return &entity.ElectionPeriod{}, err
	}

	return &electionPeriod, nil
}

func (e ElectionPeriodRepository) Get(ctx context.Context, id uint) (*entity.ElectionPeriod, error) {
	var electionPeriod entity.ElectionPeriod

	err := e.db.WithContext(ctx).Where("ID = ?", id).First(&electionPeriod).Error
	if err != nil {
		e.log.Error("post id not found")
		return &entity.ElectionPeriod{}, err
	}

	return &electionPeriod, nil
}

func (e ElectionPeriodRepository) Update(ctx context.Context, req model.UpdateElectionPeriod) (*entity.ElectionPeriod, error) {
	electionPeriod, err := e.Get(ctx, req.ID)
	if err != nil {
		e.log.Error("id not found")
		return &entity.ElectionPeriod{}, err
	}

	arg := entity.ElectionPeriod{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	err = e.db.WithContext(ctx).Model(&electionPeriod).Updates(arg).Error
	if err != nil {
		e.log.Error("cannot update election period")
		return &entity.ElectionPeriod{}, err
	}

	return electionPeriod, nil
}

func (e ElectionPeriodRepository) Delete(ctx context.Context, id uint) error {
	var electionPeriod entity.ElectionPeriod

	err := e.db.WithContext(ctx).Delete(&electionPeriod, id).Error
	if err != nil {
		e.log.Error("cannot delete election period")
		return err
	}

	return nil
}

func (e ElectionPeriodRepository) GetActiveElectionPeriod(ctx context.Context) (*entity.ElectionPeriod, error) {
	var electionPeriod entity.ElectionPeriod

	now := time.Now()
	err := e.db.WithContext(ctx).Where("start_time < ? AND end_time > ?", now, now).First(&electionPeriod).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		e.log.Error("cannot find active election period")
		return nil, err
	}

	return &electionPeriod, nil
}
