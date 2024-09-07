package handler

import (
	"context"
	"github.com/rusystem/crm-accounts/internal/service"
	"github.com/rusystem/crm-accounts/pkg/domain"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/company"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CompanyHandler struct {
	service *service.Service
}

func NewCompanyHandler(service *service.Service) *CompanyHandler {
	return &CompanyHandler{
		service: service,
	}
}

func (ch *CompanyHandler) GetById(ctx context.Context, req *company.CompanyId) (*company.Company, error) {
	c, err := ch.service.Company.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &company.Company{
		Id:         c.ID,
		NameRu:     c.NameRu,
		NameEn:     c.NameEn,
		Country:    c.Country,
		Address:    c.Address,
		Phone:      c.Phone,
		Email:      c.Email,
		Website:    c.Website,
		IsActive:   c.IsActive,
		CreatedAt:  timestamppb.New(c.CreatedAt),
		UpdatedAt:  timestamppb.New(c.UpdatedAt),
		IsApproved: c.IsApproved,
		Timezone:   c.Timezone,
	}, nil
}

func (ch *CompanyHandler) Create(ctx context.Context, req *company.Company) (*company.CompanyId, error) {
	id, err := ch.service.Company.Create(ctx, domain.Company{
		NameRu:     req.NameRu,
		NameEn:     req.NameEn,
		Country:    req.Country,
		Address:    req.Address,
		Phone:      req.Phone,
		Email:      req.Email,
		Website:    req.Website,
		IsActive:   req.IsActive,
		CreatedAt:  req.CreatedAt.AsTime(),
		UpdatedAt:  req.UpdatedAt.AsTime(),
		IsApproved: req.IsApproved,
		Timezone:   req.Timezone,
	})
	if err != nil {
		return nil, err
	}

	return &company.CompanyId{Id: id}, nil
}

func (ch *CompanyHandler) Update(ctx context.Context, req *company.Company) (*emptypb.Empty, error) {
	err := ch.service.Company.Update(ctx, domain.Company{
		ID:         req.Id,
		NameRu:     req.NameRu,
		NameEn:     req.NameEn,
		Country:    req.Country,
		Address:    req.Address,
		Phone:      req.Phone,
		Email:      req.Email,
		Website:    req.Website,
		IsActive:   req.IsActive,
		CreatedAt:  req.CreatedAt.AsTime(),
		UpdatedAt:  req.UpdatedAt.AsTime(),
		IsApproved: req.IsApproved,
		Timezone:   req.Timezone,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (ch *CompanyHandler) Delete(ctx context.Context, req *company.CompanyId) (*emptypb.Empty, error) {
	if err := ch.service.Company.Delete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (ch *CompanyHandler) IsExist(ctx context.Context, req *company.CompanyId) (*company.Bool, error) {
	exists, err := ch.service.Company.IsExist(ctx, req.Id)
	if err != nil {
		return &company.Bool{IsExist: exists}, err
	}

	return &company.Bool{IsExist: exists}, nil
}

func (ch *CompanyHandler) GetList(ctx context.Context, _ *emptypb.Empty) (*company.CompanyList, error) {
	companies, err := ch.service.Company.List(ctx)
	if err != nil {
		return nil, err
	}

	var resp []*company.Company
	for _, c := range companies {
		resp = append(resp, &company.Company{
			Id:         c.ID,
			NameRu:     c.NameRu,
			NameEn:     c.NameEn,
			Country:    c.Country,
			Address:    c.Address,
			Phone:      c.Phone,
			Email:      c.Email,
			Website:    c.Website,
			IsActive:   c.IsActive,
			CreatedAt:  timestamppb.New(c.CreatedAt),
			UpdatedAt:  timestamppb.New(c.UpdatedAt),
			IsApproved: c.IsApproved,
			Timezone:   c.Timezone,
		})
	}

	return &company.CompanyList{Companies: resp}, nil
}
