package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/angelbachev/go-money-api/models"
)

type BudgetStore interface {
	CreateBudget(budget *models.Budget) error
	GetBudgetByID(userID, budgetID int64) (*models.Budget, error)
}

func (s MySQLStore) CreateBudget(budget *models.Budget) error {
	query := `
	INSERT INTO budgets (user_id, name, description, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		budget.UserID,
		budget.Name,
		budget.Description,
		budget.CreatedAt,
		budget.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	budget.ID = id

	return nil
}

func (s MySQLStore) GetBudgetByID(userID, id int64) (*models.Budget, error) {
	query := `SELECT * FROM budgets WHERE id = ? AND user_id = ? LIMIT 1`
	row := s.db.QueryRow(query, id, userID)

	var budget models.Budget
	switch err := row.Scan(
		&budget.ID,
		&budget.UserID,
		&budget.Name,
		&budget.Description,
		&budget.CreatedAt,
		&budget.UpdatedAt,
	); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &budget, nil
	default:
		return nil, err
	}
}
