package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Repository interface {
	Create(ctx context.Context, e Entity) (uint, error)
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, where map[string]WhereClause) (Entity, error)
}
type WhereClause struct {
	Condition string
	Value     interface{}
}

type postgresRepo struct {
	log   *zap.SugaredLogger
	conn  *pgx.Conn
	trace trace.Tracer
}

func NewPostgresRepo(l *zap.SugaredLogger, c *pgx.Conn, t trace.Tracer) Repository {
	return &postgresRepo{
		log:   l,
		conn:  c,
		trace: t,
	}
}

func (rp *postgresRepo) Create(ctx context.Context, e Entity) (uint, error) {
	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()

	// sql statement to insert row
	stmt := `INSERT INTO users (
		name,email,phone_no,
		homeaddress_phoneno,homeaddress_address_line,
		homeaddress_city,homeaddress_pin_code,
		homeaddress_landmark,is_admin) VALUES 
		($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id
		`
	var insertedID uint
	err := rp.conn.QueryRow(ctxSpan, stmt, e.Name, e.Email, e.PhoneNo,
		e.HomeAddress.PhoneNo, e.HomeAddress.AdressLine,
		e.HomeAddress.City, e.HomeAddress.PinCode,
		e.HomeAddress.Landmark, e.IsAdmin).Scan(&insertedID)

	if err != nil {
		rp.errLogWithSpanAttributes("err while inserting into users", err, span)
	}

	return insertedID, err
}

func (rp *postgresRepo) Update(ctx context.Context, e Entity) error {
	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()
	var (
		result     map[string]interface{}
		updateStmt string
		values     []interface{}
	)
	err := mapstructure.Decode(e, &result)
	if err != nil {
		fmt.Println(err.Error())
	}
	i := 0
	len := len(result)
	for k := range result {
		i++
		if i == len {
			updateStmt = updateStmt + fmt.Sprintf(" %s=$%d", k, i)
		} else {
			updateStmt = updateStmt + fmt.Sprintf(" %s=$%d,", k, i)
		}

		values = append(values, result[k])
	}

	stmt := fmt.Sprintf("UPDATE users SET %s WHERE id=%d", updateStmt, e.ID)
	fmt.Println(stmt)
	_, err = rp.conn.Exec(ctxSpan, stmt, values...)
	if err != nil {
		rp.errLogWithSpanAttributes("err while updating ", err, span)
	}
	return err
}

func (rp *postgresRepo) Delete(ctx context.Context, ID uint) error {
	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()
	stmt := `DELETE FROM users WHERE id=$1`
	_, err := rp.conn.Exec(ctxSpan, stmt, ID)
	if err != nil {
		rp.errLogWithSpanAttributes("err while deleting users", err, span)
	}
	return err
}

func (rp *postgresRepo) Get(ctx context.Context, where map[string]WhereClause) (Entity, error) {

	ctxSpan, span := rp.trace.Start(ctx, "create-repo-func")
	defer span.End()

	var result Entity

	// var res map[string]interface{}
	stmt, val := rp.buildGetStmt(where)
	err := rp.conn.QueryRow(ctxSpan, stmt, val...).Scan(&result.ID, &result.Name, &result.Email,
		&result.PhoneNo, &result.HomeAddress.AdressLine, &result.HomeAddress.City,
		&result.HomeAddress.PhoneNo, &result.HomeAddress.PinCode, &result.HomeAddress.Landmark,
		&result.IsAdmin)

	// if err != nil {
	// 	rp.log.Errorf("err while Scaning: %s", err.Error())
	// }
	if err != nil {
		rp.errLogWithSpanAttributes("err while Getting Users", err, span)
	}

	return result, err
}

func (rp *postgresRepo) buildGetStmt(where map[string]WhereClause) (string, []interface{}) {
	sb := strings.Builder{}
	whereLen := len(where)
	val := make([]interface{}, 0, whereLen)

	sb.WriteString("FROM users ")

	if whereLen > 0 {
		sb.WriteString("WHERE ")
		idx := 1

		for k, v := range where {
			if v.Condition == "" {
				v.Condition = "="
			}

			stmt := fmt.Sprintf("%s %s $%v AND ", k, v.Condition, idx)
			if idx == whereLen {
				stmt = fmt.Sprintf("%s %s $%v", k, v.Condition, idx)
			}

			sb.WriteString(stmt)
			val = append(val, v.Value)
			idx++
		}
	}

	stmt := fmt.Sprintf("SELECT * %s", sb.String())

	return stmt, val
}

func (rp *postgresRepo) errLogWithSpanAttributes(msg string, err error, span trace.Span) {
	// mark span with the error
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())

	// extracting traceID for logging purpose
	traceID := span.SpanContext().TraceID().String()
	rp.log.Errorf(msg+"err: %v", err, zap.String("traceID", traceID))
}
