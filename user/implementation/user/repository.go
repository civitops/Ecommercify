package user

import (
	"context"

	"github.com/civitops/Ecommercify/user/pkg/config"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, e Entity) (uint, error)
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, sel string, where map[string]interface{}) (Entity, error)
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

func (rp *postgresRepo) Create(ctx context.Context, e Entity) (uint, error) {
	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()

	err := rp.pgConn.WithContext(ctxSpan).Table(e.TableName()).Create(&e).Error
	if err != nil {
		errLogWithSpanAttributes("err while inserting into users", err, span, rp.log)
	}

	return e.ID, err
}

func (rp *postgresRepo) Update(ctx context.Context, e Entity) error {
	ctxSpan, span := rp.trace.Start(ctx, "update-repo-func")
	defer span.End()

	err := rp.pgConn.WithContext(ctxSpan).Table(e.TableName()).Updates(e).Error
	if err != nil {
		errLogWithSpanAttributes("err while updating ", err, span, rp.log)
	}

	return err
}

func (rp *postgresRepo) Delete(ctx context.Context, ID uint) error {
	ctxSpan, span := rp.trace.Start(ctx, "delete-repo-func")
	defer span.End()

	err := rp.pgConn.WithContext(ctxSpan).Table(config.UserTable).Delete(&Entity{}, ID).Error
	if err != nil {
		errLogWithSpanAttributes("err while deleting users", err, span, rp.log)
	}

	return err
}

func (rp *postgresRepo) Get(ctx context.Context, sel string, where map[string]interface{}) (Entity, error) {
	ctxSpan, span := rp.trace.Start(ctx, "get-repo-func")
	defer span.End()

	var user Entity
	err := rp.pgConn.WithContext(ctxSpan).Table(config.UserTable).Select(sel).Where(where).First(&user).Error
	if err != nil {
		errLogWithSpanAttributes("err while Getting Users", err, span, rp.log)
	}

	return user, err
}
