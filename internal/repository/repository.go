package repository

import (
	"database/sql"
	"github.com/rusystem/crm-accounts/internal/config"
)

type Repository struct {
	Company  Company
	User     User
	Sections Sections
}

func New(cfg *config.Config, postgres *sql.DB) *Repository {
	return &Repository{
		Company:  NewCompanyRepository(cfg, postgres),
		User:     NewUserRepository(cfg, postgres),
		Sections: NewSectionsRepository(cfg, postgres),
	}
}
