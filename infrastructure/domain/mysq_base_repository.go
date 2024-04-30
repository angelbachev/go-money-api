package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLBaseRepository struct {
	DB *sql.DB
}

func NewMySQLBaseRepository(connectionString string) (*MySQLBaseRepository, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLBaseRepository{DB: db}, nil
}
