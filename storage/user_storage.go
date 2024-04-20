package storage

import (
	"database/sql"

	"github.com/angelbachev/go-money-api/models"
	_ "github.com/go-sql-driver/mysql"
)

type UserStore interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int64) (*models.User, error)
}

func (s MySQLStore) CreateUser(user *models.User) error {
	query := `
	INSERT INTO users (email, password, created_at, updated_at)
	VALUES (?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func (s MySQLStore) GetUserByID(id int64) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = ? LIMIT 1`
	row := s.db.QueryRow(query, id)

	return scanIntoUser(row)
}

func (s MySQLStore) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = ? LIMIT 1`
	row := s.db.QueryRow(query, email)

	return scanIntoUser(row)
}

func scanIntoUser(row *sql.Row) (*models.User, error) {
	var user models.User
	switch err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}
