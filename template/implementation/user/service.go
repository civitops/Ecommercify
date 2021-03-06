package user

import (
	"context"

	"github.com/civitops/Ecommercify/user/pkg/config"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, e Entity) (uint, error)
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, where map[string]WhereClause) (Entity, error)
}

type userService struct {
	log   *zap.SugaredLogger
	cfg   config.UserConfig
	repo  Repository
	trace trace.Tracer
}

func NewUserService(l *zap.SugaredLogger, c config.UserConfig, r Repository, t trace.Tracer) Service {
	return &userService{
		log:   l,
		cfg:   c,
		repo:  r,
		trace: t,
	}
}

func (s *userService) Create(ctx context.Context, e Entity) (uint, error) {
	ctxSpan, span := s.trace.Start(ctx, "create-svc-func")
	defer span.End()

	return s.repo.Create(ctxSpan, e)
}

func (s *userService) Update(ctx context.Context, e Entity) error {
	return s.repo.Update(ctx, e)
}

func (s *userService) Delete(ctx context.Context, ID uint) error {
	return s.repo.Delete(ctx, ID)
}

func (s *userService) Get(ctx context.Context, where map[string]WhereClause) (Entity, error) {
	return s.repo.Get(ctx, where)
}
