package repository

import (
	"context"
	"database/sql"
	"github.com/rusystem/crm-accounts/internal/config"
	"github.com/rusystem/crm-accounts/internal/repository/postgres"
	"github.com/rusystem/crm-accounts/pkg/domain"
)

type User interface {
	GetById(ctx context.Context, id int64) (domain.User, error)
	Create(ctx context.Context, user domain.User) (int64, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id int64) error
	GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.User, error)
}

type UserRepository struct {
	cfg  *config.Config
	psql postgres.User
}

func NewUserRepository(cfg *config.Config, db *sql.DB) *UserRepository {
	return &UserRepository{
		cfg:  cfg,
		psql: postgres.NewUserPostgresRepository(db),
	}
}

func (r *UserRepository) GetById(ctx context.Context, id int64) (domain.User, error) {
	return r.psql.GetById(ctx, id)
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) (int64, error) {
	return r.psql.Create(ctx, user)
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) error {
	return r.psql.Update(ctx, user)
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.psql.Delete(ctx, id)
}

func (r *UserRepository) GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.User, error) {
	return r.psql.GetListByCompanyId(ctx, companyId)
}
