package user

import (
	"context"

	"github.com/civitops/Ecommercify/user/pkg/config"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, e Entity) (uint, error)
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, ID uint) error
}

type userService struct {
	log  *zap.SugaredLogger
	cfg  config.UserConfig
	repo Repository
}

func NewUserService(l *zap.SugaredLogger, c config.UserConfig, r Repository) Service {
	return &userService{
		log:  l,
		cfg:  c,
		repo: r,
	}
}

func (s *userService) Create(ctx context.Context, e Entity) (uint, error) {

	return s.repo.Create(ctx, e)
}

func (s *userService) Update(ctx context.Context, e Entity) error {
	return s.repo.Update(ctx, e)
}

func (s *userService) Delete(ctx context.Context, ID uint) error {
	return nil
}

func (s *userService) Get(ctx context.Context, ID uint) error {
	return nil
}
