package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rusystem/crm-accounts/pkg/domain"
)

type Sections interface {
	GetById(ctx context.Context, id int64) (domain.Section, error)
	Create(ctx context.Context, section domain.Section) (int64, error)
	Update(ctx context.Context, section domain.Section) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]domain.Section, error)
}

type SectionPostgresRepository struct {
	db *sql.DB
}

func NewSectionsPostgresRepository(db *sql.DB) *SectionPostgresRepository {
	return &SectionPostgresRepository{db: db}
}

func (spr *SectionPostgresRepository) GetById(ctx context.Context, id int64) (domain.Section, error) {
	var section domain.Section

	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id = $1", domain.SectionsTable)

	if err := spr.db.QueryRowContext(ctx, query, id).Scan(&section.Id, &section.Name); err != nil {
		return domain.Section{}, err
	}

	return section, nil
}

func (spr *SectionPostgresRepository) Create(ctx context.Context, section domain.Section) (int64, error) {
	var id int64

	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", domain.SectionsTable)

	if err := spr.db.QueryRowContext(ctx, query, section.Name).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (spr *SectionPostgresRepository) Update(ctx context.Context, section domain.Section) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", domain.SectionsTable)

	_, err := spr.db.ExecContext(ctx, query, section.Name, section.Id)
	if err != nil {
		return err
	}

	return nil
}

func (spr *SectionPostgresRepository) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", domain.SectionsTable)

	_, err := spr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (spr *SectionPostgresRepository) List(ctx context.Context) ([]domain.Section, error) {
	var sections []domain.Section

	query := fmt.Sprintf("SELECT id, name FROM %s", domain.SectionsTable)
	rows, err := spr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var section domain.Section
		if err := rows.Scan(&section.Id, &section.Name); err != nil {
			return nil, err
		}

		sections = append(sections, section)
	}

	return sections, nil
}
