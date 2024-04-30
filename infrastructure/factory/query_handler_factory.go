package factory

import (
	"github.com/angelbachev/go-money-api/application/query/account/list_accounts"
	"github.com/angelbachev/go-money-api/application/query/category/get_category"
	"github.com/angelbachev/go-money-api/application/query/category/list_categories"
	"github.com/angelbachev/go-money-api/application/query/category/list_category_icons"
	"github.com/angelbachev/go-money-api/application/query/expense/list_expenses"
	"github.com/angelbachev/go-money-api/application/query/user/get_user_settings"
)

type QueryHandlerFactory struct {
	handlers            map[string]interface{}
	dbRepositoryFactory *MySQLRepositoryFactory
}

func NewQueryHandlerFactory(dbRepositoryFactory *MySQLRepositoryFactory) *QueryHandlerFactory {
	return &QueryHandlerFactory{handlers: make(map[string]interface{}), dbRepositoryFactory: dbRepositoryFactory}
}

func (f QueryHandlerFactory) ListAccountsQueryHandler() *list_accounts.ListAccountsQueryHandler {
	_, ok := f.handlers["listAccountsQueryHandler"]
	if !ok {
		f.handlers["listAccountsQueryHandler"] = list_accounts.NewListAccountsQueryHandler(
			f.dbRepositoryFactory.AccountRepository(),
		)
	}

	return f.handlers["listAccountsQueryHandler"].(*list_accounts.ListAccountsQueryHandler)
}

func (f QueryHandlerFactory) ListCategoriesQueryHandler() *list_categories.ListCategoriesQueryHandler {
	_, ok := f.handlers["updateUserSettingsCommandHandler"]
	if !ok {
		f.handlers["updateUserSettingsCommandHandler"] = list_categories.NewListCategoriesQueryHandler(
			f.dbRepositoryFactory.CategoryRepository(),
		)
	}

	return f.handlers["updateUserSettingsCommandHandler"].(*list_categories.ListCategoriesQueryHandler)
}

func (f QueryHandlerFactory) ListCategoryIconsQueryHandler() *list_category_icons.ListCategoryIconsQueryHandler {
	_, ok := f.handlers["listCategoryIconsQueryHandler"]
	if !ok {
		f.handlers["listCategoryIconsQueryHandler"] = list_category_icons.NewListCategoryIconsQueryHandler()
	}

	return f.handlers["listCategoryIconsQueryHandler"].(*list_category_icons.ListCategoryIconsQueryHandler)
}

func (f QueryHandlerFactory) GetCategoryQueryHandler() *get_category.GetCategoryQueryHandler {
	_, ok := f.handlers["getCategoryQueryHandler"]
	if !ok {
		f.handlers["getCategoryQueryHandler"] = get_category.NewGetCategoryQueryHandler(
			f.dbRepositoryFactory.CategoryRepository(),
		)
	}

	return f.handlers["getCategoryQueryHandler"].(*get_category.GetCategoryQueryHandler)
}

func (f QueryHandlerFactory) GetUserSettingsQueryHandler() *get_user_settings.GetUserSettingsQueryHandler {
	_, ok := f.handlers["getUserSettingsQueryHandler"]
	if !ok {
		f.handlers["getUserSettingsQueryHandler"] = get_user_settings.NewGetUserSettingsQueryHandler(
			f.dbRepositoryFactory.UserSettingsRepository(),
		)
	}

	return f.handlers["getUserSettingsQueryHandler"].(*get_user_settings.GetUserSettingsQueryHandler)
}

func (f QueryHandlerFactory) ListExpensesQueryHandler() *list_expenses.ListExpensesQueryHandler {
	_, ok := f.handlers["listExpensesQueryHandler"]
	if !ok {
		f.handlers["listExpensesQueryHandler"] = list_expenses.NewListExpensesQueryHandler(
			f.dbRepositoryFactory.CategoryRepository(),
			f.dbRepositoryFactory.ExpenseRepository(),
		)
	}

	return f.handlers["listExpensesQueryHandler"].(*list_expenses.ListExpensesQueryHandler)
}
