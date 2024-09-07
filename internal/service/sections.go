package service

import (
	"context"
	"github.com/rusystem/crm-accounts/internal/repository"
	"github.com/rusystem/crm-accounts/pkg/domain"
)

type Sections interface {
	GetById(ctx context.Context, id int64) (domain.Section, error)
	Create(ctx context.Context, section domain.Section) (int64, error)
	Update(ctx context.Context, section domain.Section) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]domain.Section, error)
}

type SectionService struct {
	repo *repository.Repository
}

func NewSectionService(repo *repository.Repository) *SectionService {
	return &SectionService{
		repo: repo,
	}
}

func (ss *SectionService) GetById(ctx context.Context, id int64) (domain.Section, error) {
	return ss.repo.Sections.GetById(ctx, id)
}

func (ss *SectionService) Create(ctx context.Context, section domain.Section) (int64, error) {
	return ss.repo.Sections.Create(ctx, section)
}

func (ss *SectionService) Update(ctx context.Context, section domain.Section) error {
	return ss.repo.Sections.Update(ctx, section)
}

func (ss *SectionService) Delete(ctx context.Context, id int64) error {
	return ss.repo.Sections.Delete(ctx, id)
}

func (ss *SectionService) List(ctx context.Context) ([]domain.Section, error) {
	return ss.repo.Sections.List(ctx)
}
