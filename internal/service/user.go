package service

import (
	"context"
	"github.com/rusystem/crm-accounts/internal/repository"
	"github.com/rusystem/crm-accounts/pkg/domain"
)

type User interface {
	GetById(ctx context.Context, id int64) (domain.User, error)
	Create(ctx context.Context, user domain.User) (int64, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id int64) error
	GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.User, error)
}

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetById(ctx context.Context, id int64) (domain.User, error) {
	return us.repo.User.GetById(ctx, id)
}

func (us *UserService) Create(ctx context.Context, user domain.User) (int64, error) {
	return us.repo.User.Create(ctx, user)
}

func (us *UserService) Update(ctx context.Context, user domain.User) error {
	return us.repo.User.Update(ctx, user)
}

func (us *UserService) Delete(ctx context.Context, id int64) error {
	return us.repo.User.Delete(ctx, id)
}

func (us *UserService) GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.User, error) {
	return us.repo.User.GetListByCompanyId(ctx, companyId)
}
