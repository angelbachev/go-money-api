package expense

type ExpenseRepositoryInterface interface {
	CreateExpense(cateory *Expense) error
	GetExpenses(userID, accountID int64, filters *ExpenseFilters, page, limit int64) ([]*Expense, error)
	GetExpensesCount(userID, accountID int64, filters *ExpenseFilters) (int64, error)
	GetExpenseByID(userID, accountID, expenseID int64) (*Expense, error)
	DeleteExpense(id int64) error
	UpdateExpense(expense *Expense) error
}
