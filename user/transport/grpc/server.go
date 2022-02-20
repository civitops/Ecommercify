package grpc

import (
	"context"

	"github.com/civitops/Ecommercify/user/implementation/user"
	userPb "github.com/civitops/Ecommercify/user/proto/user"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func newGrpcServer(l *zap.SugaredLogger, u user.Service, t trace.Tracer) *server {
	return &server{
		log:   l,
		user:  u,
		trace: t,
	}
}

type server struct {
	userPb.AuthUserServer
	log   *zap.SugaredLogger
	user  user.Service
	trace trace.Tracer
}

func (s *server) CreateUser(ctx context.Context, u *userPb.User) (*userPb.UserId, error) {
	ctxSpan, span := s.trace.Start(ctx, "create-user-grpc")
	defer span.End()
	usr := user.Entity{
		Name:     u.Name,
		PhoneNo:  u.PhoneNo,
		Email:    u.Email,
		Password: u.Password,
		HomeAddress: user.Address{
			AdressLine: u.HomeAdress.AddressLine,
			PhoneNo:    u.HomeAdress.PhoneNo,
			City:       u.HomeAdress.City,
			PinCode:    u.HomeAdress.PinCode,
			Landmark:   u.HomeAdress.Landmark,
		},
		DeliveryAddress: user.Address{
			AdressLine: u.DeliveryAddress.AddressLine,
			PhoneNo:    u.DeliveryAddress.PhoneNo,
			City:       u.DeliveryAddress.City,
			PinCode:    u.DeliveryAddress.PinCode,
			Landmark:   u.DeliveryAddress.Landmark,
		},
		Role: user.RoleType(u.RoleType),
	}
	id, err := s.user.Create(ctxSpan, usr)
	userId := userPb.UserId{Id: uint32(id)}
	return &userId, err
}

// func (s *server) GetUserByID(ctx context.Context, u *userPb.SignInUser) (*userPb.SignInAck, error) {

// }
