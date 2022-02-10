package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Repository interface {
	Create(ctx context.Context, e Entity) (uint, error)
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, ID uint) (Entity, error)
}

type postgresRepo struct {
	log  *zap.SugaredLogger
	conn *pgx.Conn
}

func NewPostgresRepo(l *zap.SugaredLogger, c *pgx.Conn) Repository {
	return &postgresRepo{
		log:  l,
		conn: c,
	}
}

func (rp *postgresRepo) Create(ctx context.Context, e Entity) (uint, error) {
	stmt := `INSERT INTO users (
		name,email,phone_no,
		homeaddress_phoneno,homeaddress_address_line,
		homeaddress_city,homeaddress_pin_code,
		homeaddress_landmark,is_admin) VALUES 
		($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id
		`
	var insertedId uint
	err := rp.conn.QueryRow(ctx, stmt, e.Name, e.Email, e.PhoneNo,
		e.HomeAddress.PhoneNo, e.HomeAddress.AdressLine,
		e.HomeAddress.City, e.HomeAddress.PinCode,
		e.HomeAddress.Landmark, e.IsAdmin).Scan(&insertedId)

	return insertedId, err
}

func (rp *postgresRepo) Update(ctx context.Context, e Entity) error {
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
	_, err = rp.conn.Exec(ctx, stmt, values...)

	return err
}

func (rp *postgresRepo) Delete(ctx context.Context, ID uint) error {
	stmt := `DELETE FROM users WHERE id=$1`
	_, err := rp.conn.Exec(ctx, stmt, ID)
	return err
}

func (rp *postgresRepo) Get(ctx context.Context,
	sel map[string]interface{}, where map[string]interface{}) (Entity, error) {
	var result Entity
	sb := strings.Builder{}

	if len(where) > 0 {
		sb.WriteString("WHERE ")
	}

	idx := 1
	for i := range where {
		sb.WriteString(fmt.Sprintf(+" $" + idx + " "))
		idx++
	}
	stmt := `SELECT id,name FROM users WHERE id=$1`
	err := rp.conn.QueryRow(ctx, stmt, ID).Scan(&result)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(result)
	return Entity{}, nil
}
