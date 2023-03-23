package usecase

import (
	"be-sharing-vision/app/domain"
	"be-sharing-vision/app/repository"
	"context"
	"time"
)

type PostUsecaseInterface interface {
	CreatePost(ctx context.Context, domain domain.PostDomain) (PostDomain domain.PostDomain, err error)
	GetAllPost(ctx context.Context, domain domain.PostDomain) (allOrder []domain.PostDomain, err error)
	GetByID(ctx context.Context, id int) (domain.PostDomain, error)
	UpdatePost(ctx context.Context, domain domain.PostDomain) (domain.PostDomain, error)
	DeletePost(ctx context.Context, domain domain.PostDomain) (err error)
}

type PostUsecase struct {
	Repo           repository.PostRepository
	ContextTimeout time.Duration
}

func NewPostUsecase(repo repository.PostRepository, timeout time.Duration) *PostUsecase {
	return &PostUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (pu *PostUsecase) CreatePost(ctx context.Context, Post domain.PostDomain) (PostDomain domain.PostDomain, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	PostDomain, err = pu.Repo.CreatePost(ctx, Post)
	if err != nil {
		return
	}
	return
}

func (pu *PostUsecase) GetAllPost(ctx context.Context, offset int, limit int) (allOrder []domain.PostDomain, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	allOrder, err = pu.Repo.GetAllPost(ctx, offset, limit)
	if err != nil {
		return
	}

	return
}

func (pu *PostUsecase) GetPostById(ctx context.Context, codePost string) (PostDomain domain.PostDomain, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()
	PostDomain, err = pu.Repo.GetById(ctx, codePost)
	if err != nil {
		return
	}

	return
}

func (pu *PostUsecase) UpdatePost(ctx context.Context, id string, Post domain.PostDomain) (PostDomain domain.PostDomain, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()
	PostDomain, err = pu.Repo.UpdatePost(ctx, id, Post)
	if err != nil {
		return
	}

	return
}

func (pu *PostUsecase) DeletePost(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	if err = pu.Repo.DeletePost(ctx, id); err != nil {
		return
	}

	return
}
