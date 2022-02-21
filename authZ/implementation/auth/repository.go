package auth

import (
	"context"

	"github.com/civitops/Ecommercify/auth/pkg/config"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	SignIn(ctx context.Context, e Entity) (uint, error)
}

type postgresRepo struct {
	log    *zap.SugaredLogger
	pgConn *gorm.DB
	trace  trace.Tracer
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Entity) TableName() string {
	return config.UserTable
}

func NewPostgresRepo(l *zap.SugaredLogger, p *gorm.DB, t trace.Tracer) Repository {
	return &postgresRepo{
		log:    l,
		pgConn: p,
		trace:  t,
	}
}

func (rp *postgresRepo) SignIn(ctx context.Context, e Entity) (uint, error) {
	return 0, nil
}
