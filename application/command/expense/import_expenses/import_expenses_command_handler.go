package import_expenses

import (
	"encoding/csv"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/angelbachev/go-money-api/domain/category"
	"github.com/angelbachev/go-money-api/domain/expense"
)

type ImportExpensesCommandHandler struct {
	categoryRepository category.CategoryRepositoryInterface
	expenseRepository  expense.ExpenseRepositoryInterface
}

func NewImportExpensesCommandHandler(
	categoryRepository category.CategoryRepositoryInterface,
	expenseRepository expense.ExpenseRepositoryInterface,
) *ImportExpensesCommandHandler {
	return &ImportExpensesCommandHandler{
		categoryRepository: categoryRepository,
		expenseRepository:  expenseRepository,
	}
}

func (h ImportExpensesCommandHandler) Handle(command ImportExpensesCommand) error {
	reader := csv.NewReader(command.File)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	existingCategories, err := h.categoryRepository.GetCategoryNames(command.AccountID)
	if err != nil {
		return err
	}

	rootCategoryID, err := h.categoryRepository.GetRootCategoryID(command.AccountID)
	if err != nil {
		return err
	}

	// Print the CSV data
	for _, row := range data {
		if row[0] == "" || row[2] == "" {
			continue // empty line
		}
		date, err := time.Parse("02.01.2006", row[0])
		if err != nil {
			// TODO: handle error
			continue
		}

		re := regexp.MustCompile("[0-9]+.?,?[0-9]+")
		amount, err := strconv.ParseFloat(strings.Replace(re.FindString(row[2]), ",", ".", 1), 64)
		if err != nil {
			// TODO: handle error
			continue
		}
		categoryName := strings.TrimSpace(row[1])
		categoryID, ok := existingCategories[strings.ToLower(categoryName)]
		if !ok {
			category := category.NewCategory(command.UserID, command.AccountID, rootCategoryID, categoryName, "", "")
			err = h.categoryRepository.CreateCategory(category)
			if err != nil {
				// TODO: handle error
				continue
			}
			categoryID = category.ID
			existingCategories[strings.ToLower(categoryName)] = categoryID
		}

		expense := expense.NewExpense(command.UserID, command.AccountID, categoryID, strings.TrimSpace(row[3]), int64(amount*100), date)
		h.expenseRepository.CreateExpense(expense)
	}

	// TODO: prevent entering the same expense multiple times
	// TODO: return error lines
	// TODO: add support for different currencies

	return nil
}
