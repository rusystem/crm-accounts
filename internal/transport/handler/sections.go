package handler

import (
	"context"
	"github.com/rusystem/crm-accounts/internal/service"
	"github.com/rusystem/crm-accounts/pkg/domain"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/sections"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SectionsHandler struct {
	service *service.Service
}

func NewSectionsHandler(service *service.Service) *SectionsHandler {
	return &SectionsHandler{
		service: service,
	}
}

func (sh *SectionsHandler) GetById(ctx context.Context, id *sections.SectionsId) (*sections.Section, error) {
	s, err := sh.service.Section.GetById(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	return &sections.Section{
		Id:   s.Id,
		Name: s.Name,
	}, nil
}

func (sh *SectionsHandler) Create(ctx context.Context, section *sections.Section) (*sections.SectionsId, error) {
	id, err := sh.service.Section.Create(ctx, domain.Section{
		Name: section.Name,
	})
	if err != nil {
		return nil, err
	}

	return &sections.SectionsId{
		Id: id,
	}, nil
}

func (sh *SectionsHandler) Update(ctx context.Context, req *sections.Section) (*emptypb.Empty, error) {
	err := sh.service.Section.Update(ctx, domain.Section{
		Id:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (sh *SectionsHandler) Delete(ctx context.Context, req *sections.SectionsId) (*emptypb.Empty, error) {
	if err := sh.service.Section.Delete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (sh *SectionsHandler) GetList(ctx context.Context, _ *emptypb.Empty) (*sections.SectionList, error) {
	list, err := sh.service.Section.List(ctx)
	if err != nil {
		return nil, err
	}

	var listPb []*sections.Section
	for _, section := range list {
		listPb = append(listPb, &sections.Section{
			Id:   section.Id,
			Name: section.Name,
		})
	}

	return &sections.SectionList{
		Sections: listPb,
	}, nil
}
