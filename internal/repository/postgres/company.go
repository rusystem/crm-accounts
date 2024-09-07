package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

type CompanyPostgresRepository struct {
	db *sql.DB
}

func NewCompanyPostgresRepository(db *sql.DB) *CompanyPostgresRepository {
	return &CompanyPostgresRepository{db: db}
}

func (cpr *CompanyPostgresRepository) GetById(ctx context.Context, id int64) (domain.Company, error) {
	var company domain.Company
	query := fmt.Sprintf(`
		SELECT 
		    id, name_ru, name_en, country, address, phone, email, website, 
		    is_active, created_at, updated_at, is_approved, timezone 
		FROM %s WHERE id = $1;
	`, domain.CompaniesTable)

	err := cpr.db.QueryRowContext(ctx, query, id).Scan(
		&company.ID, &company.NameRu, &company.NameEn, &company.Country, &company.Address, &company.Phone, &company.Email,
		&company.Website, &company.IsActive, &company.CreatedAt, &company.UpdatedAt, &company.IsApproved, &company.Timezone,
	)
	if err != nil {
		return domain.Company{}, err
	}

	return company, nil
}

func (cpr *CompanyPostgresRepository) Create(ctx context.Context, company domain.Company) (int64, error) {
	var id int64
	query := fmt.Sprintf(`
		INSERT INTO %s
		(name_ru, name_en, country, address, phone, email, website, is_active, created_at, updated_at, is_approved, timezone)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id;
	`, domain.CompaniesTable)

	err := cpr.db.QueryRowContext(ctx, query,
		company.NameRu, company.NameEn, company.Country, company.Address, company.Phone, company.Email,
		company.Website, company.IsActive, company.CreatedAt, company.UpdatedAt, company.IsApproved, company.Timezone,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (cpr *CompanyPostgresRepository) Update(ctx context.Context, company domain.Company) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET
		    name_ru = $1, name_en = $2, country = $3, address = $4, phone = $5, email = $6,
		    website = $7, is_active = $8, updated_at = $9, is_approved = $10, timezone = $11
		WHERE id = $12;
	`, domain.CompaniesTable)

	_, err := cpr.db.ExecContext(ctx, query,
		company.NameRu, company.NameEn, company.Country, company.Address, company.Phone, company.Email,
		company.Website, company.IsActive, company.UpdatedAt, company.IsApproved, company.Timezone,
		company.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (cpr *CompanyPostgresRepository) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, domain.CompaniesTable)

	_, err := cpr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (cpr *CompanyPostgresRepository) IsExist(ctx context.Context, id int64) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)`, domain.CompaniesTable)

	var exist bool
	if err := cpr.db.QueryRowContext(ctx, query, id).Scan(&exist); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, domain.ErrCompanyNotFound
		}

		return false, err
	}

	return exist, nil
}

func (cpr *CompanyPostgresRepository) List(ctx context.Context) ([]domain.Company, error) {
	var companies []domain.Company

	query := fmt.Sprintf(`
		SELECT 
		    id, name_ru, name_en, country, address, phone, email, website, 
		    is_active, created_at, updated_at, is_approved, timezone 
		FROM %s;
	`, domain.CompaniesTable)

	rows, err := cpr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var company domain.Company

		if err := rows.Scan(
			&company.ID, &company.NameRu, &company.NameEn, &company.Country, &company.Address, &company.Phone, &company.Email,
			&company.Website, &company.IsActive, &company.CreatedAt, &company.UpdatedAt, &company.IsApproved, &company.Timezone,
		); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}
