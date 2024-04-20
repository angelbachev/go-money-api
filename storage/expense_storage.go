package storage

import (
	"github.com/angelbachev/go-money-api/models"
	_ "github.com/go-sql-driver/mysql"
)

type ExpenseStore interface {
	CreateExpense(cateory *models.Expense) error
	GetExpenses(userID, budgetID int64) ([]*models.Expense, error)
}

func (s MySQLStore) CreateExpense(expense *models.Expense) error {
	query := `
	INSERT INTO expenses (user_id, budget_id, category_id, description, amount, date, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		expense.UserID,
		expense.BudgetID,
		expense.CategoryID,
		expense.Description,
		expense.Amount,
		expense.Date,
		expense.CreatedAt,
		expense.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	expense.ID = id

	return nil
}

func (s MySQLStore) GetExpenses(userID, budgetID int64) ([]*models.Expense, error) {
	query := `SELECT * FROM expenses WHERE user_id = ? AND budget_id = ? ORDER BY date DESC, amount DESC`
	rows, err := s.db.Query(query, userID, budgetID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	expenses := []*models.Expense{}
	for rows.Next() {
		var expense models.Expense
		err = rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.BudgetID,
			&expense.CategoryID,
			&expense.Description,
			&expense.Amount,
			&expense.Date,
			&expense.CreatedAt,
			&expense.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		expenses = append(expenses, &expense)
	}
	return expenses, nil
}
