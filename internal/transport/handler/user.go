package handler

import (
	"context"
	"database/sql"
	"github.com/rusystem/crm-accounts/internal/service"
	"github.com/rusystem/crm-accounts/pkg/domain"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/user"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	service *service.Service
}

func NewUserHandler(service *service.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetById(ctx context.Context, req *user.UserId) (*user.User, error) {
	u, err := uh.service.User.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &user.User{
		Id:                       u.ID,
		CompanyId:                u.CompanyID,
		Username:                 u.Username,
		Name:                     u.Name,
		Email:                    u.Email,
		Phone:                    u.Phone,
		PasswordHash:             u.PasswordHash,
		CreatedAt:                timestamppb.New(u.CreatedAt),
		UpdatedAt:                timestamppb.New(u.UpdatedAt),
		LastLogin:                timestamppb.New(u.LastLogin.Time),
		IsActive:                 u.IsActive,
		Role:                     u.Role,
		Language:                 u.Language,
		Country:                  u.Country,
		IsApproved:               u.IsApproved,
		IsSendSystemNotification: u.IsSendSystemNotification,
		Sections:                 u.Sections,
		Position:                 u.Position,
	}, nil
}

func (uh *UserHandler) Create(ctx context.Context, req *user.User) (*user.UserId, error) {
	id, err := uh.service.User.Create(ctx, domain.User{
		CompanyID:                req.CompanyId,
		Username:                 req.Username,
		Name:                     req.Name,
		Email:                    req.Email,
		Phone:                    req.Phone,
		PasswordHash:             req.PasswordHash,
		CreatedAt:                req.CreatedAt.AsTime(),
		UpdatedAt:                req.UpdatedAt.AsTime(),
		LastLogin:                sql.NullTime{},
		IsActive:                 req.IsActive,
		Role:                     req.Role,
		Language:                 req.Language,
		Country:                  req.Country,
		IsApproved:               req.IsApproved,
		IsSendSystemNotification: req.IsSendSystemNotification,
		Sections:                 req.Sections,
		Position:                 req.Position,
	})
	if err != nil {
		return nil, err
	}

	return &user.UserId{Id: id}, nil
}

func (uh *UserHandler) Update(ctx context.Context, req *user.User) (*emptypb.Empty, error) {
	err := uh.service.User.Update(ctx, domain.User{
		ID:                       req.Id,
		CompanyID:                req.CompanyId,
		Username:                 req.Username,
		Name:                     req.Name,
		Email:                    req.Email,
		Phone:                    req.Phone,
		PasswordHash:             req.PasswordHash,
		CreatedAt:                req.CreatedAt.AsTime(),
		UpdatedAt:                req.UpdatedAt.AsTime(),
		IsActive:                 req.IsActive,
		Role:                     req.Role,
		Language:                 req.Language,
		Country:                  req.Country,
		IsApproved:               req.IsApproved,
		IsSendSystemNotification: req.IsSendSystemNotification,
		Sections:                 req.Sections,
		Position:                 req.Position,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (uh *UserHandler) Delete(ctx context.Context, req *user.UserId) (*emptypb.Empty, error) {
	if err := uh.service.User.Delete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (uh *UserHandler) GetListByCompanyId(ctx context.Context, req *user.UserId) (*user.UserList, error) {
	u, err := uh.service.User.GetListByCompanyId(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	var users []*user.User
	for _, v := range u {
		users = append(users, &user.User{
			Id:                       v.ID,
			CompanyId:                v.CompanyID,
			Username:                 v.Username,
			Name:                     v.Name,
			Email:                    v.Email,
			Phone:                    v.Phone,
			PasswordHash:             v.PasswordHash,
			CreatedAt:                timestamppb.New(v.CreatedAt),
			UpdatedAt:                timestamppb.New(v.UpdatedAt),
			LastLogin:                timestamppb.New(v.LastLogin.Time),
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

	return &user.UserList{Users: users}, nil
}
