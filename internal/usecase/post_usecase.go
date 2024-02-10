package usecase

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type PostUseCaseInterface interface {
	Create(ctx context.Context, req model.CreatePost) (*entity.Post, error)
	GetById(ctx context.Context, postId uint) (*entity.Post, error)
	Update(ctx context.Context, req model.UpdatePost) (*entity.Post, error)
	Delete(ctx context.Context, postId uint) error
}

type PostUseCase struct {
	postRepo repository.PostRepositoryInterface
	timeOut  time.Duration
}

func NewPostUseCase(postRepo repository.PostRepositoryInterface, timeOut time.Duration) PostUseCaseInterface {
	return &PostUseCase{postRepo: postRepo, timeOut: timeOut}
}

func (p PostUseCase) Create(ctx context.Context, req model.CreatePost) (*entity.Post, error) {
	post, err := p.postRepo.Create(ctx, req)
	if err != nil {
		return &entity.Post{}, err
	}

	return post, nil
}

func (p PostUseCase) GetById(ctx context.Context, postId uint) (*entity.Post, error) {
	post, err := p.postRepo.GetById(ctx, postId)
	if err != nil {
		return &entity.Post{}, err
	}

	return post, nil
}

func (p PostUseCase) Update(ctx context.Context, req model.UpdatePost) (*entity.Post, error) {
	post, err := p.postRepo.Update(ctx, req)
	if err != nil {
		return &entity.Post{}, err
	}

	return post, nil
}

func (p PostUseCase) Delete(ctx context.Context, postId uint) error {
	err := p.postRepo.Delete(ctx, postId)
	if err != nil {
		return err
	}

	post, _ := p.postRepo.GetById(ctx, postId)
	if post.ID != 0 {
		return err
	}
	return nil
}
