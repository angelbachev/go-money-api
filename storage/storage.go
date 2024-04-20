package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store interface {
	UserStore
	BudgetStore
	CategoryStore
	ExpenseStore
}

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(connectionString string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLStore{db: db}, nil
}
