package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	// postgresql driver
	_ "github.com/lib/pq"
)

type SQLOperations interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type DB interface {
	SQLOperations
	Begin() (*sql.Tx, error)
	Close() error
}

type RowScanner interface {
	Scan(dest ...interface{}) error
}

type AppDB struct {
	*sql.DB
}

func InitDB() *AppDB {
	return InitDBWithURL(
		os.Getenv("DATABASE_URL"),
	)
}

func InitDBWithURL(databaseURL string) *AppDB {
	if databaseURL == "" {
		panic("database url is required")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic(fmt.Sprintf("sql.Open failed because err=[%v]", err))
	}

	return &AppDB{
		DB: db,
	}
}

func (db *AppDB) InTransaction(ctx context.Context, operations func(context.Context, SQLOperations) error) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err = operations(ctx, tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}

		return err
	}

	return tx.Commit()
}
