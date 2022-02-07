package user

import (
	"context"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type Repository interface {
	Create(ctx context.Context, e Entity) error
	Update(ctx context.Context, e Entity) error
	Delete(ctx context.Context, ID uint) error
	Get(ctx context.Context, ID uint) error
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

func (rp *postgresRepo) Create(ctx context.Context, e Entity) error {
	stmt := `INSERT INTO users (
		name,email,phone_no,
		homeaddress_phoneno,homeaddress_address_line,
		homeaddress_city,homeaddress_pin_code,
		homeaddress_landmark,is_admin) VALUES 
		($1,$2,$3,$4,$5,$6,$7,$8,$9)
		`
	_, err := rp.conn.Exec(ctx, stmt, e.Name, e.Email, e.PhoneNo,
		e.HomeAddress.PhoneNo, e.HomeAddress.AdressLine,
		e.HomeAddress.City, e.HomeAddress.PinCode,
		e.HomeAddress.Landmark, e.IsAdmin)
	return err
}

func (rp *postgresRepo) Update(ctx context.Context, e Entity) error {
	return nil
}

func (rp *postgresRepo) Delete(ctx context.Context, ID uint) error {
	return nil
}

func (rp *postgresRepo) Get(ctx context.Context, ID uint) error {
	return nil
}
