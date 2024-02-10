package usecase

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/entity/token"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUseCaseInterface interface {
	Register(ctx context.Context, req model.UserRegister) (*entity.User, error)
	Login(ctx context.Context, req model.UserLogin) (*entity.User, string, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}

type UserUseCase struct {
	userRepository repository.UserRepositoryInterface
	timeOut        time.Duration
	tokenMaker     token.Maker
}

func NewUserUseCase(userRepo repository.UserRepositoryInterface, timeout time.Duration, tokenMaker token.Maker) UserUseCaseInterface {
	return &UserUseCase{userRepository: userRepo, timeOut: timeout, tokenMaker: tokenMaker}
}

func (u UserUseCase) Register(ctx context.Context, req model.UserRegister) (*entity.User, error) {
	exist, err := u.userRepository.GetByUsername(ctx, req.Username)
	if err != nil {
		return &entity.User{}, err
	}

	if exist.Username != "" {
		return &entity.User{}, err
	}

	result, err := u.userRepository.Create(ctx, req)
	if err != nil {
		return &entity.User{}, err
	}

	return result, nil
}

func (u UserUseCase) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := u.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return &entity.User{}, err
	}

	return user, nil
}

func (u UserUseCase) Login(ctx context.Context, req model.UserLogin) (*entity.User, string, error) {
	result, err := u.userRepository.GetByUsername(ctx, req.Username)
	if err != nil {
		return &entity.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password))
	if err != nil {
		return &entity.User{}, "", err
	}

	createdToken, err := u.tokenMaker.CreateToken(result.Username, time.Hour*24)
	if err != nil {
		return &entity.User{}, "", err
	}

	return result, createdToken, nil
}
