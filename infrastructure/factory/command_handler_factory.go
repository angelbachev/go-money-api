package factory

import (
	"github.com/angelbachev/go-money-api/application/command/account/create_account"
	"github.com/angelbachev/go-money-api/application/command/account/delete_account"
	"github.com/angelbachev/go-money-api/application/command/account/update_account"
	"github.com/angelbachev/go-money-api/application/command/auth/create_auth_token"
	"github.com/angelbachev/go-money-api/application/command/category/create_category"
	"github.com/angelbachev/go-money-api/application/command/category/delete_category"
	"github.com/angelbachev/go-money-api/application/command/category/update_category"
	"github.com/angelbachev/go-money-api/application/command/expense/create_expense"
	"github.com/angelbachev/go-money-api/application/command/expense/delete_expense"
	"github.com/angelbachev/go-money-api/application/command/expense/import_expenses"
	"github.com/angelbachev/go-money-api/application/command/expense/update_expense"
	"github.com/angelbachev/go-money-api/application/command/user/create_user"
	"github.com/angelbachev/go-money-api/application/command/user/update_user_settings"
	"github.com/angelbachev/go-money-api/domain/auth"
)

type CommandHandlerFactory struct {
	handlers            map[string]interface{}
	dbRepositoryFactory *MySQLRepositoryFactory
	authService         auth.AuthService
}

func NewCommandHandlerFactory(dbRepositoryFactory *MySQLRepositoryFactory, authService auth.AuthService) *CommandHandlerFactory {
	return &CommandHandlerFactory{handlers: make(map[string]interface{}), dbRepositoryFactory: dbRepositoryFactory, authService: authService}
}

func (f CommandHandlerFactory) CreateUserCommandHandler() *create_user.CreateUserCommandHandler {
	_, ok := f.handlers["createUserCommandHandler"]
	if !ok {
		f.handlers["createUserCommandHandler"] = create_user.NewCreateUserCommandHandler(
			f.dbRepositoryFactory.UserRepository(),
			f.dbRepositoryFactory.UserSettingsRepository(),
			f.dbRepositoryFactory.AccountRepository(),
			f.dbRepositoryFactory.CategoryRepository(),
		)
	}

	return f.handlers["createUserCommandHandler"].(*create_user.CreateUserCommandHandler)
}

func (f CommandHandlerFactory) UpdateUserSettingsCommandHandler() *update_user_settings.UpdateUserSettingsCommandHandler {
	_, ok := f.handlers["updateUserSettingsCommandHandler"]
	if !ok {
		f.handlers["updateUserSettingsCommandHandler"] = update_user_settings.NewUpdateUserSettingsCommandHandler(
			f.dbRepositoryFactory.UserSettingsRepository(),
		)
	}

	return f.handlers["updateUserSettingsCommandHandler"].(*update_user_settings.UpdateUserSettingsCommandHandler)
}

func (f CommandHandlerFactory) CreateAuthTokenCommandHandler() *create_auth_token.CreateAuthTokenCommandHandler {
	_, ok := f.handlers["createAuthTokenCommandHandler"]
	if !ok {
		f.handlers["createAuthTokenCommandHandler"] = create_auth_token.NewCreateAuthTokenCommandHandler(
			f.dbRepositoryFactory.UserRepository(),
			f.dbRepositoryFactory.UserSettingsRepository(),
			f.authService,
		)
	}

	return f.handlers["createAuthTokenCommandHandler"].(*create_auth_token.CreateAuthTokenCommandHandler)
}

func (f CommandHandlerFactory) CreateAccountCommandHandler() *create_account.CreateAccountCommandHandler {
	_, ok := f.handlers["createAccountCommandHandler"]
	if !ok {
		f.handlers["createAccountCommandHandler"] = create_account.NewCreateAccountCommandHandler(
			f.dbRepositoryFactory.AccountRepository(),
		)
	}

	return f.handlers["createAccountCommandHandler"].(*create_account.CreateAccountCommandHandler)
}

func (f CommandHandlerFactory) UpdateAccountCommandHandler() *update_account.UpdateAccountCommandHandler {
	_, ok := f.handlers["updateAccountCommandHandler"]
	if !ok {
		f.handlers["updateAccountCommandHandler"] = update_account.NewUpdateAccountCommandHandler(
			f.dbRepositoryFactory.AccountRepository(),
		)
	}

	return f.handlers["updateAccountCommandHandler"].(*update_account.UpdateAccountCommandHandler)
}

