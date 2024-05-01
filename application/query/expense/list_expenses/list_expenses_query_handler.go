package list_expenses

import (
	"github.com/angelbachev/go-money-api/domain/category"
	"github.com/angelbachev/go-money-api/domain/expense"
)

type ListExpensesQueryHandler struct {
	categoryRepository category.CategoryRepository
	expenseRepository  expense.ExpenseRepository
}

func NewListExpensesQueryHandler(
	categoryRepository category.CategoryRepository,
	expenseRepository expense.ExpenseRepository,
) *ListExpensesQueryHandler {
	return &ListExpensesQueryHandler{
		categoryRepository: categoryRepository,
		expenseRepository:  expenseRepository,
	}
}

func (h ListExpensesQueryHandler) Handle(query ListExpensesQuery) (*ListExpensesResponse, error) {
	var categoryIDs []int64
	if len(query.CategoryIDs) > 0 {
		catIDS, err := h.categoryRepository.GetListCategoryIDsAndTheirSubcategories(query.CategoryIDs)
		if err != nil {
			return nil, err
		}
		categoryIDs = catIDS
	}

	filters := &expense.ExpenseFilters{
		MinAmount:   query.MinAmount,
		MaxAmount:   query.MaxAmount,
		MinDate:     query.MinDate,
		MaxDate:     query.MaxDate,
		CategoryIDs: categoryIDs,
	}

	// TODO: handle subcategories
	// TODO: validate user owns the account
	expenses, err := h.expenseRepository.GetExpenses(query.UserID, query.AccountID, filters, query.Page, query.Limit)
	if err != nil {
		return nil, err
	}

	// TODO: refactor to optimize queries (reduce their number)
	var formattedExpenses []*ExpenseResponse
	for _, exp := range expenses {
		c, err := h.categoryRepository.GetCategoryByID(exp.CategoryID)
		if err != nil {
			return nil, err
		}
		formattedExpenses = append(formattedExpenses, NewExpenseResponse(exp, c.Name))
	}

	totalCount, err := h.expenseRepository.GetExpensesCount(query.UserID, query.AccountID, filters)
	if err != nil {
		return nil, err
	}

	return NewListExpensesResponse(formattedExpenses, totalCount), nil
}
