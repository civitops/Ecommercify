package grpc

import (
	"context"

	"github.com/civitops/Ecommercify/user/implementation/user"
	userPb "github.com/civitops/Ecommercify/user/proto/user"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGrpcServer(l *zap.SugaredLogger, u user.Service, t trace.Tracer) *server {
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
	//TODO: Check conditions
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
	return &userId, status.Errorf(codes.Internal, err.Error())
}

func (s *server) GetUserByID(ctx context.Context, u *userPb.SignInUser) (*userPb.SignInAck, error) {
	ctxSpan, span := s.trace.Start(ctx, "get-user-grpc")
	defer span.End()
	where := make(map[string]interface{})
	if u.Email == "" && u.PhoneNo != "" {
		where["phoneNo"] = u.PhoneNo
	} else if u.Email != "" && u.PhoneNo == "" {
		where["email"] = u.Email
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "email or phoneNo must be provided")
	}
	where["role"] = u.RoleType
	user, err := s.user.Get(ctxSpan, "password", where)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &userPb.SignInAck{IsValid: true, Password: user.Password}, nil
}
