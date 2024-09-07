package grpc

import (
	"context"
	"github.com/rusystem/crm-accounts/pkg/domain"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/sections"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SectionsAccountsClient struct {
	conn    *grpc.ClientConn
	section sections.SectionsServiceClient
}

func NewSectionsAccountsClient(addr string) (*SectionsAccountsClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &SectionsAccountsClient{
		conn:    conn,
		section: sections.NewSectionsServiceClient(conn),
	}, nil
}

func (sc *SectionsAccountsClient) Close() error {
	return sc.conn.Close()
}

func (sc *SectionsAccountsClient) GetById(ctx context.Context, id int64) (domain.Section, error) {
	resp, err := sc.section.GetById(ctx, &sections.SectionsId{Id: id})
	if err != nil {
		return domain.Section{}, err
	}

	return domain.Section{
		Id:   resp.Id,
		Name: resp.Name,
	}, nil
}

func (sc *SectionsAccountsClient) Create(ctx context.Context, section domain.Section) (int64, error) {
	resp, err := sc.section.Create(ctx, &sections.Section{
		Name: section.Name,
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}

func (sc *SectionsAccountsClient) Update(ctx context.Context, section domain.Section) error {
	_, err := sc.section.Update(ctx, &sections.Section{
		Id:   section.Id,
		Name: section.Name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (sc *SectionsAccountsClient) Delete(ctx context.Context, id int64) error {
	_, err := sc.section.Delete(ctx, &sections.SectionsId{Id: id})
	if err != nil {
		return err
	}

	return nil
}

func (sc *SectionsAccountsClient) GetList(ctx context.Context) ([]domain.Section, error) {
	resp, err := sc.section.GetList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var list []domain.Section
	for _, section := range resp.Sections {
		list = append(list, domain.Section{
			Id:   section.Id,
			Name: section.Name,
		})
	}

	return list, nil
}
