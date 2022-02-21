package auth

import (
	"context"

	"github.com/civitops/Ecommercify/auth/pkg/config"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Service interface {
	SignIn(ctx context.Context, e Entity) (uint, error)
}

type userService struct {
	log   *zap.SugaredLogger
	cfg   config.UserConfig
	repo  Repository
	trace trace.Tracer
}

func NewAuthService(l *zap.SugaredLogger, c config.UserConfig, r Repository, t trace.Tracer) Service {
	return &userService{
		log:   l,
		cfg:   c,
		repo:  r,
		trace: t,
	}
}

func (s *userService) SignIn(ctx context.Context, e Entity) (uint, error) {
	return 0, nil
}
func errLogWithSpanAttributes(msg string, err error, span trace.Span, log *zap.SugaredLogger) {
	// mark span with the error
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())

	// extracting traceID for logging purpose
	traceID := span.SpanContext().TraceID().String()
	log.Errorf(msg+"err: %v", err, zap.String("traceID", traceID))
}
