package storage

import (
	"fmt"
	"strings"

	"github.com/angelbachev/go-money-api/models"
	_ "github.com/go-sql-driver/mysql"
)

type ExpenseStore interface {
	CreateExpense(cateory *models.Expense) error
	GetExpenses(userID, accountID int64, filters *models.ExpenseFilters, page, limit int64) ([]*models.Expense, error)
	GetExpensesCount(userID, accountID int64, filters *models.ExpenseFilters) (int64, error)
	GetExpenseByID(userID, accountID, expenseID int64) (*models.Expense, error)
	DeleteExpense(id int64) error
	UpdateExpense(expense *models.Expense) error
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

func (s MySQLStore) GetExpenses(userID, accountID int64, filters *models.ExpenseFilters, page, limit int64) ([]*models.Expense, error) {
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
		%s
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

	var limitClause string
	if page > 0 && limit > 0 {
		limitClause = "LIMIT ?, ?"
		offset := (page - 1) * limit
		params = append(params, offset)
		params = append(params, limit)
	}

	query = fmt.Sprintf(query, filterClause, limitClause)
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

func (s MySQLStore) GetExpensesCount(userID, accountID int64, filters *models.ExpenseFilters) (int64, error) {
	query := `
		SELECT COUNT(id) 
		FROM expenses 
		WHERE 
			user_id = ? 
			AND account_id = ?
			%s
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
	row := s.db.QueryRow(query, params...)
	var count int64
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s MySQLStore) GetExpenseByID(userID, accountID, expenseID int64) (*models.Expense, error) {
	query := `
	SELECT * 
	FROM expenses 
	WHERE 
		user_id = ? 
		AND account_id = ?
		AND id = ?
	LIMIT 1
`
	row := s.db.QueryRow(query, userID, accountID, expenseID)
	var expense models.Expense
	err := row.Scan(
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

	return &expense, nil
}

func (s MySQLStore) DeleteExpense(id int64) error {
	query := "DELETE FROM expenses WHERE id = ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s MySQLStore) UpdateExpense(expense *models.Expense) error {
	query := "UPDATE expenses SET category_id = ?, description = ?, amount = ?, date = ?, updated_at = ? WHERE id = ?"
	_, err := s.db.Exec(
		query,
		expense.CategoryID,
		expense.Description,
		expense.Amount,
		expense.Date,
		expense.UpdatedAt,
		expense.ID,
	)

	return err
}
