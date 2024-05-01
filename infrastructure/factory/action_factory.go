package factory

import (
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/angelbachev/go-money-api/presentation/api/rest/account"
	"github.com/angelbachev/go-money-api/presentation/api/rest/auth"
	"github.com/angelbachev/go-money-api/presentation/api/rest/category"
	"github.com/angelbachev/go-money-api/presentation/api/rest/expense"
	"github.com/angelbachev/go-money-api/presentation/api/rest/user"
)

type ActionFactory struct {
	actions               map[string]rest.APIAction
	commandHandlerFactory *CommandHandlerFactory
	queryHandlerFactory   *QueryHandlerFactory
}

func NewActionFactory(commandHandlerFactory *CommandHandlerFactory, queryHandlerFactory *QueryHandlerFactory) *ActionFactory {
	return &ActionFactory{actions: make(map[string]rest.APIAction), commandHandlerFactory: commandHandlerFactory, queryHandlerFactory: queryHandlerFactory}
}

func (f ActionFactory) CreateUserAction() *user.CreateUserAction {
	_, ok := f.actions["createUserAction"]
	if !ok {
		f.actions["createUserAction"] = user.NewCreateUserAction(
			f.commandHandlerFactory.CreateUserCommandHandler(),
		)
	}

	return f.actions["createUserAction"].(*user.CreateUserAction)
}

func (f ActionFactory) UpdateUserSettingsAction() *user.UpdateUserSettingsAction {
	_, ok := f.actions["updateUserSettingsAction"]
	if !ok {
		f.actions["updateUserSettingsAction"] = user.NewUpdateUserSettingsAction(
			f.commandHandlerFactory.UpdateUserSettingsCommandHandler(),
		)
	}
	return f.actions["updateUserSettingsAction"].(*user.UpdateUserSettingsAction)
}

func (f ActionFactory) GetUserSettingsAction() *user.GetUserSettingsAction {
	_, ok := f.actions["getUserSettingsAction"]
	if !ok {
		f.actions["getUserSettingsAction"] = user.NewGetUserSettingsAction(
			f.queryHandlerFactory.GetUserSettingsQueryHandler(),
		)
	}

	return f.actions["getUserSettingsAction"].(*user.GetUserSettingsAction)
}

func (f ActionFactory) CreateAuthTokenAction() *auth.CreateAuthTokenAction {
	_, ok := f.actions["createAuthTokenAction"]
	if !ok {
		f.actions["createAuthTokenAction"] = auth.NewCreateAuthTokenAction(
			f.commandHandlerFactory.CreateAuthTokenCommandHandler(),
		)
	}

	return f.actions["createAuthTokenAction"].(*auth.CreateAuthTokenAction)
}

func (f ActionFactory) CreateAccountAction() *account.CreateAccountAction {
	_, ok := f.actions["createAccountAction"]
	if !ok {
		f.actions["createAccountAction"] = account.NewCreateAccountAction(
			f.commandHandlerFactory.CreateAccountCommandHandler(),
		)
	}

	return f.actions["createAccountAction"].(*account.CreateAccountAction)
}

func (f ActionFactory) UpdateAccountAction() *account.UpdateAccountAction {
	_, ok := f.actions["updateAccountAction"]
	if !ok {
		f.actions["updateAccountAction"] = account.NewUpdateAccountAction(
			f.commandHandlerFactory.UpdateAccountCommandHandler(),
		)
	}

	return f.actions["updateAccountAction"].(*account.UpdateAccountAction)
}

func (f ActionFactory) DeleteAccountAction() *account.DeleteAccountAction {
	_, ok := f.actions["deleteAccountAction"]
	if !ok {
		f.actions["deleteAccountAction"] = account.NewDeleteAccountAction(
			f.commandHandlerFactory.DeleteAccountCommandHandler(),
		)
	}

	return f.actions["deleteAccountAction"].(*account.DeleteAccountAction)
}

func (f ActionFactory) ListAccountsAction() *account.ListAccountsAction {
	_, ok := f.actions["listAccountsAction"]
	if !ok {
		f.actions["listAccountsAction"] = account.NewListAccountsAction(
			f.queryHandlerFactory.ListAccountsQueryHandler(),
		)
	}

	return f.actions["listAccountsAction"].(*account.ListAccountsAction)
}

func (f ActionFactory) CreateCategoryAction() *category.CreateCategoryAction {
	_, ok := f.actions["createCategoryAction"]
	if !ok {
		f.actions["createCategoryAction"] = category.NewCreateCategoryAction(
			f.commandHandlerFactory.CreateCategoryCommandHandler(),
		)
	}

	return f.actions["createCategoryAction"].(*category.CreateCategoryAction)
}

