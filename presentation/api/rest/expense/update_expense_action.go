package expense

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/command/expense/update_expense"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type UpdateExpenseAction struct {
	*rest.BaseAction
	handler *update_expense.UpdateExpenseCommandHandler
}

func NewUpdateExpenseAction(handler *update_expense.UpdateExpenseCommandHandler) *UpdateExpenseAction {
	return &UpdateExpenseAction{
		BaseAction: rest.NewBaseAction("Put", "/accounts/{accountID}/expenses/{expenseID}", false),
		handler:    handler,
	}
}

func (a UpdateExpenseAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	expenseID, err := strconv.ParseInt(chi.URLParam(r, "expenseID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	var command update_expense.UpdateExpenseCommand
	if err := a.Body(r, &command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	command.ID = expenseID
	command.UserID = a.GetAuthUserID(r)
	command.AccountID = accountID

	response, err := a.handler.Handle(command)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusCreated, response.Expense)
}
