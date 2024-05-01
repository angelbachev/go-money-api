package factory

import (
	a "github.com/angelbachev/go-money-api/domain/account"
	c "github.com/angelbachev/go-money-api/domain/category"
	e "github.com/angelbachev/go-money-api/domain/expense"
	u "github.com/angelbachev/go-money-api/domain/user"

	"github.com/angelbachev/go-money-api/infrastructure/domain"
	"github.com/angelbachev/go-money-api/infrastructure/domain/account"
	"github.com/angelbachev/go-money-api/infrastructure/domain/category"
	"github.com/angelbachev/go-money-api/infrastructure/domain/expense"
	"github.com/angelbachev/go-money-api/infrastructure/domain/user"
)

type MySQLRepositoryFactory struct {
	repositories   map[string]interface{}
	baseRepository *domain.MySQLBaseRepository
}

func NewMySQLRepositoryFactory(connectionString string) (*MySQLRepositoryFactory, error) {
	baseRepository, err := domain.NewMySQLBaseRepository(connectionString)
	if err != nil {
		return nil, err
	}

	return &MySQLRepositoryFactory{repositories: make(map[string]interface{}), baseRepository: baseRepository}, nil
}

func (f MySQLRepositoryFactory) UserRepository() u.UserRepository {
	_, ok := f.repositories["userRepository"]
	if !ok {
		f.repositories["userRepository"] = user.MySQLUserRepository{MySQLBaseRepository: f.baseRepository}
	}

	return f.repositories["userRepository"].(u.UserRepository)
}

func (f MySQLRepositoryFactory) UserSettingsRepository() u.UserSettingsRepository {
	_, ok := f.repositories["userSettingsRepository"]
	if !ok {
		f.repositories["userSettingsRepository"] = user.MySQLUserSettingsRepository{MySQLBaseRepository: f.baseRepository}
	}

	return f.repositories["userSettingsRepository"].(u.UserSettingsRepository)
}

func (f MySQLRepositoryFactory) AccountRepository() a.AccountRepository {
	_, ok := f.repositories["accountRepository"]
	if !ok {
		f.repositories["accountRepository"] = account.MySQLAccountRepository{MySQLBaseRepository: f.baseRepository}
	}

	return f.repositories["accountRepository"].(a.AccountRepository)
}

func (f MySQLRepositoryFactory) CategoryRepository() c.CategoryRepository {
	_, ok := f.repositories["categoryRepository"]
	if !ok {
		f.repositories["categoryRepository"] = category.MySQLCategoryRepository{MySQLBaseRepository: f.baseRepository}
	}

	return f.repositories["categoryRepository"].(c.CategoryRepository)
}

func (f MySQLRepositoryFactory) ExpenseRepository() e.ExpenseRepository {
	_, ok := f.repositories["expenseRepository"]
	if !ok {
		f.repositories["expenseRepository"] = expense.MySQLExpenseRepository{MySQLBaseRepository: f.baseRepository}
	}

	return f.repositories["expenseRepository"].(e.ExpenseRepository)
}
