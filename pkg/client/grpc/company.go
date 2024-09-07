package grpc

import (
	"context"
	"github.com/rusystem/crm-accounts/pkg/domain"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/company"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CompanyAccountsClient struct {
	conn    *grpc.ClientConn
	company company.CompanyServiceClient
}

func NewCompanyAccountsClient(addr string) (*CompanyAccountsClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &CompanyAccountsClient{
		conn:    conn,
		company: company.NewCompanyServiceClient(conn),
	}, nil
}

func (cc *CompanyAccountsClient) Close() error {
	return cc.conn.Close()
}

func (cc *CompanyAccountsClient) GetById(ctx context.Context, id int64) (domain.Company, error) {
	resp, err := cc.company.GetById(ctx, &company.CompanyId{Id: id})
	if err != nil {
		return domain.Company{}, err
	}

	return domain.Company{
		ID:         resp.Id,
		NameRu:     resp.NameRu,
		NameEn:     resp.NameEn,
		Country:    resp.Country,
		Address:    resp.Address,
		Phone:      resp.Phone,
		Email:      resp.Email,
		Website:    resp.Website,
		IsActive:   resp.IsActive,
		CreatedAt:  resp.CreatedAt.AsTime(),
		UpdatedAt:  resp.UpdatedAt.AsTime(),
		IsApproved: resp.IsApproved,
		Timezone:   resp.Timezone,
	}, nil
}

func (cc *CompanyAccountsClient) Create(ctx context.Context, c domain.Company) (int64, error) {
	resp, err := cc.company.Create(ctx, &company.Company{
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
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}

func (cc *CompanyAccountsClient) Update(ctx context.Context, c domain.Company) error {
	_, err := cc.company.Update(ctx, &company.Company{
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

	return err
}

func (cc *CompanyAccountsClient) Delete(ctx context.Context, id int64) error {
	_, err := cc.company.Delete(ctx, &company.CompanyId{Id: id})
	return err
}

func (cc *CompanyAccountsClient) IsExist(ctx context.Context, id int64) (bool, error) {
	exist, err := cc.company.IsExist(ctx, &company.CompanyId{Id: id})
	if err != nil {
		return exist.IsExist, err
	}

	return exist.IsExist, nil
}

func (cc *CompanyAccountsClient) GetList(ctx context.Context) ([]domain.Company, error) {
	resp, err := cc.company.GetList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var companies []domain.Company
	for _, c := range resp.Companies {
		companies = append(companies, domain.Company{
			ID:         c.Id,
			NameRu:     c.NameRu,
			NameEn:     c.NameEn,
			Country:    c.Country,
			Address:    c.Address,
			Phone:      c.Phone,
			Email:      c.Email,
			Website:    c.Website,
			IsActive:   c.IsActive,
			CreatedAt:  c.CreatedAt.AsTime(),
			UpdatedAt:  c.UpdatedAt.AsTime(),
			IsApproved: c.IsApproved,
			Timezone:   c.Timezone,
		})
	}

	return companies, nil
}
