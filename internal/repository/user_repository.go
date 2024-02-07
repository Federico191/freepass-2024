package repository

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
)

type UserRepository interface {
	Register(ctx context.Context, req model.UserRegister) (entity.User, error)
	Login(ctx context.Context, req model.UserLogin) (entity.User, error)
}
