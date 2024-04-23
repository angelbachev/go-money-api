package storage

import (
	"fmt"
	"strings"

	"github.com/angelbachev/go-money-api/models"
	_ "github.com/go-sql-driver/mysql"
)

type ExpenseStore interface {
	CreateExpense(cateory *models.Expense) error
	GetExpenses(userID, accountID int64, filters *models.ExpenseFilters) ([]*models.Expense, error)
}

func (s MySQLStore) CreateExpense(expense *models.Expense) error {
	query := `
	INSERT INTO expenses (user_id, account_id, category_id, description, amount, date, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		expense.UserID,
		expense.AccountID,
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

func (s MySQLStore) GetExpenses(userID, accountID int64, filters *models.ExpenseFilters) ([]*models.Expense, error) {
	query := `
		SELECT * 
		FROM expenses 
		WHERE 
			user_id = ? 
			AND account_id = ?
			%s
		ORDER BY 
			date DESC, 
			amount DESC
	`

	var filtersParts []string
	var params = []any{userID, accountID}

	if filters != nil {
		if filters.MinAmount != nil {
			filtersParts = append(filtersParts, "amount >= ?")
			params = append(params, *filters.MinAmount)
		}

		if filters.MaxAmount != nil {
			filtersParts = append(filtersParts, "amount <= ?")
			params = append(params, *filters.MaxAmount)
		}

		if filters.MinDate != nil {
			filtersParts = append(filtersParts, "date >= ?")
			params = append(params, *filters.MinDate)
		}

		if filters.MaxDate != nil {
			filtersParts = append(filtersParts, "date <= ?")
			params = append(params, *filters.MaxDate)
		}

		if len(filters.CategoryIDs) > 0 {
			categoriesClause := "category_id IN ("

			for idx, ct := range filters.CategoryIDs {
				params = append(params, ct)
				if idx == 0 {
					categoriesClause += "?"
				} else {
					categoriesClause += ", ?"
				}
			}
			categoriesClause += ")"
			filtersParts = append(filtersParts, categoriesClause)
		}
	}

	var filterClause string
	if len(filtersParts) > 0 {
		filterClause = " AND " + strings.Join(filtersParts, " AND ")
	}
	query = fmt.Sprintf(query, filterClause)
	rows, err := s.db.Query(query, params...)

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
			&expense.AccountID,
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
