package auth

import (
	"context"

	"github.com/civitops/Ecommercify/authN/pkg/config"
	// "golang.org/x/crypto/argon2"
	// usr "github.com/civitops/Ecommercify/user/implementation/user"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Service interface {
	SignIn(ctx context.Context, email, password string) (uint, error)
}

type userService struct {
	log   *zap.SugaredLogger
	cfg   config.UserConfig
	trace trace.Tracer
}

func NewAuthService(l *zap.SugaredLogger, c config.UserConfig, t trace.Tracer) Service {
	return &userService{
		log:   l,
		cfg:   c,
		trace: t,
	}
}

func (s *userService) SignIn(ctx context.Context, email, password string) (uint, error) {
	// key := argon2.IDKey([]byte(s.cfg.ArgonPassword), []byte(s.cfg.ArgonPassword), 1, 64*1024, 4, 32)

	// argon2.

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
