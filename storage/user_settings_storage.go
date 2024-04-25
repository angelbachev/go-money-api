package storage

import (
	"database/sql"

	"github.com/angelbachev/go-money-api/models"
	_ "github.com/go-sql-driver/mysql"
)

type UserSettingsStore interface {
	CreateUserSettings(settings *models.UserSettings) error
	UpdateUserSettings(settings *models.UserSettings) error
	GetUserSettings(userID int64) (*models.UserSettings, error)
}

func (s MySQLStore) CreateUserSettings(settings *models.UserSettings) error {
	query := `
	INSERT INTO user_settings (user_id, default_account_id, theme, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		settings.UserID,
		settings.DefaultAccountID,
		settings.Theme,
		settings.CreatedAt,
		settings.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	settings.ID = id

	return nil
}

func (s MySQLStore) UpdateUserSettings(settings *models.UserSettings) error {
	query := "UPDATE user_settings SET default_account_id = ?, theme = ?, updated_at = ? WHERE user_id = ?"
	_, err := s.db.Exec(
		query,
		settings.DefaultAccountID,
		settings.Theme,
		settings.UpdatedAt,
		settings.UserID,
	)
	return err
}

func (s MySQLStore) GetUserSettings(userID int64) (*models.UserSettings, error) {
	query := `SELECT * FROM user_settings WHERE user_id = ? LIMIT 1`
	row := s.db.QueryRow(query, userID)

	return scanIntoUserSettings(row)
}

func scanIntoUserSettings(row *sql.Row) (*models.UserSettings, error) {
	var settings models.UserSettings
	switch err := row.Scan(
		&settings.ID,
		&settings.UserID,
		&settings.DefaultAccountID,
		&settings.Theme,
		&settings.CreatedAt,
		&settings.UpdatedAt,
	); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &settings, nil
	default:
		return nil, err
	}
}
