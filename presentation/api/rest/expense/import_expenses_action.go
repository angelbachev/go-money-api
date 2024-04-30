package expense

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/command/expense/import_expenses"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type ImportExpensesAction struct {
	*rest.BaseAction
	handler *import_expenses.ImportExpensesCommandHandler
}

func NewImportExpensesAction(handler *import_expenses.ImportExpensesCommandHandler) *ImportExpensesAction {
	return &ImportExpensesAction{
		BaseAction: rest.NewBaseAction("Post", "/accounts/{accountID}/expenses/import", false),
		handler:    handler,
	}
}

func (a ImportExpensesAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	var command import_expenses.ImportExpensesCommand
	if err := a.Body(r, &command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	command.UserID = a.GetAuthUserID(r)
	command.AccountID = accountID
	command.File = file

	err = a.handler.Handle(command)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
