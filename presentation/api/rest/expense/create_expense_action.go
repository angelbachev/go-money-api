package expense

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/command/expense/create_expense"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type CreateExpenseAction struct {
	*rest.BaseAction
	handler *create_expense.CreateExpenseCommandHandler
}

func NewCreateExpenseAction(handler *create_expense.CreateExpenseCommandHandler) *CreateExpenseAction {
	return &CreateExpenseAction{
		BaseAction: rest.NewBaseAction("Post", "/accounts/{accountID}/expenses", false),
		handler:    handler,
	}
}

func (a CreateExpenseAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	var command create_expense.CreateExpenseCommand
	if err := a.Body(r, &command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	command.UserID = a.GetAuthUserID(r)
	command.AccountID = accountID

	response, err := a.handler.Handle(command)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusCreated, response.Expense)
}