func (f ActionFactory) UpdateCategoryAction() *category.UpdateCategoryAction {
	_, ok := f.actions["updateCategoryAction"]
	if !ok {
		f.actions["updateCategoryAction"] = category.NewUpdateCategoryAction(
			f.commandHandlerFactory.UpdateCategoryCommandHandler(),
		)
	}

	return f.actions["updateCategoryAction"].(*category.UpdateCategoryAction)
}

func (f ActionFactory) DeleteCategoryAction() *category.DeleteCategoryAction {
	_, ok := f.actions["deleteCategoryAction"]
	if !ok {
		f.actions["deleteCategoryAction"] = category.NewDeleteCategoryAction(
			f.commandHandlerFactory.DeleteCategoryCommandHandler(),
		)
	}

	return f.actions["deleteCategoryAction"].(*category.DeleteCategoryAction)
}

func (f ActionFactory) ListCategoriesAction() *category.ListCategoriesAction {
	_, ok := f.actions["listCategoriesAction"]
	if !ok {
		f.actions["listCategoriesAction"] = category.NewListCategoriesAction(
			f.queryHandlerFactory.ListCategoriesQueryHandler(),
		)
	}

	return f.actions["listCategoriesAction"].(*category.ListCategoriesAction)
}

func (f ActionFactory) ListCategoryIconsAction() *category.ListCategoryIconsAction {
	_, ok := f.actions["listCategoryIconsAction"]
	if !ok {
		f.actions["listCategoryIconsAction"] = category.NewListCategoryIconsAction(
			f.queryHandlerFactory.ListCategoryIconsQueryHandler(),
		)
	}

	return f.actions["listCategoryIconsAction"].(*category.ListCategoryIconsAction)
}

func (f ActionFactory) GetCategoryAction() *category.GetCategoryAction {
	_, ok := f.actions["getCategoryAction"]
	if !ok {
		f.actions["getCategoryAction"] = category.NewGetCategoryAction(
			f.queryHandlerFactory.GetCategoryQueryHandler(),
		)
	}

	return f.actions["getCategoryAction"].(*category.GetCategoryAction)
}

func (f ActionFactory) CreateExpenseAction() *expense.CreateExpenseAction {
	_, ok := f.actions["createExpenseAction"]
	if !ok {
		f.actions["createExpenseAction"] = expense.NewCreateExpenseAction(
			f.commandHandlerFactory.CreateExpenseCommandHandler(),
		)
	}

	return f.actions["createExpenseAction"].(*expense.CreateExpenseAction)
}

func (f ActionFactory) UpdateExpenseAction() *expense.UpdateExpenseAction {
	_, ok := f.actions["updateExpenseAction"]
	if !ok {
		f.actions["updateExpenseAction"] = expense.NewUpdateExpenseAction(
			f.commandHandlerFactory.UpdateExpenseCommandHandler(),
		)
	}

	return f.actions["updateExpenseAction"].(*expense.UpdateExpenseAction)
}

func (f ActionFactory) DeleteExpenseAction() *expense.DeleteExpenseAction {
	_, ok := f.actions["deleteExpenseAction"]
	if !ok {
		f.actions["deleteExpenseAction"] = expense.NewDeleteExpenseAction(
			f.commandHandlerFactory.DeleteExpenseCommandHandler(),
		)
	}

	return f.actions["deleteExpenseAction"].(*expense.DeleteExpenseAction)
}

func (f ActionFactory) ListExpensesAction() *expense.ListExpensesAction {
	_, ok := f.actions["listExpensesAction"]
	if !ok {
		f.actions["listExpensesAction"] = expense.NewListExpensesAction(
			f.queryHandlerFactory.ListExpensesQueryHandler(),
		)
	}

	return f.actions["listExpensesAction"].(*expense.ListExpensesAction)
}

func (f ActionFactory) ImportExpensesAction() *expense.ImportExpensesAction {
	_, ok := f.actions["importExpensesAction"]
	if !ok {
		f.actions["importExpensesAction"] = expense.NewImportExpensesAction(
			f.commandHandlerFactory.ImportExpensesCommandHandler(),
		)
	}

	return f.actions["importExpensesAction"].(*expense.ImportExpensesAction)
}

func (f ActionFactory) All() []rest.APIAction {
	return []rest.APIAction{
		f.CreateUserAction(),
		f.UpdateUserSettingsAction(),
		f.GetUserSettingsAction(),
		f.CreateAuthTokenAction(),
		f.CreateAccountAction(),
		f.UpdateAccountAction(),
		f.DeleteAccountAction(),
		f.ListAccountsAction(),
		f.CreateCategoryAction(),
		f.UpdateCategoryAction(),
		f.DeleteCategoryAction(),
		f.ListCategoriesAction(),
		f.GetCategoryAction(),
		f.ListCategoryIconsAction(),
		f.CreateExpenseAction(),
		f.UpdateExpenseAction(),
		f.DeleteExpenseAction(),
		f.ListExpensesAction(),
		f.ImportExpensesAction(),
	}
}
