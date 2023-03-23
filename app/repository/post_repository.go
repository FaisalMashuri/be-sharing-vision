package repository

import (
	"be-sharing-vision/app/domain"
	"be-sharing-vision/app/entity"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(ctx context.Context, domain domain.PostDomain) (PostDomain domain.PostDomain, err error)
	GetAllPost(ctx context.Context, offset int, limit int) (allOrder []domain.PostDomain, err error)
	GetById(ctx context.Context, id string) (domain.PostDomain, error)
	UpdatePost(ctx context.Context, id string, post domain.PostDomain) (domain.PostDomain, error)
	DeletePost(ctx context.Context, id string) (err error)
}

func NewPostRepositoryImpl(conn *gorm.DB) PostRepository {
	return &PostRepositoryImpl{Conn: conn}
}

type PostRepositoryImpl struct {
	Conn *gorm.DB
}

func (p PostRepositoryImpl) CreatePost(ctx context.Context, Post domain.PostDomain) (PostDomain domain.PostDomain, err error) {
	//TODO implement me
	var PostModel entity.Post

	createdPost := entity.Post{
		Id:        Post.Id,
		Title:     Post.Title,
		Content:   Post.Content,
		Category:  Post.Category,
		Status:    Post.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = p.Conn.Create(&createdPost).Error
	if err != nil {
		return Post, err
	}
	return PostModel.ToDomain(createdPost), nil
}

func (p PostRepositoryImpl) GetAllPost(ctx context.Context, offset int, limit int) (allOrder []domain.PostDomain, err error) {
	//TODO implement me
	var PostModel []entity.Post

	if err = p.Conn.Raw(fmt.Sprintf("select * from posts limit %d offset %d ", limit, offset)).Scan(&PostModel).Error; err != nil {
		return nil, err
	}
	for _, item := range PostModel {
		allOrder = append(allOrder, item.ToDomain(item))
	}
	return allOrder, nil
}

func (p PostRepositoryImpl) GetById(ctx context.Context, id string) (domain.PostDomain, error) {
	//TODO implement me
	var PostModel entity.Post
	if err := p.Conn.Where("id = ?", id).First(&PostModel).Error; err != nil {
		return domain.PostDomain{}, err
	}
	return PostModel.ToDomain(PostModel), nil
}

func (p PostRepositoryImpl) UpdatePost(ctx context.Context, id string, Post domain.PostDomain) (domain.PostDomain, error) {
	//TODO implement me
	var PostModel entity.Post
	if err := p.Conn.Where("id = ?", id).First(&PostModel).Error; err != nil {
		return domain.PostDomain{}, err
	}
	PostModel.Title = Post.Title
	PostModel.Content = Post.Content
	PostModel.Category = Post.Category
	PostModel.Status = Post.Status
	PostModel.UpdatedAt = time.Now()
	if err := p.Conn.Save(&PostModel).Error; err != nil {
		return domain.PostDomain{}, err
	}
	return PostModel.ToDomain(PostModel), nil

}

func (p PostRepositoryImpl) DeletePost(ctx context.Context, id string) (err error) {
	//TODO implement me
	var PostModel entity.Post
	if err = p.Conn.Where("id = ?", id).Delete(&PostModel).Error; err != nil {
		return
	}
	return
}
