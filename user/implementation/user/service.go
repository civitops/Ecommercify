package user

import (
	"context"
	"errors"

	"github.com/civitops/Ecommercify/user/pkg/config"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, e Entity) (uint, error)
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, sel string, where map[string]interface{}) (Entity, error)
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

	if e.Name == "" {
		err := errors.New("name is Required")
		errLogWithSpanAttributes("name not Provided inside Create Svc", err, span, s.log)

		return 0, err
	}

	return s.repo.Create(ctxSpan, e)
}

func (s *userService) Update(ctx context.Context, e Entity) error {
	ctxSpan, span := s.trace.Start(ctx, "update-svc-func")
	defer span.End()

	if e.ID == 0 {
		err := errors.New("id is Required")
		errLogWithSpanAttributes("id not provided inside update svc", err, span, s.log)

		return err
	}

	return s.repo.Update(ctxSpan, e)
}

func (s *userService) Delete(ctx context.Context, ID uint) error {
	ctxSpan, span := s.trace.Start(ctx, "delete-svc-func")
	defer span.End()

	return s.repo.Delete(ctxSpan, ID)
}

func (s *userService) Get(ctx context.Context, sel string, where map[string]interface{}) (Entity, error) {
	ctxSpan, span := s.trace.Start(ctx, "get-svc-func")
	defer span.End()

	if sel == "" {
		sel = "*"
	}

	return s.repo.Get(ctxSpan, sel, where)
}

func errLogWithSpanAttributes(msg string, err error, span trace.Span, log *zap.SugaredLogger) {
	// mark span with the error
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())

	// extracting traceID for logging purpose
	traceID := span.SpanContext().TraceID().String()
	log.Errorf(msg+"err: %v", err, zap.String("traceID", traceID))
}
