package service

import (
	"github.com/nats-io/nats.go"
	"github.com/rusystem/crm-accounts/internal/repository"
)

type Service struct {
	Company Company
	Section Sections
	User    User
}

func New(repo *repository.Repository, nc *nats.Conn) *Service {
	return &Service{
		Company: NewCompanyService(repo),
		Section: NewSectionService(repo),
		User:    NewUserService(repo),
	}
}
