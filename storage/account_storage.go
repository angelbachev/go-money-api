package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/angelbachev/go-money-api/models"
)

type AccountStore interface {
	CreateAccount(account *models.Account) error
	GetAccountByID(userID, accountID int64) (*models.Account, error)
	GetAccounts(userID int64) ([]*models.Account, error)
}

func (s MySQLStore) CreateAccount(account *models.Account) error {
	query := `
	INSERT INTO accounts (user_id, name, description, currency_code, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		account.UserID,
		account.Name,
		account.Description,
		account.CurrencyCode,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	account.ID = id

	return nil
}

func (s MySQLStore) GetAccountByID(userID, id int64) (*models.Account, error) {
	query := `SELECT * FROM accounts WHERE id = ? AND user_id = ? LIMIT 1`
	row := s.db.QueryRow(query, id, userID)

	var account models.Account
	switch err := row.Scan(
		&account.ID,
		&account.UserID,
		&account.Name,
		&account.Description,
		&account.CurrencyCode,
		&account.CreatedAt,
		&account.UpdatedAt,
	); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &account, nil
	default:
		return nil, err
	}
}

func (s MySQLStore) GetAccounts(userID int64) ([]*models.Account, error) {
	query := `SELECT * FROM accounts WHERE user_id = ? ORDER BY name ASC`
	rows, err := s.db.Query(query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []*models.Account{}
	for rows.Next() {
		var account models.Account
		err = rows.Scan(
			&account.ID,
			&account.UserID,
			&account.Name,
			&account.Description,
			&account.CurrencyCode,
			&account.CreatedAt,
			&account.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}
	return accounts, nil
}
