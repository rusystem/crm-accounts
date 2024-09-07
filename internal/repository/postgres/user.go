package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/rusystem/crm-accounts/pkg/domain"
	"time"
)

type User interface {
	GetById(ctx context.Context, id int64) (domain.User, error)
	Create(ctx context.Context, user domain.User) (int64, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id int64) error
	GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.User, error)
}

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (upr *UserPostgresRepository) GetById(ctx context.Context, id int64) (domain.User, error) {
	var user domain.User
	var sections []byte

	query := fmt.Sprintf(`
		SELECT 
		    id, company_id, username, name, email, phone, password_hash, 
		    created_at, updated_at, last_login, is_active, role, language, 
		    country, is_approved, is_send_system_notification, sections, position 
		FROM %s WHERE id = $1
		`, domain.UsersTable)

	err := upr.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CompanyID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
		&user.IsActive,
		&user.Role,
		&user.Language,
		&user.Country,
		&user.IsApproved,
		&user.IsSendSystemNotification,
		&sections,
		&user.Position,
	)
	if err != nil {
		return domain.User{}, err
	}

	if err = json.Unmarshal(sections, &user.Sections); err != nil {
		return domain.User{}, fmt.Errorf("failed to unmarshal sections JSON: %v", err)
	}

	return user, nil
}

func (upr *UserPostgresRepository) Create(ctx context.Context, user domain.User) (int64, error) {
	query := fmt.Sprintf(`
        INSERT INTO %s
        (company_id, username, name, email, phone, password_hash, created_at, updated_at, last_login, is_active,
         role, language, country, is_approved, is_send_system_notification, sections, position)
        VALUES
        ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
        RETURNING id`, domain.UsersTable)

	sectionsJSON, err := json.Marshal(user.Sections)
	if err != nil {
		return 0, err
	}

	var id int64
	err = upr.db.QueryRowContext(ctx, query,
		user.CompanyID,
		user.Username,
		user.Name,
		user.Email,
		user.Phone,
		user.PasswordHash,
		time.Now().UTC(),
		time.Now().UTC(),
		user.LastLogin,
		user.IsActive,
		user.Role,
		user.Language,
		user.Country,
		user.IsApproved,
		user.IsSendSystemNotification,
		sectionsJSON,
		user.Position,
	).Scan(&id)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return 0, domain.ErrUserAlreadyExists
			}
		}

		return 0, err
	}

	return id, nil
}

func (upr *UserPostgresRepository) Update(ctx context.Context, user domain.User) error {
	sectionsJSON, err := json.Marshal(user.Sections)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
		UPDATE %s
		SET
		    company_id = $1, username = $2, name = $3, email = $4, phone = $5, password_hash = $6, updated_at = $7,
		    last_login = $8, is_active = $9, role = $10, language = $11, country = $12, is_approved = $13,
		    is_send_system_notification = $14, sections = $15, position = $16
		WHERE id = $17
		`, domain.UsersTable)

	_, err = upr.db.ExecContext(ctx, query,
		user.CompanyID,
		user.Username,
		user.Name,
		user.Email,
		user.Phone,
		user.PasswordHash,
		time.Now().UTC(),
		user.LastLogin,
		user.IsActive,
		user.Role,
		user.Language,
		user.Country,
		user.IsApproved,
		user.IsSendSystemNotification,
		sectionsJSON,
		user.Position,
		user.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (upr *UserPostgresRepository) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, domain.UsersTable)

	_, err := upr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (upr *UserPostgresRepository) GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.User, error) {
	query := fmt.Sprintf(`
		SELECT 
		    id, company_id, username, name, email, phone, password_hash, created_at, 
		    updated_at, last_login, is_active, role, language, country, 
		    is_approved, is_send_system_notification, sections, position
		FROM %s
		WHERE company_id = $1
		`, domain.UsersTable)

	var users []domain.User

	rows, err := upr.db.QueryContext(ctx, query, companyId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var user domain.User
		var b []byte

		if err := rows.Scan(
			&user.ID, &user.CompanyID, &user.Username, &user.Name, &user.Email, &user.Phone, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt,
			&user.LastLogin, &user.IsActive, &user.Role, &user.Language, &user.Country, &user.IsApproved, &user.IsSendSystemNotification,
			&b, &user.Position,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &user.Sections); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
