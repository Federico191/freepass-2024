package implementation

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) repository.UserRepository {
	return &UserRepository{db: db, log: log}
}

func (u UserRepository) Register(ctx context.Context, req model.UserRegister) (entity.User, error) {
	err := u.db.Where("username = ?", req.Username).First(&req).Error
	if err != nil {
		u.log.Error(err)
		return entity.User{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
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
		return entity.User{}, err
	}
	return user, nil
}

func (u UserRepository) Login(ctx context.Context, req model.UserLogin) (entity.User, error) {
	user := entity.User{}

	err := u.db.WithContext(ctx).Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		u.log.Error(err)
		return entity.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		u.log.Error(err)
		return entity.User{}, err
	}

	return user, nil
}
