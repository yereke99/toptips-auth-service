package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// PgxIface is pgx interface
type PgxIface interface {
	// using pgxconn interface
	// Begin(context.Context) (pgx.Tx, error)
	// Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	// QueryRow(context.Context, string, ...interface{}) pgx.Row
	// Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	// Ping(context.Context) error
	// Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
	// Close(context.Context) error

	// using pgxpool interface
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
	Close()
}

type Database struct {
	DB PgxIface
}

func NewDatabase(db PgxIface) Database {
	return Database{DB: db}
}

func (pool Database) Create(contact string, code int) error {
	q := `INSERT INTO sms_cache(contact, code) VALUES($1,$2)`

	_, err := pool.DB.Exec(context.Background(), q, contact, code)

	if err != nil {
		return err
	}

	return nil
}

// Check
func (pool Database) CheckFromRepo(contact, role string) (bool, error) {
	switch role {
	case "driver":
		q := `Select phone From driver WHERE phone=$1`
		row := pool.DB.QueryRow(context.Background(), q, contact)

		var phone string

		err := row.Scan(&phone)
		if err != nil {
			switch err.Error() {
			case "no rows in result set":
				return true, nil
			default:
				return false, err
			}
		}

		if phone == contact {
			return false, nil
		}

		return true, nil

	case "user":
		q := `Select phone From customer WHERE phone=$1`
		row := pool.DB.QueryRow(context.Background(), q, contact)

		var phone string

		err := row.Scan(&phone)
		fmt.Println(err)
		if err != nil {
			switch err.Error() {
			case "no rows in result set":
				return true, nil
			default:
				return false, err
			}
		}

		if phone == contact {
			return false, nil
		}

		return true, nil
	default:
		return false, errors.New("Wrong role is given.")
	}

}

func (pool Database) GiveToken(contact, role string) (string, error) {
	switch role {
	case "driver":
		q := `Select token From driver Where phone=$1`

		row := pool.DB.QueryRow(context.Background(), q, contact)

		var token string

		err := row.Scan(&token)

		if err != nil {
			return "", err
		}

		return token, nil
	case "user":
		q := `Select token From users Where contact=$1`

		row := pool.DB.QueryRow(context.Background(), q, contact)

		var token string

		err := row.Scan(&token)

		if err != nil {
			return "", err
		}

		return token, nil
	default:
		return "", errors.New("Wrong type of role.")
	}
}

func (pool Database) ValidateSMS(contact string) (int, error) {
	q := `Select code From sms_cache WHERE contact=$1`

	row := pool.DB.QueryRow(context.Background(), q, contact)

	var code int

	err := row.Scan(&code)

	if err != nil {
		return 0, err
	}
	return code, nil
}

func (pool Database) Login(contact string) (string, error) {
	var password string
	q := `Select password From auth WHERE contact=$1`
	row := pool.DB.QueryRow(context.Background(), q, contact)

	if err := row.Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}

func (pool Database) Clean(contact string) error {
	q := `DELETE FROM sms_cache x USING sms_cache y WHERE x.id <= y.id AND x.contact = y.contact AND x.contact = $1`

	_, err := pool.DB.Exec(context.Background(), q, contact)

	if err != nil {
		return err
	}
	return nil
}
