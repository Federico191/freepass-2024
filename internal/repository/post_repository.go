package repository

import (
	"context"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostRepositoryInterface interface {
	Create(ctx context.Context, req model.CreatePost) (*entity.Post, error)
	GetById(ctx context.Context, postId uint) (*entity.Post, error)
	Update(ctx context.Context, req model.UpdatePost) (*entity.Post, error)
	Delete(ctx context.Context, postId uint) error
}

type PostRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewPostRepository(db *gorm.DB, log *logrus.Logger) PostRepositoryInterface {
	return PostRepository{db: db, log: log}
}

func (p PostRepository) Create(ctx context.Context, req model.CreatePost) (*entity.Post, error) {
	post := &entity.Post{
		Description: req.Description,
		PicUrl:      &req.Description,
	}

	err := p.db.WithContext(ctx).Create(post).Error
	if err != nil {
		p.log.Error("cannot insert data")
		return &entity.Post{}, err
	}

	return post, nil
}

func (p PostRepository) GetById(ctx context.Context, postId uint) (*entity.Post, error) {
	var post *entity.Post

	err := p.db.WithContext(ctx).Where("ID = ?", postId).First(&post).Error
	if err != nil {
		p.log.Error("post id not found")
		return &entity.Post{}, err
	}

	return post, nil
}

func (p PostRepository) Update(ctx context.Context, req model.UpdatePost) (*entity.Post, error) {
	post, err := p.GetById(ctx, req.ID)
	if err != nil {
		p.log.Error("post id not found")
		return &entity.Post{}, err
	}

	arg := &entity.Post{
		Description: req.Description,
		PicUrl:      req.PicUrl,
	}

	err = p.db.WithContext(ctx).Model(&post).Updates(arg).Error
	if err != nil {
		p.log.Error("cannot update post")
		return &entity.Post{}, err
	}

	return post, nil
}

func (p PostRepository) Delete(ctx context.Context, postId uint) error {
	var post *entity.Post

	err := p.db.WithContext(ctx).Delete(&post, postId).Error
	if err != nil {
		p.log.Error("cannot delete post")
		return err
	}

	return nil
}
