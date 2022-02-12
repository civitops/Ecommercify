package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, e Entity) (uint, error)
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, whereAnd Entity, whereOR Entity) (Entity, error)
}

type postgresRepo struct {
	log    *zap.SugaredLogger
	conn   *pgx.Conn
	pgConn *gorm.DB
	trace  trace.Tracer
}

func NewPostgresRepo(l *zap.SugaredLogger, c *pgx.Conn, p *gorm.DB, t trace.Tracer) Repository {
	p.AutoMigrate(&Entity{})
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

	result := rp.pgConn.WithContext(ctxSpan).Create(&e)
	if result.Error != nil {
		rp.errLogWithSpanAttributes("err while inserting into users", result.Error, span)
	}

	return e.ID, result.Error
}

func (rp *postgresRepo) Update(ctx context.Context, e Entity) error {

	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()

	r := rp.pgConn.WithContext(ctxSpan).Model(&e).Updates(e)

	if r.Error != nil {
		rp.errLogWithSpanAttributes("err while updating ", r.Error, span)
	}

	return r.Error
}

func (rp *postgresRepo) Delete(ctx context.Context, ID uint) error {

	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()

	result := rp.pgConn.WithContext(ctxSpan).Delete(&Entity{}, ID)
	if result.Error != nil {
		rp.errLogWithSpanAttributes("err while deleting users", result.Error, span)
	}

	return result.Error
}

func (rp *postgresRepo) Get(ctx context.Context, whereAnd Entity, whereOR Entity) (Entity, error) {
	var (
		andRes map[string]interface{}
		orRes  map[string]interface{}
		user   Entity
	)

	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()

	err := mapstructure.Decode(whereAnd, &andRes)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = mapstructure.Decode(whereOR, &orRes)
	if err != nil {
		fmt.Println(err.Error())
	}
	pgResult := rp.pgConn.WithContext(ctxSpan).Where(andRes).Or(orRes).First(&user)

	if pgResult.Error != nil {
		rp.errLogWithSpanAttributes("err while Getting Users", pgResult.Error, span)
	}

	return user, pgResult.Error
}

func (rp *postgresRepo) errLogWithSpanAttributes(msg string, err error, span trace.Span) {
	// mark span with the error
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())

	// extracting traceID for logging purpose
	traceID := span.SpanContext().TraceID().String()
	rp.log.Errorf(msg+"err: %v", err, zap.String("traceID", traceID))
}
