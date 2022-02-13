package user

import (
	"context"

	"github.com/jackc/pgx/v4"
	"go.opentelemetry.io/otel/codes"
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
	conn   *pgx.Conn
	pgConn *gorm.DB
	trace  trace.Tracer
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Entity) TableName() string {
	return "users"
}

func NewPostgresRepo(l *zap.SugaredLogger, c *pgx.Conn, p *gorm.DB, t trace.Tracer) Repository {
	return &postgresRepo{
		log:    l,
		conn:   c,
		pgConn: p,
		trace:  t,
	}
}

func (rp *postgresRepo) Create(ctx context.Context, e Entity) (uint, error) {
	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()

	err := rp.pgConn.WithContext(ctxSpan).Table("users").Create(&e).Error
	if err != nil {
		rp.errLogWithSpanAttributes("err while inserting into users", err, span)
	}

	return e.ID, err
}

func (rp *postgresRepo) Update(ctx context.Context, e Entity) error {
	ctxSpan, span := rp.trace.Start(ctx, "update-repo-func")
	defer span.End()

	err := rp.pgConn.WithContext(ctxSpan).Table("users").Updates(e).Error
	if err != nil {
		rp.errLogWithSpanAttributes("err while updating ", err, span)
	}

	return err
}

func (rp *postgresRepo) Delete(ctx context.Context, ID uint) error {
	ctxSpan, span := rp.trace.Start(ctx, "delete-repo-func")
	defer span.End()

	err := rp.pgConn.WithContext(ctxSpan).Table("users").Delete(&Entity{}, ID).Error
	if err != nil {
		rp.errLogWithSpanAttributes("err while deleting users", err, span)
	}

	return err
}

func (rp *postgresRepo) Get(ctx context.Context, sel string, where map[string]interface{}) (Entity, error) {
	ctxSpan, span := rp.trace.Start(ctx, "get-repo-func")
	defer span.End()

	var user Entity
	err := rp.pgConn.WithContext(ctxSpan).Table("users").Select(sel).Where(where).First(&user).Error
	if err != nil {
		rp.errLogWithSpanAttributes("err while Getting Users", err, span)
	}

	return user, err
}

func (rp *postgresRepo) errLogWithSpanAttributes(msg string, err error, span trace.Span) {
	// mark span with the error
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())

	// extracting traceID for logging purpose
	traceID := span.SpanContext().TraceID().String()
	rp.log.Errorf(msg+"err: %v", err, zap.String("traceID", traceID))
}
