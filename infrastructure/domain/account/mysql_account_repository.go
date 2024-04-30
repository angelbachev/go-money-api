package account

import (
	"database/sql"

	"github.com/angelbachev/go-money-api/domain/account"
	"github.com/angelbachev/go-money-api/infrastructure/domain"
)

type MySQLAccountRepository struct {
	*domain.MySQLBaseRepository
}

func (r MySQLAccountRepository) CreateAccount(account *account.Account) error {
	query := `
	INSERT INTO accounts (user_id, name, description, currency_code, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := r.DB.Exec(
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

func (r MySQLAccountRepository) UpdateAccount(account *account.Account) error {
	query := "UPDATE accounts SET name = ?, description = ?, currency_code = ?, updated_at = ? WHERE id = ?"
	_, err := r.DB.Exec(
		query,
		account.Name,
		account.Description,
		account.CurrencyCode,
		account.UpdatedAt,
		account.ID,
	)

	return err
}

func (r MySQLAccountRepository) GetAccountByID(userID, id int64) (*account.Account, error) {
	query := `SELECT * FROM accounts WHERE id = ? AND user_id = ? LIMIT 1`
	row := r.DB.QueryRow(query, id, userID)

	var account account.Account
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

func (r MySQLAccountRepository) GetAccounts(userID int64) ([]*account.Account, error) {
	query := `SELECT * FROM accounts WHERE user_id = ? ORDER BY name ASC`
	rows, err := r.DB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []*account.Account{}
	for rows.Next() {
		var account account.Account
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

func (r MySQLAccountRepository) DeleteAccount(id int64) error {
	query := "DELETE FROM accounts WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
