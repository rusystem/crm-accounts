package grpc

import (
	"context"
	"database/sql"
	"github.com/rusystem/crm-accounts/pkg/domain"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserAccountsClient struct {
	conn *grpc.ClientConn
	user user.UserServiceClient
}

func NewUserAccountsClient(addr string) (*UserAccountsClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &UserAccountsClient{
		conn: conn,
		user: user.NewUserServiceClient(conn),
	}, nil
}

func (uc *UserAccountsClient) Close() error {
	return uc.conn.Close()
}

func (uc *UserAccountsClient) GetById(ctx context.Context, id int64) (domain.User, error) {
	resp, err := uc.user.GetById(ctx, &user.UserId{Id: id})
	if err != nil {
		return domain.User{}, err
	}

	var lastLogin sql.NullTime

	if !resp.LastLogin.AsTime().IsZero() {
		lastLogin = sql.NullTime{
			Time:  resp.LastLogin.AsTime(),
			Valid: true,
		}
	}

	return domain.User{
		ID:                       resp.Id,
		CompanyID:                resp.CompanyId,
		Username:                 resp.Username,
		Name:                     resp.Name,
		Email:                    resp.Email,
		Phone:                    resp.Phone,
		PasswordHash:             resp.PasswordHash,
		CreatedAt:                resp.CreatedAt.AsTime(),
		UpdatedAt:                resp.UpdatedAt.AsTime(),
		LastLogin:                lastLogin,
		IsActive:                 resp.IsActive,
		Role:                     resp.Role,
		Language:                 resp.Language,
		Country:                  resp.Country,
		IsApproved:               resp.IsApproved,
		IsSendSystemNotification: resp.IsSendSystemNotification,
		Sections:                 resp.Sections,
		Position:                 resp.Position,
	}, nil
}

func (uc *UserAccountsClient) Create(ctx context.Context, us domain.User) (int64, error) {
	resp, err := uc.user.Create(ctx, &user.User{
		CompanyId:                us.CompanyID,
		Username:                 us.Username,
		Name:                     us.Name,
		Email:                    us.Email,
		Phone:                    us.Phone,
		PasswordHash:             us.PasswordHash,
		CreatedAt:                timestamppb.New(us.CreatedAt),
		UpdatedAt:                timestamppb.New(us.UpdatedAt),
		IsActive:                 us.IsActive,
		Role:                     us.Role,
		Language:                 us.Language,
		Country:                  us.Country,
		IsApproved:               us.IsApproved,
		IsSendSystemNotification: us.IsSendSystemNotification,
		Sections:                 us.Sections,
		Position:                 us.Position,
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}

func (uc *UserAccountsClient) Update(ctx context.Context, us domain.User) error {
	_, err := uc.user.Update(ctx, &user.User{
		Id:                       us.ID,
		CompanyId:                us.CompanyID,
		Username:                 us.Username,
		Name:                     us.Name,
		Email:                    us.Email,
		Phone:                    us.Phone,
		PasswordHash:             us.PasswordHash,
		CreatedAt:                timestamppb.New(us.CreatedAt),
		UpdatedAt:                timestamppb.New(us.UpdatedAt),
		IsActive:                 us.IsActive,
		Role:                     us.Role,
		Language:                 us.Language,
		Country:                  us.Country,
		IsApproved:               us.IsApproved,
		IsSendSystemNotification: us.IsSendSystemNotification,
		Sections:                 us.Sections,
		Position:                 us.Position,
	})

	return err
}

func (uc *UserAccountsClient) Delete(ctx context.Context, id int64) error {
	_, err := uc.user.Delete(ctx, &user.UserId{Id: id})

	return err
}

func (uc *UserAccountsClient) GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.User, error) {
	resp, err := uc.user.GetListByCompanyId(ctx, &user.UserId{Id: companyId})
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, v := range resp.Users {
		var lastLogin sql.NullTime

		if !v.LastLogin.AsTime().IsZero() {
			lastLogin = sql.NullTime{
				Time:  v.LastLogin.AsTime(),
				Valid: true,
			}
		}

		users = append(users, domain.User{
			ID:                       v.Id,
			CompanyID:                v.CompanyId,
			Username:                 v.Username,
			Name:                     v.Name,
			Email:                    v.Email,
			Phone:                    v.Phone,
			PasswordHash:             v.PasswordHash,
			CreatedAt:                v.CreatedAt.AsTime(),
			UpdatedAt:                v.UpdatedAt.AsTime(),
			LastLogin:                lastLogin,
			IsActive:                 v.IsActive,
			Role:                     v.Role,
			Language:                 v.Language,
			Country:                  v.Country,
			IsApproved:               v.IsApproved,
			IsSendSystemNotification: v.IsSendSystemNotification,
			Sections:                 v.Sections,
			Position:                 v.Position,
		})
	}

	return users, nil
}