func (f CommandHandlerFactory) DeleteAccountCommandHandler() *delete_account.DeleteAccountCommandHandler {
	_, ok := f.handlers["deleteAccountCommandHandler"]
	if !ok {
		f.handlers["deleteAccountCommandHandler"] = delete_account.NewDeleteAccountCommandHandler(
			f.dbRepositoryFactory.AccountRepository(),
			f.dbRepositoryFactory.CategoryRepository(),
			f.dbRepositoryFactory.ExpenseRepository(),
		)
	}

	return f.handlers["deleteAccountCommandHandler"].(*delete_account.DeleteAccountCommandHandler)
}

func (f CommandHandlerFactory) CreateCategoryCommandHandler() *create_category.CreateCategoryCommandHandler {
	_, ok := f.handlers["createCategoryCommandHandler"]
	if !ok {
		f.handlers["createCategoryCommandHandler"] = create_category.NewCreateCategoryCommandHandler(
			f.dbRepositoryFactory.CategoryRepository(),
		)
	}

	return f.handlers["createCategoryCommandHandler"].(*create_category.CreateCategoryCommandHandler)
}

func (f CommandHandlerFactory) UpdateCategoryCommandHandler() *update_category.UpdateCategoryCommandHandler {
	_, ok := f.handlers["updateCategoryCommandHandler"]
	if !ok {
		f.handlers["updateCategoryCommandHandler"] = update_category.NewUpdateCategoryCommandHandler(
			f.dbRepositoryFactory.CategoryRepository(),
		)
	}

	return f.handlers["updateCategoryCommandHandler"].(*update_category.UpdateCategoryCommandHandler)
}

func (f CommandHandlerFactory) DeleteCategoryCommandHandler() *delete_category.DeleteCategoryCommandHandler {
	_, ok := f.handlers["deleteCategoryCommandHandler"]
	if !ok {
		f.handlers["deleteCategoryCommandHandler"] = delete_category.NewDeleteCategoryCommandHandler(
			f.dbRepositoryFactory.CategoryRepository(),
			f.dbRepositoryFactory.ExpenseRepository(),
		)
	}

	return f.handlers["deleteCategoryCommandHandler"].(*delete_category.DeleteCategoryCommandHandler)
}

func (f CommandHandlerFactory) CreateExpenseCommandHandler() *create_expense.CreateExpenseCommandHandler {
	_, ok := f.handlers["createExpenseCommandHandler"]
	if !ok {
		f.handlers["createExpenseCommandHandler"] = create_expense.NewCreateExpenseCommandHandler(
			f.dbRepositoryFactory.ExpenseRepository(),
		)
	}

	return f.handlers["createExpenseCommandHandler"].(*create_expense.CreateExpenseCommandHandler)
}

func (f CommandHandlerFactory) UpdateExpenseCommandHandler() *update_expense.UpdateExpenseCommandHandler {
	_, ok := f.handlers["updateExpenseCommandHandler"]
	if !ok {
		f.handlers["updateExpenseCommandHandler"] = update_expense.NewUpdateExpenseCommandHandler(
			f.dbRepositoryFactory.ExpenseRepository(),
		)
	}

	return f.handlers["updateExpenseCommandHandler"].(*update_expense.UpdateExpenseCommandHandler)
}

func (f CommandHandlerFactory) DeleteExpenseCommandHandler() *delete_expense.DeleteExpenseCommandHandler {
	_, ok := f.handlers["deleteExpenseCommandHandler"]
	if !ok {
		f.handlers["deleteExpenseCommandHandler"] = delete_expense.NewDeleteExpenseCommandHandler(
			f.dbRepositoryFactory.ExpenseRepository(),
		)
	}

	return f.handlers["deleteExpenseCommandHandler"].(*delete_expense.DeleteExpenseCommandHandler)
}

func (f CommandHandlerFactory) ImportExpensesCommandHandler() *import_expenses.ImportExpensesCommandHandler {
	_, ok := f.handlers["importExpensesCommandHandler"]
	if !ok {
		f.handlers["importExpensesCommandHandler"] = import_expenses.NewImportExpensesCommandHandler(
			f.dbRepositoryFactory.CategoryRepository(),
			f.dbRepositoryFactory.ExpenseRepository(),
		)
	}

	return f.handlers["importExpensesCommandHandler"].(*import_expenses.ImportExpensesCommandHandler)
}
