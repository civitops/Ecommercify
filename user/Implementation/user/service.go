package user

import (
	"context"

	"github.com/civitops/Ecommercify/user/pkg/config"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, e Entity) error
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, ID uint) error
}

type userService struct {
	log *zap.SugaredLogger
	cfg config.UserConfig
}

func NewUserService(l *zap.SugaredLogger, c config.UserConfig) Service {
	return &userService{
		log: l,
		cfg: c,
	}
}

func (s *userService) Create(ctx context.Context, e Entity) error {
	return nil
}

func (s *userService) Update(ctx context.Context, e Entity) error {
	return nil
}

func (s *userService) Delete(ctx context.Context, ID uint) error {
	return nil
}

func (s *userService) Get(ctx context.Context, ID uint) error {
	return nil
}
