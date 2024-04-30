package user

import (
	"database/sql"

	"github.com/angelbachev/go-money-api/domain/user"
	"github.com/angelbachev/go-money-api/infrastructure/domain"
)

type MySQLUserRepository struct {
	*domain.MySQLBaseRepository
}

func (r MySQLUserRepository) CreateUser(user *user.User) error {
	query := `
	INSERT INTO users (email, password, created_at, updated_at)
	VALUES (?, ?, ?, ?)
	`
	result, err := r.DB.Exec(
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

func (r MySQLUserRepository) GetUserByID(id int64) (*user.User, error) {
	query := `SELECT * FROM users WHERE id = ? LIMIT 1`
	row := r.DB.QueryRow(query, id)

	return scanIntoUser(row)
}

func (r MySQLUserRepository) GetUserByEmail(email string) (*user.User, error) {
	query := `SELECT * FROM users WHERE email = ? LIMIT 1`
	row := r.DB.QueryRow(query, email)

	return scanIntoUser(row)
}

func scanIntoUser(row *sql.Row) (*user.User, error) {
	var user user.User
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
