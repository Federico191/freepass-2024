package implementation

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"time"
)

type UserUseCase struct {
	userRepository repository.UserRepository
	timeOut        time.Duration
}

func NewUserUseCase(userRepo repository.UserRepository, timeout time.Duration) usecase.UserUseCase {
	return &UserUseCase{userRepository: userRepo, timeOut: timeout}
}

func (u UserUseCase) Register(ctx context.Context, req model.UserRegister) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	result, err := u.userRepository.Register(ctx, req)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (u UserUseCase) Login(ctx context.Context, req model.UserLogin) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	result, err := u.userRepository.Login(ctx, req)
	if err != nil {
		return entity.User{}, err
	}
	return result, nil
}
