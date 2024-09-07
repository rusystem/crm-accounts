package transport

import (
	"github.com/rusystem/crm-accounts/internal/service"
	"github.com/rusystem/crm-accounts/internal/transport/handler"
)

type Handler struct {
	User     *handler.UserHandler
	Company  *handler.CompanyHandler
	Sections *handler.SectionsHandler
}

func New(service *service.Service) *Handler {
	return &Handler{
		User:     handler.NewUserHandler(service),
		Company:  handler.NewCompanyHandler(service),
		Sections: handler.NewSectionsHandler(service),
	}
}
