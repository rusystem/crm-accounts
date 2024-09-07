package repository

import (
	"context"
	"database/sql"
	"github.com/rusystem/crm-accounts/internal/config"
	"github.com/rusystem/crm-accounts/internal/repository/postgres"
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

type CompanyRepository struct {
	cfg  *config.Config
	psql postgres.Company
}

func NewCompanyRepository(cfg *config.Config, psql *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		cfg:  cfg,
		psql: postgres.NewCompanyPostgresRepository(psql),
	}
}

func (cr *CompanyRepository) GetById(ctx context.Context, id int64) (domain.Company, error) {
	return cr.psql.GetById(ctx, id)
}

func (cr *CompanyRepository) Create(ctx context.Context, company domain.Company) (int64, error) {
	return cr.psql.Create(ctx, company)
}

func (cr *CompanyRepository) Update(ctx context.Context, company domain.Company) error {
	return cr.psql.Update(ctx, company)
}

func (cr *CompanyRepository) Delete(ctx context.Context, id int64) error {
	return cr.psql.Delete(ctx, id)
}

func (cr *CompanyRepository) IsExist(ctx context.Context, id int64) (bool, error) {
	return cr.psql.IsExist(ctx, id)
}

func (cr *CompanyRepository) List(ctx context.Context) ([]domain.Company, error) {
	return cr.psql.List(ctx)
}
