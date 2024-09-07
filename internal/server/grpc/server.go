package grpc

import (
	"fmt"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/company"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/sections"
	"github.com/rusystem/crm-accounts/pkg/gen/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type Server struct {
	server         *grpc.Server
	userServer     user.UserServiceServer
	companyServer  company.CompanyServiceServer
	sectionsServer sections.SectionsServiceServer
}

func New(userServer user.UserServiceServer, companyServer company.CompanyServiceServer, sectionsServer sections.SectionsServiceServer) *Server {
	opt := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 100),
		grpc.MaxSendMsgSize(1024 * 1024 * 100),
		grpc.MaxConcurrentStreams(1000),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	}

	return &Server{
		server:         grpc.NewServer(opt...),
		userServer:     userServer,
		companyServer:  companyServer,
		sectionsServer: sectionsServer,
	}
}

func (s *Server) Run(port int64) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	user.RegisterUserServiceServer(s.server, s.userServer)
	company.RegisterCompanyServiceServer(s.server, s.companyServer)
	sections.RegisterSectionsServiceServer(s.server, s.sectionsServer)

	if err = s.server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() func() {
	return s.server.GracefulStop
}
