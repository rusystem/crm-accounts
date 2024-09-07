package service

import (
	"context"
	"github.com/rusystem/crm-accounts/internal/repository"
	"github.com/rusystem/crm-accounts/pkg/domain"
)

type Company interface {
	GetById(ctx context.Context, id int64) (domain.Company, error)
	Create(ctx context.Context, company domain.Company) (int64, error)
	Update(ctx context.Context, company domain.Company) error
	Delete(ctx context.Context, id int64) error
	IsExist(ctx context.Context, id int64) (bool, error)
	List(ctx context.Context) ([]domain.Company, error)
}

type CompanyService struct {
	repo *repository.Repository
}

func NewCompanyService(repo *repository.Repository) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

func (cs *CompanyService) GetById(ctx context.Context, id int64) (domain.Company, error) {
	return cs.repo.Company.GetById(ctx, id)
}

func (cs *CompanyService) Create(ctx context.Context, company domain.Company) (int64, error) {
	return cs.repo.Company.Create(ctx, company)
}

func (cs *CompanyService) Update(ctx context.Context, company domain.Company) error {
	return cs.repo.Company.Update(ctx, company)
}

func (cs *CompanyService) Delete(ctx context.Context, id int64) error {
	return cs.repo.Company.Delete(ctx, id)
}

func (cs *CompanyService) IsExist(ctx context.Context, id int64) (bool, error) {
	return cs.repo.Company.IsExist(ctx, id)
}

func (cs *CompanyService) List(ctx context.Context) ([]domain.Company, error) {
	return cs.repo.Company.List(ctx)
}
