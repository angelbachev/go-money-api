package user

import (
	"database/sql"

	"github.com/angelbachev/go-money-api/domain/user"
	"github.com/angelbachev/go-money-api/infrastructure/domain"
)

type MySQLUserSettingsRepository struct {
	*domain.MySQLBaseRepository
}

func (r MySQLUserSettingsRepository) CreateUserSettings(settings *user.UserSettings) error {
	query := `
	INSERT INTO user_settings (user_id, default_account_id, theme, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.DB.Exec(
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

func (r MySQLUserSettingsRepository) UpdateUserSettings(settings *user.UserSettings) error {
	query := "UPDATE user_settings SET default_account_id = ?, theme = ?, updated_at = ? WHERE user_id = ?"
	_, err := r.DB.Exec(
		query,
		settings.DefaultAccountID,
		settings.Theme,
		settings.UpdatedAt,
		settings.UserID,
	)
	return err
}

func (r MySQLUserSettingsRepository) GetUserSettings(userID int64) (*user.UserSettings, error) {
	query := `SELECT * FROM user_settings WHERE user_id = ? LIMIT 1`
	row := r.DB.QueryRow(query, userID)

	return scanIntoUserSettings(row)
}

func scanIntoUserSettings(row *sql.Row) (*user.UserSettings, error) {
	var settings user.UserSettings
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
