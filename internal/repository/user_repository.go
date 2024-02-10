package repository

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, req model.UserRegister) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}

type UserRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) UserRepositoryInterface {
	return &UserRepository{db: db, log: log}
}

func (u UserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user *entity.User

	err := u.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		u.log.Error("username not found")
		return &entity.User{}, err
	}

	return user, nil
}

func (u UserRepository) Create(ctx context.Context, req model.UserRegister) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &entity.User{}, err
	}

	user := &entity.User{
		Username:  req.Username,
		Email:     req.Email,
		FullName:  req.FullName,
		Password:  string(hashedPassword),
		IsAdmin:   nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err = u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		u.log.Error("cannot insert data: ", err)
		return &entity.User{}, err
	}
	return user, nil
}
