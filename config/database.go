package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DB config
type DataBaseConfig struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBName   string
}

func (db DataBaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		db.Username,
		db.Password,
		db.Hostname,
		db.Port,
		db.DBName,
	)
}

var tables = []string{
	createTableSMS,
	profile,
}

type DataStore struct {
	dbPool *pgxpool.Pool
}

func NewDataStore(pool *pgxpool.Pool) DataStore {
	return DataStore{dbPool: pool}
}

func NewDBPool(dbConfig DataBaseConfig) (*pgxpool.Pool, error) {

	ctx, cancell := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancell()

	pool, err := pgxpool.Connect(ctx, dbConfig.DSN())

	if err != nil {
		return nil, errors.New("database connection error")
	}

	err = validateDBPool(pool)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
func InitTabeles(pool *pgxpool.Pool) error {
	for _, table := range tables {
		_, err := pool.Exec(context.Background(), table)

		if err != nil {
			return err
		}
		fmt.Println("Created table: ")
	}
	return nil
}

// validateDBPool will pings the database and logs the current user and database
func validateDBPool(pool *pgxpool.Pool) error {
	// tried to ping connection
	err := pool.Ping(context.Background())

	// return error if error found
	if err != nil {
		return errors.New("database connection error")
	}

	var (
		currentDatabase string
		currentUser     string
		dbVersion       string
	)

	// Lets try to get db system info
	sqlStatement := `select current_database(), current_user, version();`
	row := pool.QueryRow(context.Background(), sqlStatement)
	err = row.Scan(&currentDatabase, &currentUser, &dbVersion)

	switch {
	case err == sql.ErrNoRows:
		return errors.New("no rows were returned")
	case err != nil:
		return errors.New("database connection error")
	default:
		log.Printf("database version: %s\n", dbVersion)
		log.Printf("current database user: %s\n", currentUser)
		log.Printf("current database: %s\n", currentDatabase)
	}

	return nil
}
